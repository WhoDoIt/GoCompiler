package syntaxtree

import (
	"strconv"

	"github.com/WhoDoIt/GoCompiler/internal/tokenizer"
)

type NumberEvalVisitor struct {
}

func (s NumberEvalVisitor) Calculate(expr Expr) int {
	return Accept(s, expr)
}

func (s NumberEvalVisitor) number(expr Expr) int {
	return Accept(s, expr)
}

func (s NumberEvalVisitor) VisitBinaryExpr(expr BinaryExpr) int {
	switch expr.Operator.Type {
	case tokenizer.PLUS:
		return s.number(expr.Left) + s.number(expr.Right)
	case tokenizer.MINUS:
		return s.number(expr.Left) - s.number(expr.Right)
	case tokenizer.SLASH:
		return s.number(expr.Left) / s.number(expr.Right)
	case tokenizer.STAR:
		return s.number(expr.Left) * s.number(expr.Right)
	case tokenizer.PIPE:
		return s.number(expr.Left) | s.number(expr.Right)
	case tokenizer.AMPERSAND:
		return s.number(expr.Left) & s.number(expr.Right)
	}
	return 0
}

func (s NumberEvalVisitor) VisitUnaryExpr(expr UnaryExpr) int {
	switch expr.Operator.Type {
	case tokenizer.EXCLAMATION:
		return s.number(expr.Right)
	case tokenizer.MINUS:
		return -s.number(expr.Right)
	}
	return 0
}

func (s NumberEvalVisitor) VisitGroupingExpr(expr GroupingExpr) int {
	return s.number(expr.Inside)
}

func (s NumberEvalVisitor) VisitLiteral(expr Literal) int {
	res, _ := strconv.Atoi(expr.Value.Content)
	return int(res)
}
