package ast

type Pos struct {
	Line   int
	Offset int
}

func NewPos(line, offset int) *Pos {
	return &Pos{Line: line, Offset: offset}
}

type Ast struct {
	pos *Pos
}

func (ast *Ast) Pos() *Pos {
	return ast.pos
}

func (ast *Ast) SetPos(pos *Pos) {
	ast.pos = pos
}

type Expr interface {
	Pos() *Pos
	SetPos(*Pos)
}

type RefvarExpr struct {
	Ast
	Name string
}

type NilExpr struct {
	Ast
}

type TrueExpr struct {
	Ast
}

type FalseExpr struct {
	Ast
}

type FixnumExpr struct {
	Ast
	Value string
}

type FloatnumExpr struct {
	Ast
	Value string
}

type StringExpr struct {
	Ast
	Value string
}

type FuncallExpr struct {
	Ast
	Fun  Expr
	Args []Expr
}

type FunctionExpr struct {
	Ast
	Parameters []string
	Body       *Block
}

type UnaryExpr struct {
	Ast
	Op    string
	Value Expr
}

type BinaryExpr struct {
	Ast
	Op    string
	Left  Expr
	Right Expr
}

type AndExpr struct {
	Ast
	Left  Expr
	Right Expr
}

type OrExpr struct {
	Ast
	Left  Expr
	Right Expr
}

type IndexExpr struct {
	Ast
	Prefix Expr
	Index  Expr
}

type SliceExpr struct {
	Ast
	Prefix Expr
	Begin  Expr
	End    Expr
}

type ListExpr struct {
	Ast
	Elements []Expr
}

type PairExpr struct {
	Ast
	Key   Expr
	Value Expr
}

type TableExpr struct {
	Ast
	Elements []Expr
}

type RangeExpr struct {
	Ast
	Start Expr
	End   Expr
	Step  Expr
}

type Stmt interface {
	Pos() *Pos
	SetPos(*Pos)
}

type VarStmt struct {
	Ast
	Vars []*VarDeclStmt
}

type VarDeclStmt struct {
	Ast
	Name  string
	Value Expr
}

type AssignStmt struct {
	Ast
	Lhs Expr
	Rhs Expr
}

type IfStmt struct {
	Ast
	Test Expr
	Then Stmt
	Else Stmt
}

type ForStmt struct {
	Ast
	Name string
	Exp  Expr
	Body *Block
}

type WhileStmt struct {
	Ast
	Test Expr
	Body *Block
}

type ReturnStmt struct {
	Ast
	Value Expr
}

type BreakStmt struct {
	Ast
}

type ContinueStmt struct {
	Ast
}

type ExprStmt struct {
	Ast
	Exp Expr
}

type Block struct {
	Ast
	Stmts []Stmt
}
