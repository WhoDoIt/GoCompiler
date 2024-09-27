package syntaxtree

import (
	"github.com/WhoDoIt/GoCompiler/internal/tokenizer"
)

type Expr interface {
}

// type AssignExpr struct {
// 	Left  tokenizer.Token
// 	Right Expr
// }

type BinaryExpr struct {
	Left     Expr
	Operator tokenizer.Token
	Right    Expr
}

type UnaryExpr struct {
	Operator tokenizer.Token
	Right    Expr
}

type GroupingExpr struct {
	Inside Expr
}

type LiteralExpr struct {
	Value tokenizer.Token
}

type VariableExpr struct {
	Value tokenizer.Token
}

type Stmt interface {
}

type ExpressionStmt struct {
	Expression Expr
}

type PrintStmt struct {
	Expression Expr
}

type IfStmt struct {
	Condition Expr
	Block     Stmt
}

type BlockStmt struct {
	Statements []Stmt
}

type VarDeclStmt struct {
	Name       tokenizer.Token
	Expression Expr
}

type ForStmt struct {
	PreStatement  Stmt
	Condition     Expr
	PostStatement Expr
	Block         Stmt
}
