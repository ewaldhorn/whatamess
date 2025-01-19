package lexer

import (
	"interpreter/src/token"
)

// ----------------------------------------------------------------------------
type Lexer struct {
	input        string
	position     int  // current pos in input, points to current char
	readPosition int  // current reading position in input, after current char
	ch           byte // current char being inspected
}

// ----------------------------------------------------------------------------
// New creates a new lexer that will process the input string.
// Returns a pointer to the newly created lexer
func New(input string) *Lexer {
	newLexer := &Lexer{input: input}
	newLexer.readChar()

	return newLexer
}

// ----------------------------------------------------------------------------
// readChar reads the next character in the input and advances l.position and
// l.readPosition. When it reaches the end of input, it sets l.ch to 0 (NUL)
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
// peekChar reads the next character in the input but does not increment the
// position. When it reaches the end of input, it returns 0 (NUL)
func (l *Lexer) peekChar() byte {
	if l.readPosition >= len(l.input) {
		return 0
	} else {
		return l.input[l.readPosition]
	}
}

// ----------------------------------------------------------------------------
// NextToken processes and returns the next token in the input.
// The token.Token struct contains both the type of token and the
// literal value associated with it
func (l *Lexer) NextToken() token.Token {
	var tok token.Token

	l.skipWhitespace()

	switch l.ch {
	case '=':
		if l.peekChar() == '=' { // TODO: extract common behaviour out
			ch := l.ch
			l.readChar()
			tok = token.Token{Type: token.EQUALS, Literal: string(ch) + string(l.ch)}
		} else {
			tok = newToken(token.ASSIGN, l.ch)
		}
	case ';':
		tok = newToken(token.SEMICOLON, l.ch)
	case '(':
		tok = newToken(token.OPAREN, l.ch)
	case ')':
		tok = newToken(token.CPAREN, l.ch)
	case '{':
		tok = newToken(token.OBRACE, l.ch)
	case '}':
		tok = newToken(token.CBRACE, l.ch)
	case ',':
		tok = newToken(token.COMMA, l.ch)
	case '+':
		tok = newToken(token.PLUS, l.ch)
	case '-':
		tok = newToken(token.MINUS, l.ch)
	case '*':
		tok = newToken(token.ASTERISK, l.ch)
	case '/':
		tok = newToken(token.SLASH, l.ch)
	case '!':
		if l.peekChar() == '=' {
			ch := l.ch
			l.readChar()
			tok = token.Token{Type: token.NOTEQUALS, Literal: string(ch) + string(l.ch)}
		} else {
			tok = newToken(token.BANG, l.ch)
		}
	case '<':
		if l.peekChar() == '=' {
			ch := l.ch
			l.readChar()
			tok = token.Token{Type: token.LTE, Literal: string(ch) + string(l.ch)}
		} else {
			tok = newToken(token.LT, l.ch)
		}
	case '>':
		if l.peekChar() == '=' {
			ch := l.ch
			l.readChar()
			tok = token.Token{Type: token.GTE, Literal: string(ch) + string(l.ch)}
		} else {
			tok = newToken(token.GT, l.ch)
		}
	case 0:
		tok.Literal = ""
		tok.Type = token.EOF
	default:
		if isLetter(l.ch) {
			tok.Literal = l.readIdentifier()
			tok.Type = token.LookupIdentifier(tok.Literal)
			return tok
		} else if isDigit(l.ch) {
			tok.Type = token.INT
			tok.Literal = l.readNumber()
			return tok
		} else {
			tok = newToken(token.ILLEGAL, l.ch)
		}
	}

	l.readChar()
	return tok
}

// ----------------------------------------------------------------------------
func (l *Lexer) readIdentifier() string {
	position := l.position
	for isLetter(l.ch) {
		l.readChar()
	}
	return l.input[position:l.position]
}

// ----------------------------------------------------------------------------
// TODO: This only reads integers, no floats
func (l *Lexer) readNumber() string {
	position := l.position
	for isDigit(l.ch) {
		l.readChar()
	}
	return l.input[position:l.position]
}

// ----------------------------------------------------------------------------
// skipWhitespace skips over any whitespace characters (space, tab, newline,
// carriage return) in the input by repeatedly calling readChar() until a
// non-whitespace character is found
func (l *Lexer) skipWhitespace() {
	for l.ch == ' ' || l.ch == '\t' || l.ch == '\n' || l.ch == '\r' {
		l.readChar()
	}
}

// ----------------------------------------------------------------------------
// isLetter checks that a characters is an acceptable letter/character for
// use in an identifier
func isLetter(ch byte) bool {
	return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z' || ch == '_'
}

// ----------------------------------------------------------------------------
// isDigit will tell us if it's a digit from 0..9
func isDigit(ch byte) bool {
	return '0' <= ch && ch <= '9'
}

// ----------------------------------------------------------------------------
func newToken(tokenType token.TokenType, ch byte) token.Token {
	return token.Token{Type: tokenType, Literal: string(ch)}
}
