package syntaxtree

type StmtVisitor[E any] interface {
	VisitExpressionStmt(stmt ExpressionStmt) E
	VisitPrintStmt(stmt PrintStmt) E
	VisitBlockStmt(stmt BlockStmt) E
	VisitIfStmt(stmt IfStmt) E
	VisitVarDeclStmt(stmt VarDeclStmt) E
	VisitForStmt(stmt ForStmt) E
}

func AcceptStmt[E any](visitor StmtVisitor[E], expr Stmt) E {
	switch val := expr.(type) {
	case ExpressionStmt:
		return visitor.VisitExpressionStmt(val)
	case PrintStmt:
		return visitor.VisitPrintStmt(val)
	case VarDeclStmt:
		return visitor.VisitVarDeclStmt(val)
	case IfStmt:
		return visitor.VisitIfStmt(val)
	case BlockStmt:
		return visitor.VisitBlockStmt(val)
	case ForStmt:
		return visitor.VisitForStmt(val)
	}
	return *new(E)
}
