package regenerate

import (
	"fmt"
	"github.com/ncbray/compilerutil/writer"
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
	Name    string
	Structs []*StructDecl
}

type TreeDecl struct {
	Dump    bool
	Structs []*StructDecl
	Groups  []*GroupDecl
}

func generateStructDecl(node *StructDecl, out *writer.TabbedWriter) {
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

	for _, node := range group.Structs {
		generateStructDecl(node, out)
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
	for _, node := range group.Structs {
		generateNodeDump(node, out)
		out.EndOfLine()
	}
	out.WriteLine(fmt.Sprintf("func dump%s(node %s, out *writer.TabbedWriter) {", group.Name, group.Name))
	out.Indent()
	out.WriteLine("switch node := node.(type) {")
	for _, node := range group.Structs {
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

func generateDump(decl *TreeDecl, out *writer.TabbedWriter) {
	for _, node := range decl.Structs {
		generateNodeDump(node, out)
		out.EndOfLine()
	}

	for _, group := range decl.Groups {
		generateGroupDump(group, out)
	}
}

func getTreeImports(decl *TreeDecl, imports map[string]bool) {
	if decl.Dump {
		imports["fmt"] = true
		imports["github.com/ncbray/compilerutil/writer"] = true
	}
}

func generateTree(decl *TreeDecl, out *writer.TabbedWriter) {
	for _, node := range decl.Structs {
		generateStructDecl(node, out)
		out.EndOfLine()
	}
	for _, group := range decl.Groups {
		generateGroupDecl(group, out)
	}
	if decl.Dump {
		out.EndOfLine()
		generateDump(decl, out)
	}
}
