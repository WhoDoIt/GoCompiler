package parser

import (
	"errors"

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

func (p *parser) check(tokenType []tokenizer.TokenType) bool {
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

func (p *parser) synchronize() {
	p.advance()

	for !p.isAtEnd() {
		if p.previus().Type == tokenizer.SEMICOLON {
			return
		}
		switch p.peek().Type {
		case tokenizer.FN, tokenizer.VAR, tokenizer.IF, tokenizer.ELSE, tokenizer.RETURN, tokenizer.FOR, tokenizer.STRUCT:
			return
		}
		p.advance()
	}
}

func Parse(tokens []tokenizer.Token) (syntaxtree.Expr, error) {
	parser := parser{tokens: tokens}
	return parser.expression()
}

func (p *parser) expression() (syntaxtree.Expr, error) {
	return p.bitwise()
}

func (p *parser) bitwise() (syntaxtree.Expr, error) {
	expr, err := p.equality()
	if err != nil {
		return nil, err
	}
	for p.check([]tokenizer.TokenType{tokenizer.AMPERSAND, tokenizer.PIPE}) {
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
	for p.check([]tokenizer.TokenType{tokenizer.EQUAL_EQUAL, tokenizer.EXCLAMATION_EQUAL}) {
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
	for p.check([]tokenizer.TokenType{tokenizer.LESS, tokenizer.LESS_EQUAL, tokenizer.GREATER, tokenizer.GREATER_EQUAL}) {
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
	for p.check([]tokenizer.TokenType{tokenizer.PLUS, tokenizer.MINUS}) {
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
	for p.check([]tokenizer.TokenType{tokenizer.SLASH, tokenizer.STAR}) {
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
	if p.check([]tokenizer.TokenType{tokenizer.EXCLAMATION, tokenizer.MINUS}) {
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
	if p.check([]tokenizer.TokenType{tokenizer.NUMBER, tokenizer.STRING}) {
		return syntaxtree.Expr(syntaxtree.Literal{Value: p.advance()}), nil
	} else if p.check([]tokenizer.TokenType{tokenizer.LEFT_PAREN}) {
		p.advance()
		expr, err := p.expression()
		if err != nil {
			return nil, err
		}
		p.advance()
		return syntaxtree.Expr(syntaxtree.GroupingExpr{Inside: expr}), nil
	} else {
		return nil, errors.New("unexpected end")
	}
}
