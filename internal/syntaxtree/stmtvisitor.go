package syntaxtree

import "reflect"

type StmtVisitor[E any] interface {
	VisitExpressionStmt(stmt ExpressionStmt) E
	VisitPrintStmt(stmt PrintStmt) E
	VisitVarDeclStmt(stmt VarDeclStmt) E
}

func AcceptStmt[E any](visitor StmtVisitor[E], expr int) E {
	switch reflect.TypeOf(expr).Name() {
	// case "BinaryExpr":
	// 	return visitor.VisitBinaryExpr(expr.(BinaryExpr))
	// case "UnaryExpr":
	// 	return visitor.VisitUnaryExpr(expr.(UnaryExpr))
	// case "GroupingExpr":
	// 	return visitor.VisitGroupingExpr(expr.(GroupingExpr))
	// case "Literal":
	// 	return visitor.VisitLiteral(expr.(Literal))
	}
	return *new(E)
}
