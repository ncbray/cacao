package irgen

/* Generated from regenerate_cacao, do not edit by hand. */

import (
	"github.com/ncbray/compilerutil/writer"
	"strconv"
)

type LLField struct {
	Name   string
	Type   LLType
	Export bool
}

type LLType interface {
	isLLType()
}

type LLIntrinsicType struct {
	Element string
}

func (node *LLIntrinsicType) isLLType() {
}

type LLStruct struct {
	Name      string
	Fields    []*LLField
	Export    bool
	PtrCache  *PointerType
	ListCache *ListType
}

func (node *LLStruct) isLLType() {
}

type PointerType struct {
	Element LLType
}

func (node *PointerType) isLLType() {
}

type ListType struct {
	Element LLType
}

func (node *ListType) isLLType() {
}

func dumpLLField(node *LLField, out *writer.TabbedWriter) {
	out.WriteString("LLField {")
	out.EndOfLine()
	out.Indent()
	out.WriteString("Name: ")
	out.WriteString(strconv.Quote(node.Name))
	out.EndOfLine()
	out.WriteString("Type: ")
	dumpLLType(node.Type, out)
	out.EndOfLine()
	out.WriteString("Export: ")
	out.WriteString(strconv.FormatBool(node.Export))
	out.EndOfLine()
	out.Dedent()
	out.WriteString("}")
}

func dumpLLIntrinsicType(node *LLIntrinsicType, out *writer.TabbedWriter) {
	out.WriteString("LLIntrinsicType {")
	out.EndOfLine()
	out.Indent()
	out.WriteString("Element: ")
	out.WriteString(strconv.Quote(node.Element))
	out.EndOfLine()
	out.Dedent()
	out.WriteString("}")
}

func dumpLLStruct(node *LLStruct, out *writer.TabbedWriter) {
	out.WriteString("LLStruct {")
	out.EndOfLine()
	out.Indent()
	out.WriteString("Name: ")
	out.WriteString(strconv.Quote(node.Name))
	out.EndOfLine()
	out.WriteString("Fields: []*LLField {")
	out.EndOfLine()
	out.Indent()
	for i, child := range node.Fields {
		out.WriteString(strconv.Itoa(i))
		out.WriteString(": ")
		dumpLLField(child, out)
		out.EndOfLine()
	}
	out.Dedent()
	out.WriteString("}")
	out.EndOfLine()
	out.WriteString("Export: ")
	out.WriteString(strconv.FormatBool(node.Export))
	out.EndOfLine()
	out.Dedent()
	out.WriteString("}")
}

func dumpPointerType(node *PointerType, out *writer.TabbedWriter) {
	out.WriteString("PointerType {")
	out.EndOfLine()
	out.Indent()
	out.WriteString("Element: ")
	dumpLLType(node.Element, out)
	out.EndOfLine()
	out.Dedent()
	out.WriteString("}")
}

func dumpListType(node *ListType, out *writer.TabbedWriter) {
	out.WriteString("ListType {")
	out.EndOfLine()
	out.Indent()
	out.WriteString("Element: ")
	dumpLLType(node.Element, out)
	out.EndOfLine()
	out.Dedent()
	out.WriteString("}")
}

func dumpLLType(node LLType, out *writer.TabbedWriter) {
	switch node := node.(type) {
	case *LLIntrinsicType:
		dumpLLIntrinsicType(node, out)
	case *LLStruct:
		dumpLLStruct(node, out)
	case *PointerType:
		dumpPointerType(node, out)
	case *ListType:
		dumpListType(node, out)
	default:
		panic(node)
	}
}
