package parser

import (
	"interpreter/src/ast"
	"interpreter/src/lexer"
	"interpreter/src/token"
)

// ----------------------------------------------------------------------------
type Parser struct {
	lexer     *lexer.Lexer
	curToken  token.Token
	peekToken token.Token
}

// ----------------------------------------------------------------------------
func (p *Parser) nextToken() {
	p.curToken = p.peekToken
	p.peekToken = p.lexer.NextToken()
}

// ----------------------------------------------------------------------------
func (p *Parser) ParseProgram() *ast.Program {
	return nil
}

// ----------------------------------------------------------------------------
func New(l *lexer.Lexer) *Parser {
	p := &Parser{lexer: l}

	// read two tokens, so the current and peek tokens are set up
	p.nextToken()
	p.nextToken()

	return p
}
