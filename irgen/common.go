package irgen

import (
	"fmt"
	"github.com/ncbray/compilerutil/fs"
	"github.com/ncbray/compilerutil/writer"
	"go/format"
	"os"
	"path/filepath"
	"strings"
)

func isZeroValue(v interface{}) bool {
	switch v := v.(type) {
	case bool:
		return !v
	case int:
		return v == 0
	case string:
		return v == ""
	default:
		return v == nil
	}
}

func dirExists(path string) bool {
	_, err := os.Stat(path)
	return !os.IsNotExist(err)
}

func formatGoFile(src fs.DataInput, dst fs.DataOutput) error {
	data, err := src.GetBytes()
	if err != nil {
		return err
	}
	data, err = format.Source(data)
	if err != nil {
		return err
	}
	return dst.SetBytes(data)
}

func writeHeader(packageName string, sourceFile string, out *writer.TabbedWriter) {
	out.WriteLine("package " + packageName)
	out.EndOfLine()
	out.WriteLine(fmt.Sprintf("/* Generated from %s, do not edit by hand. */", sourceFile))
}

func extractPackageName(filename string) string {
	parts := strings.Split(filename, "/")
	return parts[len(parts)-2]
}

func writeImports(imports []string, out *writer.TabbedWriter) {
	if len(imports) > 0 {
		out.WriteLine("import (")
		out.Indent()
		for _, imp := range imports {
			out.WriteLine(fmt.Sprintf("%#v", imp))
		}
		out.Dedent()
		out.WriteLine(")")
	}
}

type Declarations struct {
	DataSource string
	Package    string
	File       string
	Enums      []*EnumDecl
	Trees      []*TreeDecl
	Nodes      []*NodeDecl
	Parents    []*ParentRelationshipDecl
	Counted    []*CountedRelationshipDecl
}

func processDeclaration(decl *Declarations, output_dir string, fsys fs.FileSystem) {
	package_dir := filepath.Join(output_dir, decl.Package)
	if !dirExists(package_dir) {
		panic(package_dir)
	}

	outfile := filepath.Join(package_dir, decl.File)
	//fmt.Println("generating:", filepath.Join(decl.Package, decl.File))

	t := fsys.TempFile()
	tw, err := t.GetWriter()
	if err != nil {
		panic(err)
	}
	defer tw.Close()
	out := writer.MakeTabbedWriter("\t", tw)
	packageName := extractPackageName(outfile)
	writeHeader(packageName, decl.DataSource, out)

	imports := map[string]bool{}
	for _, e := range decl.Enums {
		getEnumImports(e, imports)
	}
	for _, t := range decl.Trees {
		getTreeImports(t, imports)
	}
	if len(decl.Nodes) > 0 {
		getNodeImports(decl.Nodes, imports)
	}

	if len(imports) > 0 {
		importList := make([]string, len(imports))
		i := 0
		for k, _ := range imports {
			importList[i] = k
			i++
		}
		out.EndOfLine()
		writeImports(importList, out)
	}

	for _, e := range decl.Enums {
		out.EndOfLine()
		generateEnum(e, out)
	}
	for _, t := range decl.Trees {
		out.EndOfLine()
		generateTree(t, out)
	}
	if len(decl.Nodes) > 0 {
		generateGraph(decl.Nodes, decl.Parents, decl.Counted, out)
	}

	tw.Close()
	err = formatGoFile(t, fsys.OutputFile(outfile, 0640))
	if err != nil {
		panic(err)
	}
}

func ProcessDeclarations(decls []*Declarations, output_dir string, fsys fs.FileSystem) {
	for _, decl := range decls {
		processDeclaration(decl, output_dir, fsys)
	}
}
