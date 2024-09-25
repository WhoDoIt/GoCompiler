package parser

import (
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

func Parse(tokens []tokenizer.Token) syntaxtree.Expr {
	parser := parser{tokens: tokens}
	return parser.expression()
}

func (p *parser) expression() syntaxtree.Expr {
	return p.equality()
}

func (p *parser) equality() syntaxtree.Expr {
	expr := p.comparison()
	for p.check([]tokenizer.TokenType{tokenizer.EQUAL_EQUAL, tokenizer.EXCLAMATION_EQUAL}) {
		token := p.peek()
		p.advance()
		expr = syntaxtree.Expr(syntaxtree.BinaryExpr{Left: syntaxtree.Expr(expr), Operator: token, Right: p.comparison()})
	}
	return expr
}

func (p *parser) comparison() syntaxtree.Expr {
	expr := p.term()
	for p.check([]tokenizer.TokenType{tokenizer.LESS, tokenizer.LESS_EQUAL, tokenizer.GREATER, tokenizer.GREATER_EQUAL}) {
		token := p.peek()
		p.advance()
		expr = syntaxtree.Expr(syntaxtree.BinaryExpr{Left: syntaxtree.Expr(expr), Operator: token, Right: p.term()})
	}
	return expr
}

func (p *parser) term() syntaxtree.Expr {
	expr := p.factor()
	for p.check([]tokenizer.TokenType{tokenizer.PLUS, tokenizer.MINUS}) {
		token := p.peek()
		p.advance()
		expr = syntaxtree.Expr(syntaxtree.BinaryExpr{Left: syntaxtree.Expr(expr), Operator: token, Right: p.factor()})
	}
	return expr
}

func (p *parser) factor() syntaxtree.Expr {
	expr := p.unary()
	for p.check([]tokenizer.TokenType{tokenizer.SLASH, tokenizer.STAR}) {
		token := p.peek()
		p.advance()
		expr = syntaxtree.Expr(syntaxtree.BinaryExpr{Left: syntaxtree.Expr(expr), Operator: token, Right: p.unary()})
	}
	return expr
}

func (p *parser) unary() syntaxtree.Expr {
	if p.check([]tokenizer.TokenType{tokenizer.EXCLAMATION, tokenizer.MINUS}) {
		token := p.peek()
		p.advance()
		return syntaxtree.Expr(syntaxtree.UnaryExpr{Operator: token, Right: p.unary()})
	} else {
		return p.primary()
	}
}

func (p *parser) primary() syntaxtree.Expr {
	if p.check([]tokenizer.TokenType{tokenizer.NUMBER, tokenizer.STRING}) {
		return syntaxtree.Expr(syntaxtree.Literal{Value: p.advance()})
	} else if p.check([]tokenizer.TokenType{tokenizer.LEFT_PAREN}) {
		p.advance()
		expr := p.expression()
		p.advance()
		return syntaxtree.Expr(syntaxtree.GroupingExpr{Inside: expr})
	} else {
		return syntaxtree.Literal{Value: tokenizer.Token{Type: tokenizer.EOF, Content: "BRUH", Len: 4}}
	}
}
