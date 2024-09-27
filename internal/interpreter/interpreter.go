package interpreter

import (
	"fmt"
	"strconv"

	"github.com/WhoDoIt/GoCompiler/internal/syntaxtree"
	"github.com/WhoDoIt/GoCompiler/internal/tokenizer"
)

type intepreter struct {
	vars map[string]WorkingType
}

func Evaluate(stmt []syntaxtree.Stmt) {
	program := intepreter{vars: make(map[string]WorkingType)}
	for _, v := range stmt {
		program.eval(v)
	}
}

func (s *intepreter) eval(stmt syntaxtree.Stmt) {
	syntaxtree.AcceptStmt(s, stmt)
}

func (s *intepreter) number(expr syntaxtree.Expr) WorkingType {
	return syntaxtree.AcceptExpr(s, expr)
}

func (s *intepreter) VisitBinaryExpr(expr syntaxtree.BinaryExpr) WorkingType {
	switch expr.Operator.Type {
	case tokenizer.PLUS:
		return s.number(expr.Left).operatorPlus(s.number(expr.Right))
	case tokenizer.MINUS:
		return s.number(expr.Left).operatorMinus(s.number(expr.Right))
	case tokenizer.SLASH:
		return s.number(expr.Left).operatorSlash(s.number(expr.Right))
	case tokenizer.STAR:
		return s.number(expr.Left).operatorStar(s.number(expr.Right))
		// case tokenizer.PIPE:
		// 	return s.number(expr.Left) | s.number(expr.Right)
		// case tokenizer.AMPERSAND:
		// 	return s.number(expr.Left) & s.number(expr.Right)
	}
	return s.number(expr.Left).zeroValue()
}

func (s *intepreter) VisitUnaryExpr(expr syntaxtree.UnaryExpr) WorkingType {
	switch expr.Operator.Type {
	case tokenizer.EXCLAMATION:
		return s.number(expr.Right)
	case tokenizer.MINUS:
		return s.number(expr.Right).operatorSelfminus()
	}
	return s.number(expr.Right).zeroValue()
}

func (s *intepreter) VisitGroupingExpr(expr syntaxtree.GroupingExpr) WorkingType {
	return s.number(expr.Inside)
}

func (s *intepreter) VisitLiteral(expr syntaxtree.LiteralExpr) WorkingType {
	if expr.Value.Type == tokenizer.IDENTIFIER {
		return s.vars[expr.Value.Content]
	}
	if expr.Value.Type == tokenizer.STRING {
		return String{value: expr.Value.Content}
	} else {
		res, _ := strconv.Atoi(expr.Value.Content)
		return Int{value: int(res)}
	}
}

func (s *intepreter) VisitExpressionStmt(stmt syntaxtree.ExpressionStmt) any {
	return nil
}
func (s *intepreter) VisitPrintStmt(stmt syntaxtree.PrintStmt) any {
	value := syntaxtree.AcceptExpr(s, stmt.Expression)
	fmt.Println("> ", value)
	return nil
}
func (s *intepreter) VisitVarDeclStmt(stmt syntaxtree.VarDeclStmt) any {
	value := syntaxtree.AcceptExpr(s, stmt.Expression)
	// if _, ok := s.vars[stmt.Name.Content]; ok {

	// }
	s.vars[stmt.Name.Content] = value
	return nil
}
