package irgen

import (
	"github.com/ncbray/compilerutil/names"
	"github.com/ncbray/compilerutil/writer"
	"text/template"
)

var graphCodegenTemplates = template.Must(template.New("main").Parse(`
{{define "parent_relationship" -}}
func (src *{{.Src.Impl.Name}}) appendTo{{.SrcName}}(dst *{{.Dst.Impl.Name}}) {
	if dst.{{.PrevField.Name}} != nil || dst.{{.NextField.Name}} != nil {
		panic(dst)
	}
	if src.{{.HeadField.Name}} == nil {
		src.{{.HeadField.Name}} = dst
		src.{{.TailField.Name}} = dst
	} else {
		tail := src.{{.TailField.Name}}
		tail.{{.NextField.Name}} = dst
		dst.{{.PrevField.Name}} = tail
		src.{{.TailField.Name}} = dst
	}
}

func (src *{{.Src.Impl.Name}}) Iter{{.SrcName}}() {{.Iterator.Name}} {
	return {{.Iterator.Name}}{current: src.{{.HeadField.Name}}}
}

func (iter *{{.Iterator.Name}}) HasNext() bool {
	return iter.current != nil
}

func (iter *{{.Iterator.Name}}) GetNext() *{{.Dst.Impl.Name}} {
	temp := iter.current
	iter.current = temp.{{.NextField.Name}}
	return temp
}
{{end}}

{{define "counted_relationship" -}}
func (src *{{.Src.Impl.Name}}) Create{{.SrcName}}(count int) {{.Creator.Name}} {
	src.{{.CountedField.Name}} = make([]*{{.Dst.Impl.Name}}, count)
	return {{.Creator.Name}}{src: src}
}

func (iter *{{.Creator.Name}}) SetNext(dst *{{.Dst.Impl.Name}}) {
	iter.src.{{.CountedField.Name}}[iter.index] = dst
	iter.index++
}

func (src *{{.Src.Impl.Name}}) Iter{{.SrcName}}() {{.Iterator.Name}} {
	return {{.Iterator.Name}}{src: src}
}

func (iter *{{.Iterator.Name}}) HasNext() bool {
	return iter.index < len(iter.src.{{.CountedField.Name}})
}

func (iter *{{.Iterator.Name}}) GetNext() *{{.Dst.Impl.Name}} {
	temp := iter.src.{{.CountedField.Name}}[iter.index]
	iter.index++
	return temp
}
{{end}}
`))

type NodeDecl struct {
	Name   string
	Region string
}

type ParentRelationshipDecl struct {
	Src     string
	SrcName string

	Dst     string
	DstName string

	Required bool
}

type CountedRelationshipDecl struct {
	Src     string
	SrcName string

	Dst     string
	DstName string
}

type Relationship interface {
	isRelationship()
}

type NodeInfo struct {
	Decl *NodeDecl
	Impl *LLStruct

	Region               *NodeInfo
	Allocated            []*NodeInfo
	ForwardRelationships []Relationship
	ReverseRelationships []Relationship
}

type ParentRelationshipInfo struct {
	Src     *NodeInfo
	SrcName string

	Dst     *NodeInfo
	DstName string

	HeadField *LLField
	TailField *LLField

	NextField *LLField
	PrevField *LLField

	Iterator *LLStruct

	Required bool
}

func (r *ParentRelationshipInfo) isRelationship() {
}

type CountedRelationshipInfo struct {
	Src     *NodeInfo
	SrcName string

	Dst     *NodeInfo
	DstName string

	CountedField *LLField

	Creator  *LLStruct
	Iterator *LLStruct
}

func (r *CountedRelationshipInfo) isRelationship() {
}

func getNodeImports(decls []*NodeDecl, imports map[string]bool) {
}

func generateCommaList(args []string, out *writer.TabbedWriter) {
	for i, arg := range args {
		if i > 0 {
			out.WriteString(", ")
		}
		out.WriteString(arg)
	}
}

func ptrToStruct(s *LLStruct) LLType {
	return &PointerType{
		Element: &StructType{
			Element: s,
		},
	}
}

func createField(s *LLStruct, name string, t LLType) *LLField {
	f := &LLField{Name: name, Type: t}
	s.Fields = append(s.Fields, f)
	return f
}

func toLocalVarName(name string) string {
	return names.JoinSnakeCase(names.SplitCamelCase(name), false)
}

func constructorGenerator(info *NodeInfo) *callbackGenerator {
	return &callbackGenerator{callback: func(out *writer.TabbedWriter) {
		generateConstructor(info, out)
	}}
}

func generateConstructor(info *NodeInfo, out *writer.TabbedWriter) {
	name := info.Impl.Name
	out.WriteString("func ")
	if info.Region != nil {
		out.WriteString("(region *")
		out.WriteString(info.Region.Impl.Name)
		out.WriteString(") ")
	}
	out.WriteString("Create")
	out.WriteString(name)
	out.WriteString("(")
	args := []string{}
	for _, rel := range info.ReverseRelationships {
		switch rel := rel.(type) {
		case *ParentRelationshipInfo:
			if rel.Required {
				args = append(args, toLocalVarName(rel.DstName)+" *"+rel.Src.Impl.Name)
			}
		}
	}
	generateCommaList(args, out)
	out.WriteString(") *")
	out.WriteString(name)
	out.WriteString(" {")
	out.EndOfLine()
	out.Indent()
	out.WriteString("o := &")
	out.WriteString(name)
	out.WriteString("{}")
	out.EndOfLine()
	for _, rel := range info.ReverseRelationships {
		switch rel := rel.(type) {
		case *ParentRelationshipInfo:
			if rel.Required {
				out.WriteString(toLocalVarName(rel.DstName))
				out.WriteString(".appendTo")
				out.WriteString(rel.SrcName)
				out.WriteString("(o)")
				out.EndOfLine()
			}
		}
	}
	out.WriteLine("return o")
	out.Dedent()
	out.WriteLine("}")
}

func generateGraph(decls []*NodeDecl, parent_relationship_decls []*ParentRelationshipDecl, counted_relationship_decls []*CountedRelationshipDecl, out *writer.TabbedWriter) {
	prog := &LLProgram{}

	nodes := make([]*NodeInfo, len(decls))
	lut := map[string]*NodeInfo{}

	// Initialize and index.
	for i, decl := range decls {
		info := &NodeInfo{Decl: decl, Impl: prog.CreateLLStruct(decl.Name, true)}
		nodes[i] = info
		lut[decl.Name] = info
	}

	// Region relationships.
	for _, info := range nodes {
		decl := info.Decl
		if decl.Region != "" {
			ri, ok := lut[decl.Region]
			if !ok {
				panic(decl.Region)
			}
			info.Region = ri
			ri.Allocated = append(ri.Allocated, info)
		}
	}

	// One-to-many relationships.
	for _, decl := range parent_relationship_decls {
		si, ok := lut[decl.Src]
		if !ok {
			panic(decl.Src)
		}
		di, ok := lut[decl.Dst]
		if !ok {
			panic(decl.Dst)
		}

		rel := &ParentRelationshipInfo{
			Src:      si,
			SrcName:  decl.SrcName,
			Dst:      di,
			DstName:  decl.DstName,
			Required: decl.Required,
		}
		si.ForwardRelationships = append(si.ForwardRelationships, rel)
		di.ReverseRelationships = append(di.ReverseRelationships, rel)

		// Src fields
		rel.HeadField = createField(si.Impl, rel.SrcName+"Head", ptrToStruct(di.Impl))
		rel.TailField = createField(si.Impl, rel.SrcName+"Tail", ptrToStruct(di.Impl))

		// Dst fields
		rel.NextField = createField(di.Impl, rel.SrcName+"Next", ptrToStruct(di.Impl))
		rel.PrevField = createField(di.Impl, rel.SrcName+"Prev", ptrToStruct(di.Impl))

		// Iterator structure
		rel.Iterator = prog.CreateLLStruct(si.Impl.Name+decl.SrcName+"Iterator", false)
		createField(rel.Iterator, "Current", ptrToStruct(di.Impl))
	}

	for _, decl := range counted_relationship_decls {
		si, ok := lut[decl.Src]
		if !ok {
			panic(decl.Src)
		}
		di, ok := lut[decl.Dst]
		if !ok {
			panic(decl.Dst)
		}

		rel := &CountedRelationshipInfo{
			Src:     si,
			SrcName: decl.SrcName,
			Dst:     di,
			DstName: decl.DstName,
		}
		si.ForwardRelationships = append(si.ForwardRelationships, rel)
		di.ReverseRelationships = append(di.ReverseRelationships, rel)

		rel.CountedField = createField(si.Impl, rel.SrcName, &ListType{Element: ptrToStruct(di.Impl)})

		rel.Creator = prog.CreateLLStruct(si.Impl.Name+decl.SrcName+"Creator", false)
		createField(rel.Creator, "Src", ptrToStruct(si.Impl))
		createField(rel.Creator, "Index", &IntrinsicType{Element: "int"})

		rel.Iterator = prog.CreateLLStruct(si.Impl.Name+decl.SrcName+"Iterator", false)
		createField(rel.Iterator, "Src", ptrToStruct(si.Impl))
		createField(rel.Iterator, "Index", &IntrinsicType{Element: "int"})
	}

	// Fix up names to match Go's symbol exporting scheme.
	for _, s := range prog.structs {
		s.Name = names.JoinCamelCase(names.SplitCamelCase(s.Name), s.Export)
		for _, f := range s.Fields {
			f.Name = names.JoinCamelCase(names.SplitCamelCase(f.Name), f.Export)
		}
	}

	// Prep for code generation.
	file_decls := []LLDecl{}
	for _, info := range nodes {
		file_decls = append(file_decls, info.Impl)
		file_decls = append(file_decls, constructorGenerator(info))
		for _, rel := range info.ForwardRelationships {
			switch rel := rel.(type) {
			case *ParentRelationshipInfo:
				file_decls = append(file_decls, rel.Iterator)
				file_decls = append(file_decls, &templateGenerator{tmpl: graphCodegenTemplates, name: "parent_relationship", data: rel})
			case *CountedRelationshipInfo:
				file_decls = append(file_decls, rel.Creator)
				file_decls = append(file_decls, rel.Iterator)
				file_decls = append(file_decls, &templateGenerator{tmpl: graphCodegenTemplates, name: "counted_relationship", data: rel})
			default:
				panic(rel)
			}
		}
	}
	generateGo(file_decls, out)
}
