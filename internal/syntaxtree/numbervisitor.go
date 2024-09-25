package syntaxtree

import (
	"strconv"

	"github.com/WhoDoIt/GoCompiler/internal/tokenizer"
)

type NumberEvalVisitor struct {
}

func (s NumberEvalVisitor) Calculate(expr Expr) float32 {
	return Accept(s, expr)
}

func (s NumberEvalVisitor) number(expr Expr) float32 {
	return Accept(s, expr)
}

func (s NumberEvalVisitor) VisitBinaryExpr(expr BinaryExpr) float32 {
	switch expr.Operator.Type {
	case tokenizer.PLUS:
		return s.number(expr.Left) + s.number(expr.Right)
	case tokenizer.MINUS:
		return s.number(expr.Left) - s.number(expr.Right)
	case tokenizer.SLASH:
		return s.number(expr.Left) / s.number(expr.Right)
	case tokenizer.STAR:
		return s.number(expr.Left) * s.number(expr.Right)
	}
	return 0
}

func (s NumberEvalVisitor) VisitUnaryExpr(expr UnaryExpr) float32 {
	switch expr.Operator.Type {
	case tokenizer.EXCLAMATION:
		return s.number(expr.Right)
	case tokenizer.MINUS:
		return -s.number(expr.Right)
	}
	return 0
}

func (s NumberEvalVisitor) VisitGroupingExpr(expr GroupingExpr) float32 {
	return s.number(expr.Inside)
}

func (s NumberEvalVisitor) VisitLiteral(expr Literal) float32 {
	res, _ := strconv.Atoi(expr.Value.Content)
	return float32(res)
}
