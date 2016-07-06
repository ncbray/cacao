package irgen

import (
	"github.com/ncbray/compilerutil/writer"
	"text/template"
)

func generateType(t LLType, out *writer.TabbedWriter) {
	for {
		switch st := t.(type) {
		case *PointerType:
			out.WriteString("*")
			t = st.Element
		case *ListType:
			out.WriteString("[]")
			t = st.Element
		case *StructType:
			out.WriteString(st.Element.Name)
			return
		case *IntrinsicType:
			out.WriteString(st.Element)
			return
		default:
			panic(t)
		}
	}
}

func generateStruct(s *LLStruct, out *writer.TabbedWriter) {
	out.WriteString("type ")
	out.WriteString(s.Name)
	out.WriteString(" struct {")
	out.EndOfLine()
	out.Indent()
	for _, f := range s.Fields {
		out.WriteString(f.Name)
		out.WriteString(" ")
		generateType(f.Type, out)
		out.EndOfLine()
	}
	out.Dedent()
	out.WriteLine("}")
}

type LLProgram struct {
	structs []*LLStruct
}

func (prog *LLProgram) CreateLLStruct(name string, export bool) *LLStruct {
	s := &LLStruct{Name: name, Export: export}
	prog.structs = append(prog.structs, s)
	return s
}

type LLDecl interface {
	isLLDecl()
}

func (decl *LLStruct) isLLDecl() {
}

type callbackGenerator struct {
	callback func(out *writer.TabbedWriter)
}

func (decl *callbackGenerator) isLLDecl() {
}

type templateGenerator struct {
	tmpl *template.Template
	name string
	data interface{}
}

func (decl *templateGenerator) isLLDecl() {
}

func generateGo(decls []LLDecl, out *writer.TabbedWriter) {
	for _, decl := range decls {
		switch decl := decl.(type) {
		case *LLStruct:
			out.EndOfLine()
			generateStruct(decl, out)
		case *callbackGenerator:
			out.EndOfLine()
			decl.callback(out)
		case *templateGenerator:
			out.EndOfLine()
			err := decl.tmpl.ExecuteTemplate(out, decl.name, decl.data)
			if err != nil {
				panic(err)
			}
		default:
			panic(decl)
		}
	}
}
