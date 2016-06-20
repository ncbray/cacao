package language

/* Generated from regenerate_cacao, do not edit by hand. */

import (
	"fmt"
	"github.com/ncbray/compilerutil/writer"
)

type Token struct {
	Pos  int
	Text string
}

type Param struct {
	Name *Token
	Type TypeRef
}

type File struct {
	Decls []*FunctionDecl
}

type TypeRef interface {
	isTypeRef()
}

type TypeName struct {
	Name *Token
}

func (node *TypeName) isTypeRef() {
}

type Expr interface {
	isExpr()
}

type GetName struct {
	Name *Token
}

func (node *GetName) isExpr() {
}

type IntegerLiteral struct {
	Pos  int
	Text string
}

func (node *IntegerLiteral) isExpr() {
}

type PrefixExpr struct {
	Op   PrefixOperator
	Expr Expr
}

func (node *PrefixExpr) isExpr() {
}

type InfixExpr struct {
	Pos   int
	Left  Expr
	Op    InfixOperator
	Right Expr
}

func (node *InfixExpr) isExpr() {
}

type Return struct {
	Pos  int
	Expr Expr
}

func (node *Return) isExpr() {
}

type Call struct {
	Pos  int
	Expr Expr
	Args []Expr
}

func (node *Call) isExpr() {
}

type If struct {
	Pos  int
	Cond Expr
	T    *Block
	F    *Block
}

func (node *If) isExpr() {
}

type Block struct {
	Pos   int
	Exprs []Expr
}

func (node *Block) isExpr() {
}

type FunctionDecl struct {
	Name    *Token
	Params  []*Param
	Results []*Param
	Body    *Block
}

func (node *FunctionDecl) isExpr() {
}

func dumpToken(node *Token, out *writer.TabbedWriter) {
	out.WriteString("Token {")
	out.EndOfLine()
	out.Indent()
	out.WriteString("Pos: ")
	dumpint(node.Pos, out)
	out.EndOfLine()
	out.WriteString("Text: ")
	dumpstring(node.Text, out)
	out.EndOfLine()
	out.Dedent()
	out.WriteString("}")
}

func dumpParam(node *Param, out *writer.TabbedWriter) {
	out.WriteString("Param {")
	out.EndOfLine()
	out.Indent()
	if node.Name != nil {
		out.WriteString("Name: ")
		dumpToken(node.Name, out)
		out.EndOfLine()
	}
	out.WriteString("Type: ")
	dumpTypeRef(node.Type, out)
	out.EndOfLine()
	out.Dedent()
	out.WriteString("}")
}

func dumpFile(node *File, out *writer.TabbedWriter) {
	out.WriteString("File {")
	out.EndOfLine()
	out.Indent()
	out.WriteString("Decls: []*FunctionDecl {")
	out.EndOfLine()
	out.Indent()
	for i, child := range node.Decls {
		out.WriteString(fmt.Sprintf("%d: ", i))
		dumpFunctionDecl(child, out)
		out.EndOfLine()
	}
	out.Dedent()
	out.WriteString("}")
	out.EndOfLine()
	out.Dedent()
	out.WriteString("}")
}

func dumpTypeName(node *TypeName, out *writer.TabbedWriter) {
	out.WriteString("TypeName {")
	out.EndOfLine()
	out.Indent()
	if node.Name != nil {
		out.WriteString("Name: ")
		dumpToken(node.Name, out)
		out.EndOfLine()
	}
	out.Dedent()
	out.WriteString("}")
}

func dumpTypeRef(node TypeRef, out *writer.TabbedWriter) {
	switch node := node.(type) {
	case *TypeName:
		dumpTypeName(node, out)
	default:
		panic(node)
	}
}
func dumpGetName(node *GetName, out *writer.TabbedWriter) {
	out.WriteString("GetName {")
	out.EndOfLine()
	out.Indent()
	if node.Name != nil {
		out.WriteString("Name: ")
		dumpToken(node.Name, out)
		out.EndOfLine()
	}
	out.Dedent()
	out.WriteString("}")
}

func dumpIntegerLiteral(node *IntegerLiteral, out *writer.TabbedWriter) {
	out.WriteString("IntegerLiteral {")
	out.EndOfLine()
	out.Indent()
	out.WriteString("Pos: ")
	dumpint(node.Pos, out)
	out.EndOfLine()
	out.WriteString("Text: ")
	dumpstring(node.Text, out)
	out.EndOfLine()
	out.Dedent()
	out.WriteString("}")
}

func dumpPrefixExpr(node *PrefixExpr, out *writer.TabbedWriter) {
	out.WriteString("PrefixExpr {")
	out.EndOfLine()
	out.Indent()
	out.WriteString("Op: ")
	dumpPrefixOperator(node.Op, out)
	out.EndOfLine()
	out.WriteString("Expr: ")
	dumpExpr(node.Expr, out)
	out.EndOfLine()
	out.Dedent()
	out.WriteString("}")
}

func dumpInfixExpr(node *InfixExpr, out *writer.TabbedWriter) {
	out.WriteString("InfixExpr {")
	out.EndOfLine()
	out.Indent()
	out.WriteString("Pos: ")
	dumpint(node.Pos, out)
	out.EndOfLine()
	out.WriteString("Left: ")
	dumpExpr(node.Left, out)
	out.EndOfLine()
	out.WriteString("Op: ")
	dumpInfixOperator(node.Op, out)
	out.EndOfLine()
	out.WriteString("Right: ")
	dumpExpr(node.Right, out)
	out.EndOfLine()
	out.Dedent()
	out.WriteString("}")
}

func dumpReturn(node *Return, out *writer.TabbedWriter) {
	out.WriteString("Return {")
	out.EndOfLine()
	out.Indent()
	out.WriteString("Pos: ")
	dumpint(node.Pos, out)
	out.EndOfLine()
	out.WriteString("Expr: ")
	dumpExpr(node.Expr, out)
	out.EndOfLine()
	out.Dedent()
	out.WriteString("}")
}

func dumpCall(node *Call, out *writer.TabbedWriter) {
	out.WriteString("Call {")
	out.EndOfLine()
	out.Indent()
	out.WriteString("Pos: ")
	dumpint(node.Pos, out)
	out.EndOfLine()
	out.WriteString("Expr: ")
	dumpExpr(node.Expr, out)
	out.EndOfLine()
	out.WriteString("Args: []Expr {")
	out.EndOfLine()
	out.Indent()
	for i, child := range node.Args {
		out.WriteString(fmt.Sprintf("%d: ", i))
		dumpExpr(child, out)
		out.EndOfLine()
	}
	out.Dedent()
	out.WriteString("}")
	out.EndOfLine()
	out.Dedent()
	out.WriteString("}")
}

func dumpIf(node *If, out *writer.TabbedWriter) {
	out.WriteString("If {")
	out.EndOfLine()
	out.Indent()
	out.WriteString("Pos: ")
	dumpint(node.Pos, out)
	out.EndOfLine()
	out.WriteString("Cond: ")
	dumpExpr(node.Cond, out)
	out.EndOfLine()
	if node.T != nil {
		out.WriteString("T: ")
		dumpBlock(node.T, out)
		out.EndOfLine()
	}
	if node.F != nil {
		out.WriteString("F: ")
		dumpBlock(node.F, out)
		out.EndOfLine()
	}
	out.Dedent()
	out.WriteString("}")
}

func dumpBlock(node *Block, out *writer.TabbedWriter) {
	out.WriteString("Block {")
	out.EndOfLine()
	out.Indent()
	out.WriteString("Pos: ")
	dumpint(node.Pos, out)
	out.EndOfLine()
	out.WriteString("Exprs: []Expr {")
	out.EndOfLine()
	out.Indent()
	for i, child := range node.Exprs {
		out.WriteString(fmt.Sprintf("%d: ", i))
		dumpExpr(child, out)
		out.EndOfLine()
	}
	out.Dedent()
	out.WriteString("}")
	out.EndOfLine()
	out.Dedent()
	out.WriteString("}")
}

func dumpFunctionDecl(node *FunctionDecl, out *writer.TabbedWriter) {
	out.WriteString("FunctionDecl {")
	out.EndOfLine()
	out.Indent()
	if node.Name != nil {
		out.WriteString("Name: ")
		dumpToken(node.Name, out)
		out.EndOfLine()
	}
	out.WriteString("Params: []*Param {")
	out.EndOfLine()
	out.Indent()
	for i, child := range node.Params {
		out.WriteString(fmt.Sprintf("%d: ", i))
		dumpParam(child, out)
		out.EndOfLine()
	}
	out.Dedent()
	out.WriteString("}")
	out.EndOfLine()
	out.WriteString("Results: []*Param {")
	out.EndOfLine()
	out.Indent()
	for i, child := range node.Results {
		out.WriteString(fmt.Sprintf("%d: ", i))
		dumpParam(child, out)
		out.EndOfLine()
	}
	out.Dedent()
	out.WriteString("}")
	out.EndOfLine()
	if node.Body != nil {
		out.WriteString("Body: ")
		dumpBlock(node.Body, out)
		out.EndOfLine()
	}
	out.Dedent()
	out.WriteString("}")
}

func dumpExpr(node Expr, out *writer.TabbedWriter) {
	switch node := node.(type) {
	case *GetName:
		dumpGetName(node, out)
	case *IntegerLiteral:
		dumpIntegerLiteral(node, out)
	case *PrefixExpr:
		dumpPrefixExpr(node, out)
	case *InfixExpr:
		dumpInfixExpr(node, out)
	case *Return:
		dumpReturn(node, out)
	case *Call:
		dumpCall(node, out)
	case *If:
		dumpIf(node, out)
	case *Block:
		dumpBlock(node, out)
	case *FunctionDecl:
		dumpFunctionDecl(node, out)
	default:
		panic(node)
	}
}
