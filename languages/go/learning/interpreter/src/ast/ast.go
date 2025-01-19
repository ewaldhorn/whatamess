package ast

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
