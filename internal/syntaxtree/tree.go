package syntaxtree

import (
	"github.com/WhoDoIt/GoCompiler/internal/tokenizer"
)

type Expr interface {
}

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

type Stmt interface {
}

type ExpressionStmt struct {
	Expression Expr
}

type PrintStmt struct {
	Expression Expr
}

type VarDeclStmt struct {
	Name       tokenizer.Token
	Expression Expr
}
