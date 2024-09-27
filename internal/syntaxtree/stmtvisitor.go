package syntaxtree

type StmtVisitor[E any] interface {
	VisitExpressionStmt(stmt ExpressionStmt) E
	VisitPrintStmt(stmt PrintStmt) E
	VisitVarDeclStmt(stmt VarDeclStmt) E
}

func AcceptStmt[E any](visitor StmtVisitor[E], expr Stmt) E {
	switch val := expr.(type) {
	case ExpressionStmt:
		return visitor.VisitExpressionStmt(val)
	case PrintStmt:
		return visitor.VisitPrintStmt(val)
	case VarDeclStmt:
		return visitor.VisitVarDeclStmt(val)
	}
	return *new(E)
}
