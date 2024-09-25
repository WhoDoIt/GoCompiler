package syntaxtree

import (
	"reflect"

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

type Literal struct {
	Value tokenizer.Token
}

type Visitor[E any] interface {
	VisitUnaryExpr(expr UnaryExpr) E
	VisitBinaryExpr(expr BinaryExpr) E
	VisitGroupingExpr(expr GroupingExpr) E
	VisitLiteral(expr Literal) E
}

func Accept[E any](visitor Visitor[E], expr Expr) E {
	switch reflect.TypeOf(expr).Name() {
	case "BinaryExpr":
		return visitor.VisitBinaryExpr(expr.(BinaryExpr))
	case "UnaryExpr":
		return visitor.VisitUnaryExpr(expr.(UnaryExpr))
	case "GroupingExpr":
		return visitor.VisitGroupingExpr(expr.(GroupingExpr))
	case "Literal":
		return visitor.VisitLiteral(expr.(Literal))
	}
	return *new(E)
}
