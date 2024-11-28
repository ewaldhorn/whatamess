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

	ASSIGN   = "="
	DIVIDE   = "/"
	MINUS    = "-"
	MULTIPLY = "*"
	PLUS     = "+"

	COMMA     = ","
	SEMICOLON = ";"

	LPAREN = "("
	RPAREN = ")"
	LBRACE = "{"
	RBRACE = "}"

	FUNCTION = "FUNCTION"
	LET      = "LET"
)

// ----------------------------------------------------------------------------
var keywords = map[string]TokenType{
	"func": FUNCTION,
	"let":  LET,
}

// ----------------------------------------------------------------------------
// LookupIdentifier returns the TokenType corresponding to the given identifier.
// If the identifier matches a language keyword, it returns the keyword's TokenType.
// Otherwise, it returns IDENT for regular identifiers.
func LookupIdentifier(identifier string) TokenType {
	if tok, ok := keywords[identifier]; ok {
		return tok
	}
	return IDENT
}
