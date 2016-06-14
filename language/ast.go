package language

/* Generated from regenerate_cacao, do not edit by hand. */

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
