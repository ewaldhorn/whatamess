package lexer

import "interpreter/src/token"

// ----------------------------------------------------------------------------
type Lexer struct {
	input        string
	position     int  // current pos in input, points to current char
	readPosition int  // current reading position in input, after current char
	ch           byte // current char being inspected
}

// ----------------------------------------------------------------------------
// New creates a new lexer that will process the input string.
// Returns a pointer to the newly created lexer.
func New(input string) *Lexer {
	newLexer := &Lexer{input: input}
	newLexer.readChar()

	return newLexer
}

// ----------------------------------------------------------------------------
// readChar reads the next character in the input and advances l.position and
// l.readPosition. When it reaches the end of input, it sets l.ch to 0 (NUL).
func (l *Lexer) readChar() {
	if l.readPosition >= len(l.input) {
		l.ch = 0
	} else {
		l.ch = l.input[l.readPosition]
	}
	l.position = l.readPosition
	l.readPosition += 1
}

// ----------------------------------------------------------------------------
// NextToken processes and returns the next token in the input.
// The token.Token struct contains both the type of token and the
// literal value associated with it.
func (l *Lexer) NextToken() token.Token {
	var tok token.Token

	switch l.ch {
	case '=':
		tok = newToken(token.ASSIGN, l.ch)
	case ';':
		tok = newToken(token.SEMICOLON, l.ch)
	case '(':
		tok = newToken(token.LPAREN, l.ch)
	case ')':
		tok = newToken(token.RPAREN, l.ch)
	case '{':
		tok = newToken(token.LBRACE, l.ch)
	case '}':
		tok = newToken(token.RBRACE, l.ch)
	case ',':
		tok = newToken(token.COMMA, l.ch)
	case '+':
		tok = newToken(token.PLUS, l.ch)
	case 0:
		tok.Literal = ""
		tok.Type = token.EOF
	default:
		tok = newToken(token.ILLEGAL, l.ch)
	}

	l.readChar()
	return tok
}

// ----------------------------------------------------------------------------
func newToken(tokenType token.TokenType, ch byte) token.Token {
	return token.Token{Type: tokenType, Literal: string(ch)}
}
