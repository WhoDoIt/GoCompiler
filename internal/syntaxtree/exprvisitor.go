package syntaxtree

type ExprVisitor[E any] interface {
	VisitUnaryExpr(expr UnaryExpr) E
	VisitBinaryExpr(expr BinaryExpr) E
	VisitGroupingExpr(expr GroupingExpr) E
	VisitLiteral(expr LiteralExpr) E
}

func AcceptExpr[E any](visitor ExprVisitor[E], expr Expr) E {
	switch val := expr.(type) {
	case BinaryExpr:
		return visitor.VisitBinaryExpr(val)
	case UnaryExpr:
		return visitor.VisitUnaryExpr(val)
	case GroupingExpr:
		return visitor.VisitGroupingExpr(val)
	case LiteralExpr:
		return visitor.VisitLiteral(val)
	}
	return *new(E)
}
