package token

// ----------------------------------------------------------------------------
type TokenType string

// ----------------------------------------------------------------------------
type Token struct {
	Type    TokenType
	Literal string
}

// ----------------------------------------------------------------------------
// Possible token types for the language
const (
	ILLEGAL = "ILLEGAL" // unknown token
	EOF     = "EOF"     // end of file

	IDENT = "IDENT" // identifiers
	INT   = "INT"   // integers

	ASSIGN = "="
	PLUS   = "+"

	COMMA     = ","
	SEMICOLON = ";"

	LPAREN = "("
	RPAREN = ")"
	LBRACE = "{"
	RBRACE = "}"

	FUNCTION = "FUNCTION"
	LET      = "LET"
)
