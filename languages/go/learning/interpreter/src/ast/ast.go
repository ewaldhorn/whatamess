package ast

import "interpreter/src/token"

// Abstract Syntax Trees (AST) are hierarchical tree representations of the
// syntactic structure of source code. Each node in the tree denotes a construct
// occurring in the code. ASTs are commonly used in compilers and interpreters
// as an intermediate step between parsing source code and generating executable
// code or evaluating expressions. Unlike parse trees that represent exact
// grammar syntactic details, ASTs abstract away syntax details (like
// parentheses and delimiters) to focus on the logical structure and meaning
// of the code.

// ----------------------------------------------------------------------------
type Node interface {
	TokenLiteral() string
}

// ----------------------------------------------------------------------------
type Statement interface {
	Node
	statementNode()
}

// ----------------------------------------------------------------------------
type Expression interface {
	Node
	expressionNode()
}

// ----------------------------------------------------------------------------
// Program Node representation of a Parser's program structure that
// contains a sequence of statements.
type Program struct {
	Statements []Statement
}

func (p *Program) TokenLiteral() string {
	if len(p.Statements) > 0 {
		return p.Statements[0].TokenLiteral()
	} else {
		return ""
	}
}

// ----------------------------------------------------------------------------
type LetStatement struct {
	Token token.Token
	Name  *Identifier
	Value Expression
}

func (ls *LetStatement) statementNode()       {}
func (ls *LetStatement) TokenLiteral() string { return ls.Token.Literal }

// ----------------------------------------------------------------------------
type Identifier struct {
	Token token.Token
	Value string
}

func (i *Identifier) expressionNode()      {}
func (i *Identifier) TokenLiteral() string { return i.Token.Literal }
