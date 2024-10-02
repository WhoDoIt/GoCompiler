package syntaxtree

import (
	"github.com/WhoDoIt/GoCompiler/internal/tokenizer"
)

type Stmt interface{}
type ExpressionStmt struct {
	Expression Expr
}
type PrintStmt struct {
	Expression Expr
}
type BlockStmt struct {
	Statements []Stmt
}
type IfStmt struct {
	Condition Expr
	Block     Stmt
}
type VarDeclStmt struct {
	Name       tokenizer.Token
	Expression Expr
}
type ForStmt struct {
	PreStatement  Stmt
	Condition     Expr
	PostStatement Expr
	Block         Stmt
}
type StmtVisitor[E any] interface {
	VisitExpressionStmt(stmt ExpressionStmt) E
	VisitPrintStmt(stmt PrintStmt) E
	VisitBlockStmt(stmt BlockStmt) E
	VisitIfStmt(stmt IfStmt) E
	VisitVarDeclStmt(stmt VarDeclStmt) E
	VisitForStmt(stmt ForStmt) E
}

func AcceptStmt[E any](visitor StmtVisitor[E], stmt Stmt) E {
	switch val := stmt.(type) {
	case ExpressionStmt:
		return visitor.VisitExpressionStmt(val)
	case PrintStmt:
		return visitor.VisitPrintStmt(val)
	case BlockStmt:
		return visitor.VisitBlockStmt(val)
	case IfStmt:
		return visitor.VisitIfStmt(val)
	case VarDeclStmt:
		return visitor.VisitVarDeclStmt(val)
	case ForStmt:
		return visitor.VisitForStmt(val)
	}
	return *new(E)
}
