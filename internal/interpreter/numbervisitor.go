package interpreter

import (
	"strconv"

	"github.com/WhoDoIt/GoCompiler/internal/syntaxtree"
	"github.com/WhoDoIt/GoCompiler/internal/tokenizer"
)

type NumberEvalVisitor struct {
}

func (s NumberEvalVisitor) Calculate(expr syntaxtree.Expr) int {
	return syntaxtree.AcceptExpr(s, expr)
}

func (s NumberEvalVisitor) number(expr syntaxtree.Expr) int {
	return syntaxtree.AcceptExpr(s, expr)
}

func (s NumberEvalVisitor) VisitBinaryExpr(expr syntaxtree.BinaryExpr) int {
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

func (s NumberEvalVisitor) VisitUnaryExpr(expr syntaxtree.UnaryExpr) int {
	switch expr.Operator.Type {
	case tokenizer.EXCLAMATION:
		return s.number(expr.Right)
	case tokenizer.MINUS:
		return -s.number(expr.Right)
	}
	return 0
}

func (s NumberEvalVisitor) VisitGroupingExpr(expr syntaxtree.GroupingExpr) int {
	return s.number(expr.Inside)
}

func (s NumberEvalVisitor) VisitLiteral(expr syntaxtree.LiteralExpr) int {
	res, _ := strconv.Atoi(expr.Value.Content)
	return int(res)
}
