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
	ASTERISK = "*"
	BANG     = "!"
	SLASH    = "/"
	MINUS    = "-"
	PLUS     = "+"

	COMMA     = ","
	SEMICOLON = ";"

	LT = "<"
	GT = ">"

	OPAREN = "("
	CPAREN = ")"
	OBRACE = "{"
	CBRACE = "}"

	FUNCTION = "FUNCTION"
	LET      = "LET"
	TRUE     = "TRUE"
	FALSE    = "FALSE"
	IF       = "IF"
	ELSE     = "ELSE"
	RETURN   = "RETURN"
)

// ----------------------------------------------------------------------------
var keywords = map[string]TokenType{
	"else":   ELSE,
	"false":  FALSE,
	"func":   FUNCTION,
	"if":     IF,
	"let":    LET,
	"return": RETURN,
	"true":   TRUE,
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
