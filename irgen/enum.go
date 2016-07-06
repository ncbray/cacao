package irgen

import (
	"fmt"
	"github.com/ncbray/compilerutil/fs"
	"github.com/ncbray/compilerutil/writer"
	"sort"
)

type parseTree struct {
	Action   string
	Children map[rune]*parseTree
}

func (tree *parseTree) Add(path []rune, action string) {
	if len(path) > 0 {
		current := path[0]
		child, ok := tree.Children[current]
		if !ok {
			child = newParseTree()
			tree.Children[current] = child
		}
		child.Add(path[1:], action)
	} else {
		if tree.Action != "" {
			panic(action)
		}
		tree.Action = action
	}
}

func newParseTree() *parseTree {
	return &parseTree{Children: map[rune]*parseTree{}}
}

type EnumInstanceField struct {
	Name  string
	Value interface{}
}

type EnumInstance struct {
	Name   string
	Fields []*EnumInstanceField
}

type EnumDeclField struct {
	Name string
	Type string
}

type EnumIndex struct {
	Name string
	Path []string
}

type EnumDecl struct {
	Name           string
	Prefix         string
	GenerateParser string
	Fields         []*EnumDeclField
	Indexes        []*EnumIndex
	Enums          []*EnumInstance
}

func (decl *EnumDecl) getField(name string) (*EnumDeclField, bool) {
	for _, f := range decl.Fields {
		if f.Name == name {
			return f, true
		}
	}
	return nil, false
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
		if tree.Action != "" {
			out.WriteLine(tree.Action)
		} else if fail != "" {
			out.WriteLine(fail)
		}
		out.Dedent()
		out.WriteLine("}")
	} else {
		if tree.Action != "" {
			out.WriteLine(tree.Action)
		} else if fail != "" {
			out.WriteLine(fail)
		}
	}
}

func formatFieldValue(f *EnumDeclField, value interface{}) string {
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

func pathRemainerType(decl *EnumDecl, idx *EnumIndex, pos int) string {
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

func (e *EnumInstance) getField(name string) (*EnumInstanceField, bool) {
	for _, f := range e.Fields {
		if f.Name == name {
			return f, true
		}
	}
	return nil, false
}

func WriteIndex(decl *EnumDecl, idx *EnumIndex, pos int, enums []*EnumInstance, out *writer.TabbedWriter) {
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

		lut := map[string][]*EnumInstance{}
		keys := []string{}

		for _, e := range enums {
			value, ok := e.getField(f.Name)
			if !ok {
				continue
			}
			// HACK bucket by string value
			text := formatFieldValue(f, value.Value)
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

func getEnumImports(decl *EnumDecl, imports map[string]bool) {
	if decl.GenerateParser != "" {
		imports["github.com/ncbray/cacao/framework"] = true
	}
}

func typeForEnum(decl *EnumDecl) string {
	n := len(decl.Enums)
	if n <= 256 {
		return "uint8"
	} else if n <= 65536 {
		return "uint16"
	} else {
		return "uint32"
	}
}

func generateEnum(decl *EnumDecl, out *writer.TabbedWriter) {
	out.WriteLine(fmt.Sprintf("type %s %s", decl.Name, typeForEnum(decl)))
	out.EndOfLine()
	out.WriteLine("const (")
	out.Indent()
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
			value, ok := e.getField(f.Name)
			if ok && !isZeroValue(value.Value) {
				out.WriteLine(fmt.Sprintf("%s: %s,", f.Name, formatFieldValue(f, value.Value)))
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
		//fullName := decl.Prefix + "_" + e.Name
		for _, e := range decl.Enums {
			f, _ := e.getField(decl.GenerateParser)
			key := f.Value.(string)
			tree.Add([]rune(key), "return "+decl.Prefix+"_"+e.Name+", true")
		}

		out.WriteLine(fmt.Sprintf("func Parse%s(state *framework.ParserState) (%s, bool) {", decl.Name, decl.Name))
		out.Indent()
		generateParseTree(tree, "return 0, false", out)
		out.Dedent()
		out.WriteLine("}")
		out.EndOfLine()
	}
}
