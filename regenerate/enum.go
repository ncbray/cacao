package regenerate

import (
	"encoding/json"
	"fmt"
	"github.com/ncbray/compilerutil/fs"
	"github.com/ncbray/compilerutil/writer"
	"go/format"
	"io/ioutil"
	"path/filepath"
	"sort"
	"strings"
)

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

type parseTree struct {
	Value    string
	Children map[rune]*parseTree
}

func (tree *parseTree) Add(path []rune, value string) {
	if len(path) > 0 {
		current := path[0]
		child, ok := tree.Children[current]
		if !ok {
			child = newParseTree()
			tree.Children[current] = child
		}
		child.Add(path[1:], value)
	} else {
		if tree.Value != "" {
			panic(value)
		}
		tree.Value = value
	}
}

func newParseTree() *parseTree {
	return &parseTree{Children: map[rune]*parseTree{}}
}

type EnumDecl struct {
	Name   string
	Fields map[string]interface{}
}

type EnumFieldDecl struct {
	Name string
	Type string
}

type EnumIndexDecl struct {
	Name string
	Path []string
}

type EnumTypeDecl struct {
	File           string
	Name           string
	Prefix         string
	GenerateParser string
	Fields         []*EnumFieldDecl
	Indexes        []*EnumIndexDecl
	Enums          []*EnumDecl
}

func (decl *EnumTypeDecl) getField(name string) (*EnumFieldDecl, bool) {
	for _, f := range decl.Fields {
		if f.Name == name {
			return f, true
		}
	}
	return nil, false
}

func fallback(a string, b string) string {
	if a != "" {
		return a
	}
	return b
}

func generateParseTree(tree *parseTree, fail string, out *writer.TabbedWriter) {
	if len(tree.Children) > 0 {
		runes := []int{}
		for r := range tree.Children {
			runes = append(runes, int(r))
		}
		sort.Ints(runes)

		out.WriteLine("switch state.Current() {")
		for _, i := range runes {
			r := rune(i)
			child := tree.Children[r]
			out.WriteLine(fmt.Sprintf("case %q:", r))
			out.Indent()
			out.WriteLine("state.Consume()")
			generateParseTree(child, fail, out)
			out.Dedent()
		}
		out.WriteLine("default:")
		out.Indent()
		out.WriteLine(fmt.Sprintf("return %s", fallback(tree.Value, fail)))
		out.Dedent()
		out.WriteLine("}")
	} else {
		out.WriteLine(fmt.Sprintf("return %s", fallback(tree.Value, fail)))
	}
}

func extractPackageName(filename string) string {
	parts := strings.Split(filename, "/")
	return parts[len(parts)-2]
}

func writeHeader(packageName string, sourceFile string, out *writer.TabbedWriter) {
	out.WriteLine("package " + packageName)
	out.EndOfLine()
	out.WriteLine(fmt.Sprintf("/* Generated from %s, do not edit by hand. */", sourceFile))
	out.EndOfLine()
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
		out.EndOfLine()
	}
}

func formatFieldValue(f *EnumFieldDecl, value interface{}) string {
	switch value := value.(type) {
	case string:
		if f.Type == "string" {
			return fmt.Sprintf("%#v", value)
		} else {
			// Must be a named reference rather than a string literal.
			return value
		}
	default:
		return fmt.Sprintf("%#v", value)
	}
}

func pathRemainerType(decl *EnumTypeDecl, idx *EnumIndexDecl, pos int) string {
	if pos >= len(idx.Path) {
		return "*" + decl.Name + "Info"
	} else {
		name := idx.Path[pos]
		f, ok := decl.getField(name)
		if !ok {
			panic(name)
		}
		return "map[" + f.Type + "]" + pathRemainerType(decl, idx, pos+1)
	}
}

func WriteIndex(decl *EnumTypeDecl, idx *EnumIndexDecl, pos int, enums []*EnumDecl, out *writer.TabbedWriter) {
	if pos >= len(idx.Path) {
		if len(enums) != 1 {
			panic(enums)
		}
		e := enums[0]
		out.WriteString(decl.Name)
		out.WriteString("Infos[")
		out.WriteString(decl.Prefix)
		out.WriteString("_")
		out.WriteString(e.Name)
		out.WriteString("]")
	} else {
		name := idx.Path[pos]
		f, ok := decl.getField(name)
		if !ok {
			panic(name)
		}

		if pos == 0 {
			out.WriteString(pathRemainerType(decl, idx, pos))
		}
		out.WriteString(" {")
		out.EndOfLine()
		out.Indent()

		lut := map[string][]*EnumDecl{}
		keys := []string{}

		for _, e := range enums {
			value, ok := e.Fields[f.Name]
			if !ok {
				continue
			}
			// HACK bucket by string value
			text := formatFieldValue(f, value)
			existing, ok := lut[text]
			if !ok {
				keys = append(keys, text)
			}
			lut[text] = append(existing, e)
		}

		for _, k := range keys {
			out.WriteString(k)
			out.WriteString(": ")
			WriteIndex(decl, idx, pos+1, lut[k], out)
			out.WriteString(",")
			out.EndOfLine()
		}

		out.Dedent()
		out.WriteString("}")
	}
}

type ProcessEnumContext struct {
	Input  fs.DataInput
	Output fs.DataOutput
}

func generateEnumGo(decl *EnumTypeDecl, packageName string, rel_src string, out *writer.TabbedWriter) {
	writeHeader(packageName, rel_src, out)

	imports := []string{}
	if decl.GenerateParser != "" {
		imports = append(imports, "github.com/ncbray/cacao/framework")
	}
	writeImports(imports, out)

	out.WriteLine(fmt.Sprintf("type %s int", decl.Name))
	out.EndOfLine()
	out.WriteLine("const (")
	out.Indent()
	out.WriteLine(fmt.Sprintf("INVALID_%s %s = %d", decl.Prefix, decl.Name, -1))
	for i, e := range decl.Enums {
		fullName := decl.Prefix + "_" + e.Name
		out.WriteLine(fmt.Sprintf("%s %s = %d", fullName, decl.Name, i))
	}
	out.WriteLine(fmt.Sprintf("NUM_%s %s = %d", decl.Prefix, decl.Name, len(decl.Enums)))
	out.Dedent()
	out.WriteLine(")")
	out.EndOfLine()

	out.WriteLine(fmt.Sprintf("type %sInfo struct {", decl.Name))
	out.Indent()
	out.WriteLine("Enum " + decl.Name)
	out.WriteLine("Name string")
	for _, f := range decl.Fields {
		out.WriteLine(fmt.Sprintf("%s %s", f.Name, f.Type))
	}
	out.Dedent()
	out.WriteLine("}")
	out.EndOfLine()

	out.WriteLine(fmt.Sprintf("var %sInfos = []*%sInfo {", decl.Name, decl.Name))
	out.Indent()
	for _, e := range decl.Enums {
		fullName := decl.Prefix + "_" + e.Name
		out.WriteLine(" {")
		out.Indent()
		out.WriteLine(fmt.Sprintf("Enum: %s,", fullName))
		out.WriteLine(fmt.Sprintf("Name: %#v,", fullName))
		for _, f := range decl.Fields {
			value, ok := e.Fields[f.Name]
			if ok {
				out.WriteLine(fmt.Sprintf("%s: %s,", f.Name, formatFieldValue(f, value)))
			}
		}
		out.Dedent()
		out.WriteLine("},")
	}
	out.Dedent()
	out.WriteLine("}")
	out.EndOfLine()

	out.WriteLine(fmt.Sprintf("func (value %s) String() string {", decl.Name))
	out.Indent()
	out.WriteLine(fmt.Sprintf("if value < 0 || value >= %d {", len(decl.Enums)))
	out.Indent()
	out.WriteLine(fmt.Sprintf("return %#v", "INVALID_"+decl.Prefix))
	out.Dedent()
	out.WriteLine("}")

	out.WriteLine(fmt.Sprintf("return %sInfos[value].Name", decl.Name))
	out.Dedent()
	out.WriteLine("}")
	out.EndOfLine()

	for _, idx := range decl.Indexes {
		out.WriteString(fmt.Sprintf("var %s = ", idx.Name))
		WriteIndex(decl, idx, 0, decl.Enums, out)
		out.EndOfLine()
	}

	if decl.GenerateParser != "" {
		tree := newParseTree()
		fail := "INVALID_" + decl.Prefix + ", false"
		for _, e := range decl.Enums {
			key := e.Fields[decl.GenerateParser].(string)
			tree.Add([]rune(key), decl.Prefix+"_"+e.Name+", true")
		}

		out.WriteLine(fmt.Sprintf("func Parse%s(state *framework.ParserState) (%s, bool) {", decl.Name, decl.Name))
		out.Indent()
		generateParseTree(tree, fail, out)
		out.Dedent()
		out.WriteLine("}")
		out.EndOfLine()
	}
}

func processEnumFile(filename string, output_dir string, fsys fs.FileSystem) {
	// Parse the JSON
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	decl := &EnumTypeDecl{}
	err = json.Unmarshal(data, decl)
	if err != nil {
		panic(err)
	}

	outfile := filepath.Join(output_dir, decl.File)
	fmt.Println("    ", filename, "=>", outfile)

	t := fsys.TempFile()
	tw, err := t.GetWriter()
	if err != nil {
		panic(err)
	}
	out := writer.MakeTabbedWriter("\t", tw)
	packageName := extractPackageName(outfile)
	rel_src, _ := filepath.Rel(filepath.Dir(outfile), filename)
	generateEnumGo(decl, packageName, rel_src, out)
	tw.Close()

	formatGoFile(t, fsys.OutputFile(outfile, 0640))
}

func ProcessEnumFile(filename string, output_dir string, fsys fs.FileSystem) {
	fmt.Println("Processing", filename)
	processEnumFile(filename, output_dir, fsys)
}
