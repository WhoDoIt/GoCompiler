package parser

import (
	"errors"
	"fmt"
	"strconv"

	"github.com/WhoDoIt/GoCompiler/internal/syntaxtree"
	"github.com/WhoDoIt/GoCompiler/internal/tokenizer"
)

type parser struct {
	tokens  []tokenizer.Token
	current int
}

func (p *parser) isAtEnd() bool {
	return p.tokens[p.current].Type == tokenizer.EOF
}

func (p *parser) previus() tokenizer.Token {
	return p.tokens[p.current-1]
}

func (p *parser) peek() tokenizer.Token {
	return p.tokens[p.current]
}

func (p *parser) advance() tokenizer.Token {
	if !p.isAtEnd() {
		p.current++
	}
	return p.previus()
}

func (p *parser) checkMany(tokenType []tokenizer.TokenType) bool {
	if p.isAtEnd() {
		return false
	}
	for _, v := range tokenType {
		if v == p.peek().Type {
			return true
		}
	}
	return false
}

func (p *parser) check(tokenType tokenizer.TokenType) bool {
	if p.isAtEnd() {
		return false
	}
	if tokenType == p.peek().Type {
		return true
	}

	return false
}

func (p *parser) generateError(str string) error {
	if p.current == 0 {
		return errors.New("[" + str + "]" + " on line " + strconv.Itoa(p.peek().Line))
	} else {
		return errors.New("[" + str + "]" + " on line " + strconv.Itoa(p.previus().Line))
	}
}

func (p *parser) synchronize() {
	p.advance()

	for !p.isAtEnd() {
		if p.previus().Type == tokenizer.SEMICOLON {
			return
		}
		switch p.peek().Type {
		case tokenizer.FN, tokenizer.VAR, tokenizer.IF, tokenizer.ELSE, tokenizer.RETURN, tokenizer.FOR, tokenizer.STRUCT, tokenizer.PRINT:
			return
		}
		p.advance()
	}
}

func Parse(tokens []tokenizer.Token) ([]syntaxtree.Stmt, error) {
	parser := parser{tokens: tokens}
	var tree []syntaxtree.Stmt
	var errs []error
	for !parser.isAtEnd() {
		stmt, err := parser.declaration()
		if err != nil {
			errs = append(errs, err)
		}
		tree = append(tree, stmt)
	}
	for _, v := range errs {
		fmt.Println(v.Error())
	}
	if errs != nil {
		return nil, errors.New("got parser error")
	} else {
		return tree, nil
	}
}

func (p *parser) declaration() (syntaxtree.Stmt, error) {
	var stmt syntaxtree.Stmt
	var err error
	if p.check(tokenizer.VAR) {
		stmt, err = p.varDelc()
	} else {
		stmt, err = p.statement()
	}
	if err != nil {
		p.synchronize()
		return nil, err
	} else {
		return stmt, nil
	}
}

func (p *parser) varDelc() (syntaxtree.Stmt, error) {
	// ASSUME VAR ALREADY CHECKED
	if !p.check(tokenizer.VAR) {
		return nil, p.generateError("expected var")
	}
	p.advance()
	name := p.advance()
	if name.Type != tokenizer.IDENTIFIER {
		return nil, p.generateError("bad name for var")
	}
	if !p.check(tokenizer.EQUAL) {
		return nil, p.generateError("expected =")
	}
	p.advance()
	expr, err := p.expression()
	if err != nil {
		return nil, err
	}
	if !p.check(tokenizer.SEMICOLON) {
		return nil, p.generateError("expected ;")
	}
	p.advance()
	return syntaxtree.VarDeclStmt{Name: name, Expression: expr}, nil

}

func (p *parser) statement() (syntaxtree.Stmt, error) {
	if p.check(tokenizer.PRINT) {
		return p.printStmt()
	} else if p.check(tokenizer.IF) {
		return p.ifStmt()
	} else if p.check(tokenizer.LEFT_BRACE) {
		return p.block()
	} else if p.check(tokenizer.FOR) {
		return p.forStmt()
	} else {
		return p.exprStmt()
	}
}

func (p *parser) forStmt() (syntaxtree.Stmt, error) {
	p.advance()
	if !p.check(tokenizer.LEFT_PAREN) {
		return nil, p.generateError("expected ( after for")
	}
	p.advance()
	prestmt, err := p.varDelc()
	if err != nil {
		return nil, err
	}
	// if !p.check(tokenizer.SEMICOLON) {
	// 	return nil, p.generateError("expected ; after PreStatement")
	// }
	// p.advance()
	cond, err := p.expression()
	if err != nil {
		return nil, err
	}
	if !p.check(tokenizer.SEMICOLON) {
		return nil, p.generateError("expected ; after condition")
	}
	p.advance()
	poststmt, err := p.expression()
	if err != nil {
		return nil, err
	}

	if !p.check(tokenizer.RIGHT_PAREN) {
		return nil, p.generateError("expected ) after for block")
	}
	p.advance()

	block, err := p.statement()
	if err != nil {
		return nil, err
	}
	return syntaxtree.ForStmt{PreStatement: prestmt, Condition: cond, PostStatement: poststmt, Block: block}, nil
}

func (p *parser) ifStmt() (syntaxtree.Stmt, error) {
	p.advance()
	if !p.check(tokenizer.LEFT_PAREN) {
		return nil, p.generateError("expected ( after if")
	}
	p.advance()
	expr, err := p.expression()
	if err != nil {
		return nil, err
	}
	if !p.check(tokenizer.RIGHT_PAREN) {
		return nil, p.generateError("expected ) after if statement")
	}
	p.advance()
	if !p.check(tokenizer.LEFT_BRACE) {
		return nil, p.generateError("expected { after if statement")
	}
	block, err := p.statement()
	if err != nil {
		return nil, err
	}
	return syntaxtree.IfStmt{Condition: expr, Block: block}, nil
}

func (p *parser) block() (syntaxtree.Stmt, error) {
	p.advance()
	var stmt []syntaxtree.Stmt
	for !p.check(tokenizer.RIGHT_BRACE) {
		a, err := p.declaration()
		if err != nil {
			return nil, err
		}
		stmt = append(stmt, a)
	}
	p.advance()
	return syntaxtree.BlockStmt{Statements: stmt}, nil
}

func (p *parser) exprStmt() (syntaxtree.Stmt, error) {
	expr, err := p.expression()
	if err != nil {
		return nil, err
	}
	if !p.check(tokenizer.SEMICOLON) {
		return nil, p.generateError("expected ;")
	}
	p.advance()
	return syntaxtree.ExpressionStmt{Expression: expr}, nil
}

func (p *parser) printStmt() (syntaxtree.Stmt, error) {
	p.advance()
	expr, err := p.expression()
	if err != nil {
		return nil, err
	}
	if !p.check(tokenizer.SEMICOLON) {
		return nil, p.generateError("expected ;")
	}
	p.advance()
	return syntaxtree.PrintStmt{Expression: expr}, nil
}

func (p *parser) expression() (syntaxtree.Expr, error) {
	return p.assignment()
}

func (p *parser) assignment() (syntaxtree.Expr, error) {
	name, err := p.bitwise()
	if err != nil {
		return nil, err
	}
	if !p.check(tokenizer.EQUAL) {
		return name, nil
	}
	if val, ok := name.(syntaxtree.LiteralExpr); !ok || val.Value.Type != tokenizer.IDENTIFIER {
		return nil, p.generateError("expected name")
	}
	op := p.advance()
	expr, err := p.expression()
	if err != nil {
		return nil, err
	}
	return syntaxtree.BinaryExpr{Left: name, Operator: op, Right: expr}, nil
}

func (p *parser) bitwise() (syntaxtree.Expr, error) {
	expr, err := p.equality()
	if err != nil {
		return nil, err
	}
	for p.checkMany([]tokenizer.TokenType{tokenizer.AMPERSAND, tokenizer.PIPE}) {
		token := p.peek()
		p.advance()
		next, err := p.equality()
		if err != nil {
			return nil, err
		}
		expr = syntaxtree.Expr(syntaxtree.BinaryExpr{Left: syntaxtree.Expr(expr), Operator: token, Right: next})
	}
	return expr, nil
}

func (p *parser) equality() (syntaxtree.Expr, error) {
	expr, err := p.comparison()
	if err != nil {
		return nil, err
	}
	for p.checkMany([]tokenizer.TokenType{tokenizer.EQUAL_EQUAL, tokenizer.EXCLAMATION_EQUAL}) {
		token := p.peek()
		p.advance()
		next, err := p.comparison()
		if err != nil {
			return nil, err
		}
		expr = syntaxtree.Expr(syntaxtree.BinaryExpr{Left: syntaxtree.Expr(expr), Operator: token, Right: next})
	}
	return expr, nil
}

func (p *parser) comparison() (syntaxtree.Expr, error) {
	expr, err := p.term()
	if err != nil {
		return nil, err
	}
	for p.checkMany([]tokenizer.TokenType{tokenizer.LESS, tokenizer.LESS_EQUAL, tokenizer.GREATER, tokenizer.GREATER_EQUAL}) {
		token := p.peek()
		p.advance()
		next, err := p.term()
		if err != nil {
			return nil, err
		}
		expr = syntaxtree.Expr(syntaxtree.BinaryExpr{Left: syntaxtree.Expr(expr), Operator: token, Right: next})
	}
	return expr, nil
}

func (p *parser) term() (syntaxtree.Expr, error) {
	expr, err := p.factor()
	if err != nil {
		return nil, err
	}
	for p.checkMany([]tokenizer.TokenType{tokenizer.PLUS, tokenizer.MINUS}) {
		token := p.peek()
		p.advance()
		next, err := p.factor()
		if err != nil {
			return nil, err
		}
		expr = syntaxtree.Expr(syntaxtree.BinaryExpr{Left: syntaxtree.Expr(expr), Operator: token, Right: next})
	}
	return expr, nil
}

func (p *parser) factor() (syntaxtree.Expr, error) {
	expr, err := p.unary()
	if err != nil {
		return nil, err
	}
	for p.checkMany([]tokenizer.TokenType{tokenizer.SLASH, tokenizer.STAR}) {
		token := p.peek()
		p.advance()
		next, err := p.unary()
		if err != nil {
			return nil, err
		}
		expr = syntaxtree.Expr(syntaxtree.BinaryExpr{Left: syntaxtree.Expr(expr), Operator: token, Right: next})
	}
	return expr, nil
}

func (p *parser) unary() (syntaxtree.Expr, error) {
	if p.checkMany([]tokenizer.TokenType{tokenizer.EXCLAMATION, tokenizer.MINUS}) {
		token := p.peek()
		p.advance()
		next, err := p.unary()
		if err != nil {
			return nil, err
		}
		return syntaxtree.Expr(syntaxtree.UnaryExpr{Operator: token, Right: next}), nil
	} else {
		return p.primary()
	}
}

func (p *parser) primary() (syntaxtree.Expr, error) {
	if p.checkMany([]tokenizer.TokenType{tokenizer.NUMBER, tokenizer.STRING, tokenizer.IDENTIFIER}) {
		return syntaxtree.Expr(syntaxtree.LiteralExpr{Value: p.advance()}), nil
	} else if p.check(tokenizer.LEFT_PAREN) {
		p.advance()
		expr, err := p.expression()
		if err != nil {
			return nil, err
		}
		p.advance()
		return syntaxtree.Expr(syntaxtree.GroupingExpr{Inside: expr}), nil
	} else {
		return nil, p.generateError("unexpected end")
	}
}

// func (p *parser) variable() (syntaxtree.Expr, error) {
// 	if p.check(tokenizer.IDENTIFIER) {
// 		return syntaxtree.Expr(syntaxtree.LiteralExpr{Value: p.advance()}), nil
// 	} else {
// 		return nil, p.generateError("expected name of variable")
// 	}
// }
