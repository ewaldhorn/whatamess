package parser

import (
	"interpreter/src/ast"
	"interpreter/src/lexer"
	"interpreter/src/token"
)

// ----------------------------------------------------------------------------
type Parser struct {
	l *lexer.Lexer

	curToken  token.Token
	peekToken token.Token
}

// ----------------------------------------------------------------------------
func (p *Parser) nextToken() {
	p.curToken = p.peekToken
	p.peekToken = p.l.NextToken()
}

// ----------------------------------------------------------------------------
func (p *Parser) ParseProgram() *ast.Program {
	return nil
}

// ----------------------------------------------------------------------------
func New(l *lexer.Lexer) *Parser {
	p := &Parser{l: l}

	// read two tokens, so cur and peek tokens are set up
	p.nextToken()
	p.nextToken()

	return p
}
