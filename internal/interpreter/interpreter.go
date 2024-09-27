package interpreter

import (
	"fmt"
	"strconv"

	"github.com/WhoDoIt/GoCompiler/internal/syntaxtree"
	"github.com/WhoDoIt/GoCompiler/internal/tokenizer"
)

type scope struct {
	vars   map[string]WorkingType
	parent *scope
}

func (s *scope) GetValue(name string) WorkingType {
	if val, ok := s.vars[name]; ok {
		return val
	} else {
		if s.parent != nil {
			return s.parent.GetValue(name)
		} else {
			return nil
		}
	}
}

func (s *scope) AssignValue(name string, val WorkingType) {
	if _, ok := s.vars[name]; ok {
		s.vars[name] = val
	} else {
		if s.parent != nil {
			s.parent.AssignValue(name, val)
		} else {
			return
		}
	}
}

func (s *scope) DeclareValue(name string, val WorkingType) {
	s.vars[name] = val
}

type intepreter struct {
	sc scope
}

func Evaluate(stmt []syntaxtree.Stmt) {
	program := intepreter{sc: scope{vars: make(map[string]WorkingType), parent: nil}}
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
	case tokenizer.EQUAL_EQUAL:
		return s.number(expr.Left).operatorEqual(s.number(expr.Right))
	case tokenizer.EXCLAMATION_EQUAL:
		return s.number(expr.Left).operatorEqual(s.number(expr.Right)).operatorExclamation()
	case tokenizer.LESS:
		return s.number(expr.Left).operatorLess(s.number(expr.Right))
	case tokenizer.LESS_EQUAL:
		left := s.number(expr.Left)
		right := s.number(expr.Right)
		result := Bool{value: left.operatorLess(right).value || left.operatorEqual(right).value}
		return result
	case tokenizer.GREATER:
		left := s.number(expr.Left)
		right := s.number(expr.Right)
		result := Bool{value: left.operatorLess(right).value || left.operatorEqual(right).value}
		return result.operatorExclamation()
	case tokenizer.GREATER_EQUAL:
		return s.number(expr.Left).operatorLess(s.number(expr.Right)).operatorExclamation()
	case tokenizer.EQUAL:
		val := s.number(expr.Right)
		s.sc.AssignValue(expr.Left.(syntaxtree.LiteralExpr).Value.Content, val)
		return val
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
		return s.sc.GetValue(expr.Value.Content)
	}
	if expr.Value.Type == tokenizer.STRING {
		return String{value: expr.Value.Content}
	} else {
		res, _ := strconv.Atoi(expr.Value.Content)
		return Int{value: int(res)}
	}
}

func (s *intepreter) VisitExpressionStmt(stmt syntaxtree.ExpressionStmt) any {
	s.number(stmt.Expression)
	return nil
}
func (s *intepreter) VisitPrintStmt(stmt syntaxtree.PrintStmt) any {
	value := syntaxtree.AcceptExpr(s, stmt.Expression)
	fmt.Println("> ", value)
	return nil
}
func (s *intepreter) VisitVarDeclStmt(stmt syntaxtree.VarDeclStmt) any {
	value := syntaxtree.AcceptExpr(s, stmt.Expression)
	s.sc.DeclareValue(stmt.Name.Content, value)
	return nil
}

func (s *intepreter) VisitForStmt(stmt syntaxtree.ForStmt) any {
	for s.eval(stmt.PreStatement); s.number(stmt.Condition).(Bool).value; s.number(stmt.PostStatement) {
		s.eval(stmt.Block)
	}
	return nil
}

func (s *intepreter) VisitIfStmt(stmt syntaxtree.IfStmt) any {
	cond := s.number(stmt.Condition)
	if val, ok := cond.(Bool); ok {
		if val.value {
			s.eval(stmt.Block)
		}
	}
	return nil
}

func (s *intepreter) VisitBlockStmt(stmt syntaxtree.BlockStmt) any {
	prev := s.sc
	s.sc = scope{vars: make(map[string]WorkingType), parent: &prev}
	for _, v := range stmt.Statements {
		s.eval(v)
	}
	s.sc = prev
	return nil
}
