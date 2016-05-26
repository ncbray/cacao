package regenerate

import (
	"encoding/json"
	"fmt"
	"github.com/ncbray/compilerutil/writer"
	"io/ioutil"
	"os"
	"path/filepath"
)

type FieldDecl struct {
	Name string
	Type string
	Ref  bool
	List bool
}

type StructDecl struct {
	Name   string
	Fields []*FieldDecl
}

type GroupDecl struct {
	Name  string
	Nodes []*StructDecl
}

type TreeDecl struct {
	File   string
	Dump   string
	Nodes  []*StructDecl
	Groups []*GroupDecl
}

func generateNodeDecl(node *StructDecl, out *writer.TabbedWriter) {
	out.WriteLine(fmt.Sprintf("type %s struct {", node.Name))
	out.Indent()
	for _, field := range node.Fields {
		out.WriteString(field.Name)
		out.WriteString(" ")
		if field.List {
			out.WriteString("[]")
		}
		if field.Ref {
			out.WriteString("*")
		}
		out.WriteString(field.Type)
		out.EndOfLine()
	}
	out.Dedent()
	out.WriteLine("}")
}

func generateGroupDecl(group *GroupDecl, out *writer.TabbedWriter) {
	out.WriteLine(fmt.Sprintf("type %s interface {", group.Name))
	out.Indent()
	out.WriteLine(fmt.Sprintf("is%s()", group.Name))
	out.Dedent()
	out.WriteLine("}")

	out.EndOfLine()

	for _, node := range group.Nodes {
		generateNodeDecl(node, out)
		out.EndOfLine()

		out.WriteLine(fmt.Sprintf("func (node *%s) is%s() {", node.Name, group.Name))
		out.WriteLine("}")
		out.EndOfLine()
	}
}

func generateNodeDump(node *StructDecl, out *writer.TabbedWriter) {
	out.WriteLine(fmt.Sprintf("func dump%s(node *%s, out *writer.TabbedWriter) {", node.Name, node.Name))
	out.Indent()
	out.WriteLine(fmt.Sprintf("out.WriteString(\"%s {\")", node.Name))
	out.WriteLine("out.EndOfLine()")
	out.WriteLine("out.Indent()")
	for _, field := range node.Fields {
		if field.List {
			out.WriteString(fmt.Sprintf("out.WriteString(\"%s: []", field.Name))
			if field.Ref {
				out.WriteString("*")
			}
			out.WriteString(field.Type)
			out.WriteString(" {\")")
			out.EndOfLine()

			out.WriteLine("out.EndOfLine()")
			out.WriteLine("out.Indent()")

			out.WriteLine(fmt.Sprintf("for i, child := range node.%s {", field.Name))
			out.Indent()
			out.WriteLine("out.WriteString(fmt.Sprintf(\"%d: \", i))")
			out.WriteLine(fmt.Sprintf("dump%s(child, out)", field.Type))
			out.WriteLine("out.EndOfLine()")
			out.Dedent()
			out.WriteLine("}")

			out.WriteLine("out.Dedent()")
			out.WriteLine("out.WriteString(\"}\")")
			out.WriteLine("out.EndOfLine()")

		} else {
			if field.Ref {
				out.WriteLine(fmt.Sprintf("if node.%s != nil {", field.Name))
				out.Indent()
			}
			out.WriteLine(fmt.Sprintf("out.WriteString(\"%s: \")", field.Name))
			out.WriteLine(fmt.Sprintf("dump%s(node.%s, out)", field.Type, field.Name))
			out.WriteLine("out.EndOfLine()")
			if field.Ref {
				out.Dedent()
				out.WriteLine("}")
			}
		}
	}
	out.WriteLine("out.Dedent()")
	out.WriteLine("out.WriteString(\"}\")")

	out.Dedent()
	out.WriteLine("}")
}

func generateGroupDump(group *GroupDecl, out *writer.TabbedWriter) {
	for _, node := range group.Nodes {
		generateNodeDump(node, out)
		out.EndOfLine()
	}
	out.WriteLine(fmt.Sprintf("func dump%s(node %s, out *writer.TabbedWriter) {", group.Name, group.Name))
	out.Indent()
	out.WriteLine("switch node := node.(type) {")
	for _, node := range group.Nodes {
		out.WriteLine(fmt.Sprintf("case *%s:", node.Name))
		out.Indent()
		out.WriteLine(fmt.Sprintf("dump%s(node, out)", node.Name))
		out.Dedent()
	}
	out.WriteLine("default:")
	out.Indent()
	out.WriteLine("panic(node)")
	out.Dedent()
	out.WriteLine("}")
	out.Dedent()
	out.WriteLine("}")

}

func generateDump(filename string, decl *TreeDecl, output_dir string) {
	outfile := filepath.Join(output_dir, decl.Dump)
	fmt.Println("    ", filename, "=>", outfile)
	f, err := os.Create(outfile)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	out := writer.MakeTabbedWriter("\t", f)
	packageName := extractPackageName(outfile)
	rel_src, _ := filepath.Rel(filepath.Dir(outfile), filename)
	writeHeader(packageName, rel_src, out)

	imports := []string{
		"fmt",
		"github.com/ncbray/compilerutil/writer",
	}
	writeImports(imports, out)

	for _, node := range decl.Nodes {
		generateNodeDump(node, out)
		out.EndOfLine()
	}

	for _, group := range decl.Groups {
		generateGroupDump(group, out)
	}
	f.Close()
	formatGoFile(f.Name())
}

func generateTree(filename string, decl *TreeDecl, output_dir string) {
	outfile := filepath.Join(output_dir, decl.File)
	fmt.Println("    ", filename, "=>", outfile)
	f, err := os.Create(outfile)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	out := writer.MakeTabbedWriter("\t", f)
	packageName := extractPackageName(outfile)
	rel_src, _ := filepath.Rel(filepath.Dir(outfile), filename)
	writeHeader(packageName, rel_src, out)

	for _, node := range decl.Nodes {
		generateNodeDecl(node, out)
		out.EndOfLine()
	}

	for _, group := range decl.Groups {
		generateGroupDecl(group, out)
	}
	f.Close()
	formatGoFile(f.Name())
}

func ProcessTreeFile(filename string, output_dir string) {
	fmt.Println("Processing", filename)
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	decl := &TreeDecl{}
	err = json.Unmarshal(data, decl)
	if err != nil {
		panic(err)
	}
	generateTree(filename, decl, output_dir)
	if decl.Dump != "" {
		generateDump(filename, decl, output_dir)
	}
}
