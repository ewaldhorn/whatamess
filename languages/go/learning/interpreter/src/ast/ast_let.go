package ast

import "interpreter/src/token"

// ----------------------------------------------------------------------------
type LetStatement struct {
	Token token.Token
	Name  *Identifier
	Value Expression
}

// ----------------------------------------------------------------------------
func (ls *LetStatement) statementNode() {}

// ----------------------------------------------------------------------------
func (ls *LetStatement) TokenLiteral() string { return ls.Token.Literal }
