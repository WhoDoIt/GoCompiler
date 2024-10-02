package tokenizer

import (
	"errors"
)

type TokenType int

const (
	UNIDENTIFIED TokenType = iota

	// SINGLE CHARACTER
	LEFT_BRACE
	RIGHT_BRACE
	LEFT_PAREN
	RIGHT_PAREN
	PIPE
	AMPERSAND
	PLUS
	MINUS
	SLASH
	STAR
	SEMICOLON
	COMMA
	DOT
	EXCLAMATION
	EQUAL
	LESS
	GREATER

	// 2 CHARACTERS
	EXCLAMATION_EQUAL
	EQUAL_EQUAL
	LESS_EQUAL
	GREATER_EQUAL

	// MULTIPLE CHARACTERS
	IDENTIFIER
	STRING
	NUMBER

	// KEYWORDS
	AND
	OR
	IF
	ELSE
	WHILE
	FN
	FOR
	RETURN
	VAR
	STRUCT
	PRINT
	TRUE
	FALSE

	EOF
)

type Token struct {
	Type    TokenType
	Content string
	Len     int
	Line    int
}

type tokenizer struct {
	data    []byte
	start   int
	current int
	line    int
}

func (t *tokenizer) TakeToken() (Token, error) {
	if t.IsAtEnd() {
		return Token{Type: EOF}, nil
	}
	t.start = t.current
	var c = t.Advance()
	switch c {
	case '{':
		return Token{LEFT_BRACE, "{", 1, t.line}, nil
	case '}':
		return Token{RIGHT_BRACE, "}", 1, t.line}, nil
	case '(':
		return Token{LEFT_PAREN, "(", 1, t.line}, nil
	case ')':
		return Token{RIGHT_PAREN, ")", 1, t.line}, nil
	case '+':
		return Token{PLUS, "+", 1, t.line}, nil
	case '-':
		return Token{MINUS, "-", 1, t.line}, nil
	case '|':
		return Token{PIPE, "|", 1, t.line}, nil
	case '&':
		return Token{AMPERSAND, "&", 1, t.line}, nil
	case '/':
		return Token{SLASH, "/", 1, t.line}, nil
	case '*':
		return Token{STAR, "*", 1, t.line}, nil
	case ';':
		return Token{SEMICOLON, ";", 1, t.line}, nil
	case '.':
		return Token{DOT, ".", 1, t.line}, nil
	case ',':
		return Token{COMMA, ",", 1, t.line}, nil
	case '!':
		if t.Match('=') {
			t.Advance()
			return Token{EXCLAMATION_EQUAL, "!=", 2, t.line}, nil
		} else {
			return Token{EXCLAMATION, "!", 1, t.line}, nil
		}
	case '=':
		if t.Match('=') {
			t.Advance()
			return Token{EQUAL_EQUAL, "==", 2, t.line}, nil
		} else {
			return Token{EQUAL, "=", 1, t.line}, nil
		}
	case '<':
		if t.Match('=') {
			t.Advance()
			return Token{LESS_EQUAL, "<=", 2, t.line}, nil
		} else {
			return Token{LESS, "<", 1, t.line}, nil
		}
	case '>':
		if t.Match('=') {
			t.Advance()
			return Token{GREATER_EQUAL, "!=", 2, t.line}, nil
		} else {
			return Token{GREATER, "!", 1, t.line}, nil
		}
	case ' ', '\t', '\r':
		return t.TakeToken()
	case '\n':
		t.line += 1
		return t.TakeToken()
	case '"':
		for t.Peak() != '"' && t.Peak() != 0 {
			t.Advance()
		}
		if t.Peak() == 0 {
			return Token{Type: UNIDENTIFIED}, errors.New("unclosed string")
		}
		t.Advance()
		return Token{STRING, string(t.data[t.start+1 : t.current-1]), t.current - t.start - 2, t.line}, nil
	}

	switch {
	case t.IsDigit(c):
		for t.IsDigit(t.Peak()) {
			t.Advance()
		}
		if t.Peak() == '.' {
			t.Advance()
			if !t.IsDigit(t.Peak()) {
				return Token{Type: UNIDENTIFIED}, errors.New("unclosed float number")
			}
			for t.IsDigit(t.Peak()) {
				t.Advance()
			}
		}
		return Token{NUMBER, string(t.data[t.start:t.current]), t.current - t.start, t.line}, nil

	}

	var keywords map[string]TokenType = make(map[string]TokenType)
	keywords["and"] = AND
	keywords["or"] = OR
	keywords["if"] = IF
	keywords["else"] = ELSE
	keywords["while"] = WHILE
	keywords["fn"] = FN
	keywords["for"] = FOR
	keywords["return"] = RETURN
	keywords["var"] = VAR
	keywords["struct"] = STRUCT
	keywords["print"] = PRINT
	keywords["true"] = TRUE
	keywords["false"] = FALSE

	for t.IsDigit(t.Peak()) || t.IsGoodChar(t.Peak()) {
		t.Advance()
	}

	var word = string(t.data[t.start:t.current])

	if val, ok := keywords[word]; ok {
		return Token{val, word, len(word), t.line}, nil
	} else {
		return Token{IDENTIFIER, word, len(word), t.line}, nil
	}
}

func (t *tokenizer) IsDigit(char byte) bool {
	return char >= '0' && char <= '9'
}

func (t *tokenizer) IsGoodChar(char byte) bool {
	return char >= 'a' && char <= 'z' || char >= 'A' && char <= 'Z' || char == '_'
}

func (t *tokenizer) IsAtEnd() bool {
	return t.current >= len(t.data)
}

func (t *tokenizer) Match(char byte) bool {
	if t.IsAtEnd() {
		return false
	}
	return t.data[t.current] == char
}

func (t *tokenizer) Advance() byte {
	if t.IsAtEnd() {
		return 0
	}
	var res = t.data[t.current]
	t.current += 1
	return res
}

func (t *tokenizer) Peak() byte {
	if t.IsAtEnd() {
		return 0
	}
	var res = t.data[t.current]
	return res
}

func (t *tokenizer) PeakNext(off int) byte {
	if t.current+off >= len(t.data) {
		return 0
	}
	var res = t.data[t.current+off]
	return res
}

func Tokenize(data []byte) ([]Token, error) {
	var tokens []Token
	var tk = tokenizer{data: data, line: 1}
	for !tk.IsAtEnd() {
		token, err := tk.TakeToken()
		if err != nil {
			return nil, err
		}
		tokens = append(tokens, token)
	}
	if tokens[len(tokens)-1].Type != EOF {
		tokens = append(tokens, Token{EOF, "EOF", 3, tk.line})
	}
	return tokens, nil
}
