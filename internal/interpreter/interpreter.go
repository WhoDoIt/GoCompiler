package interpreter

import (
	"fmt"
	"strconv"

	"github.com/WhoDoIt/GoCompiler/internal/syntaxtree"
	"github.com/WhoDoIt/GoCompiler/internal/tokenizer"
)

type scope struct {
	vars   map[string]RoseType
	parent *scope
}

func (s *scope) GetValue(name string) RoseType {
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

func (s *scope) AssignValue(name string, val RoseType) {
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

func (s *scope) DeclareValue(name string, val RoseType) {
	s.vars[name] = val
}

type intepreter struct {
	sc scope
}

func Evaluate(stmt []syntaxtree.Stmt) {
	program := intepreter{sc: scope{vars: make(map[string]RoseType), parent: nil}}
	for _, v := range stmt {
		program.eval(v)
	}
}

func (s *intepreter) eval(stmt syntaxtree.Stmt) {
	syntaxtree.AcceptStmt(s, stmt)
}

func (s *intepreter) number(expr syntaxtree.Expr) RoseType {
	res := syntaxtree.AcceptExpr(s, expr)
	if cast, ok := res.(RuntimeError); ok {
		fmt.Println("RUNTIME ERROR: " + cast.value)
	}
	return res
}

func (s *intepreter) VisitBinaryExpr(expr syntaxtree.BinaryExpr) RoseType {
	switch expr.Operator.Type {
	case tokenizer.LESS_EQUAL:
		left := s.number(expr.Left)
		right := s.number(expr.Right)
		if val, ok := left.operatorBinary(tokenizer.LESS, right).(RuntimeError); ok {
			return val
		}
		if val, ok := left.operatorBinary(tokenizer.EQUAL_EQUAL, right).(RuntimeError); ok {
			return val
		}
		result := RoseBool{value: left.operatorBinary(tokenizer.EQUAL_EQUAL, right).(RoseBool).value || left.operatorBinary(tokenizer.LESS_EQUAL, right).(RoseBool).value}
		return result
	case tokenizer.GREATER:
		left := s.number(expr.Left)
		right := s.number(expr.Right)
		if val, ok := left.operatorBinary(tokenizer.LESS, right).(RuntimeError); ok {
			return val
		}
		if val, ok := left.operatorBinary(tokenizer.EQUAL_EQUAL, right).(RuntimeError); ok {
			return val
		}
		result := RoseBool{value: left.operatorBinary(tokenizer.EQUAL_EQUAL, right).(RoseBool).value || left.operatorBinary(tokenizer.LESS_EQUAL, right).(RoseBool).value}
		return result.operatorUnary(tokenizer.EXCLAMATION)
	case tokenizer.GREATER_EQUAL:
		return s.number(expr.Left).operatorBinary(tokenizer.LESS, s.number(expr.Right)).operatorUnary(tokenizer.EXCLAMATION)
	case tokenizer.EQUAL:
		val := s.number(expr.Right)
		s.sc.AssignValue(expr.Left.(syntaxtree.LiteralExpr).Value.Content, val)
		return val
	default:
		return s.number(expr.Left).operatorBinary(expr.Operator.Type, s.number(expr.Right))
	}
}

func (s *intepreter) VisitUnaryExpr(expr syntaxtree.UnaryExpr) RoseType {
	return s.number(expr.Right).operatorUnary(expr.Operator.Type)
}

func (s *intepreter) VisitGroupingExpr(expr syntaxtree.GroupingExpr) RoseType {
	return s.number(expr.Inside)
}

func (s *intepreter) VisitLiteralExpr(expr syntaxtree.LiteralExpr) RoseType {
	if expr.Value.Type == tokenizer.IDENTIFIER {
		return s.sc.GetValue(expr.Value.Content)
	}
	if expr.Value.Type == tokenizer.STRING {
		return RoseString{value: expr.Value.Content}
	} else {
		res, _ := strconv.Atoi(expr.Value.Content)
		return RoseInt{value: int(res)}
	}
}

func (s *intepreter) VisitCallExpr(expr syntaxtree.CallExpr) RoseType {
	panic("BROTHER")
	return nil
}

func (s *intepreter) VisitExpressionStmt(stmt syntaxtree.ExpressionStmt) any {
	s.number(stmt.Expression)
	return nil
}
func (s *intepreter) VisitPrintStmt(stmt syntaxtree.PrintStmt) any {
	value := s.number(stmt.Expression)
	fmt.Println("> ", value)
	return nil
}
func (s *intepreter) VisitVarDeclStmt(stmt syntaxtree.VarDeclStmt) any {
	value := s.number(stmt.Expression)
	s.sc.DeclareValue(stmt.Name.Content, value)
	return nil
}

func (s *intepreter) VisitForStmt(stmt syntaxtree.ForStmt) any {
	for s.eval(stmt.PreStatement); s.number(stmt.Condition).(RoseBool).value; s.number(stmt.PostStatement) {
		s.eval(stmt.Block)
	}
	return nil
}

func (s *intepreter) VisitIfStmt(stmt syntaxtree.IfStmt) any {
	cond := s.number(stmt.Condition)
	if val, ok := cond.(RoseBool); ok {
		if val.value {
			s.eval(stmt.Block)
		}
	}
	return nil
}

func (s *intepreter) VisitBlockStmt(stmt syntaxtree.BlockStmt) any {
	prev := s.sc
	s.sc = scope{vars: make(map[string]RoseType), parent: &prev}
	for _, v := range stmt.Statements {
		s.eval(v)
	}
	s.sc = prev
	return nil
}
