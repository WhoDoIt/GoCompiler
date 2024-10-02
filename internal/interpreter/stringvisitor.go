package interpreter

import (
	"github.com/WhoDoIt/GoCompiler/internal/syntaxtree"
)

type StringVisitor struct {
}

func (s StringVisitor) Print(expr syntaxtree.Expr) string {
	return syntaxtree.AcceptExpr(s, expr)
}

func (s StringVisitor) string(name string, expr []syntaxtree.Expr) string {
	var result = ""
	result += "(" + name
	for _, v := range expr {
		result += " "
		result += syntaxtree.AcceptExpr(s, v)
	}
	result += ")"
	return result
}

func (s StringVisitor) VisitBinaryExpr(expr syntaxtree.BinaryExpr) string {
	return s.string(expr.Operator.Content, []syntaxtree.Expr{expr.Left, expr.Right})
}

func (s StringVisitor) VisitUnaryExpr(expr syntaxtree.UnaryExpr) string {
	return s.string(expr.Operator.Content, []syntaxtree.Expr{expr.Right})
}

func (s StringVisitor) VisitGroupingExpr(expr syntaxtree.GroupingExpr) string {
	return s.string("group", []syntaxtree.Expr{expr.Inside})
}

func (s StringVisitor) VisitLiteralExpr(expr syntaxtree.LiteralExpr) string {
	return expr.Value.Content
}

func (s StringVisitor) VisitCallExpr(expr syntaxtree.CallExpr) string {
	return s.string("call $"+s.Print(expr.Calle), expr.Arguments)
}
