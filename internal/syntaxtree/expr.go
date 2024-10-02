package syntaxtree

import (
	"github.com/WhoDoIt/GoCompiler/internal/tokenizer"
)

type Expr interface{}
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
type CallExpr struct {
	Calle     Expr
	Paren     tokenizer.Token
	Arguments []Expr
}
type LiteralExpr struct {
	Value tokenizer.Token
}
type ExprVisitor[E any] interface {
	VisitBinaryExpr(expr BinaryExpr) E
	VisitUnaryExpr(expr UnaryExpr) E
	VisitGroupingExpr(expr GroupingExpr) E
	VisitCallExpr(expr CallExpr) E
	VisitLiteralExpr(expr LiteralExpr) E
}

func AcceptExpr[E any](visitor ExprVisitor[E], expr Expr) E {
	switch val := expr.(type) {
	case BinaryExpr:
		return visitor.VisitBinaryExpr(val)
	case UnaryExpr:
		return visitor.VisitUnaryExpr(val)
	case GroupingExpr:
		return visitor.VisitGroupingExpr(val)
	case CallExpr:
		return visitor.VisitCallExpr(val)
	case LiteralExpr:
		return visitor.VisitLiteralExpr(val)
	}
	return *new(E)
}
