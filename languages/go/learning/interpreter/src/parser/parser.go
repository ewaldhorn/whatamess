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
	program := &ast.Program{}
	program.Statements = []ast.Statement{}

	for !p.curTokenIs(token.EOF) {
		stmt := p.parseStatement()
		if stmt != nil {
			program.Statements = append(program.Statements, stmt)
		}
		p.nextToken()
	}

	return program
}

// ----------------------------------------------------------------------------
func (p *Parser) parseStatement() ast.Statement {
	switch p.curToken.Type {
	case token.LET:
		return p.parseLetStatement()
	default:
		return nil
	}
}

// ----------------------------------------------------------------------------
func (p *Parser) parseLetStatement() *ast.LetStatement {
	statement := &ast.LetStatement{Token: p.curToken}

	if !p.expectPeek(token.IDENT) {
		return nil
	}

	statement.Name = &ast.Identifier{Token: p.curToken, Value: p.curToken.Literal}

	if !p.expectPeek(token.ASSIGN) {
		return nil
	}

	// TODO: skipping expressions for now until ; found
	for !p.curTokenIs(token.SEMICOLON) {
		p.nextToken()
	}

	return statement
}

// ----------------------------------------------------------------------------
func (p *Parser) curTokenIs(t token.TokenType) bool {
	return p.curToken.Type == t
}

// ----------------------------------------------------------------------------
func (p *Parser) peekTokenIs(t token.TokenType) bool {
	return p.peekToken.Type == t
}

// ----------------------------------------------------------------------------
func (p *Parser) expectPeek(t token.TokenType) bool {
	if p.peekTokenIs(t) {
		p.nextToken()
		return true
	} else {
		return false
	}
}

// ----------------------------------------------------------------------------
func New(l *lexer.Lexer) *Parser {
	p := &Parser{lexer: l}

	// read two tokens, so the current and peek tokens are set up
	p.nextToken()
	p.nextToken()

	return p
}
