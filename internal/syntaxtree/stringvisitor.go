package syntaxtree

type StringVisitor struct {
}

func (s StringVisitor) Print(expr Expr) string {
	return Accept(s, expr)
}

func (s StringVisitor) string(name string, expr []Expr) string {
	var result = ""
	result += "(" + name
	for _, v := range expr {
		result += " "
		result += Accept(s, v)
	}
	result += ")"
	return result
}

func (s StringVisitor) VisitBinaryExpr(expr BinaryExpr) string {
	return s.string(expr.Operator.Content, []Expr{expr.Left, expr.Right})
}

func (s StringVisitor) VisitUnaryExpr(expr UnaryExpr) string {
	return s.string(expr.Operator.Content, []Expr{expr.Right})
}

func (s StringVisitor) VisitGroupingExpr(expr GroupingExpr) string {
	return s.string("group", []Expr{expr.Inside})
}

func (s StringVisitor) VisitLiteral(expr Literal) string {
	return expr.Value.Content
}
