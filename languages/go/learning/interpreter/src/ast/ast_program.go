package ast

// ----------------------------------------------------------------------------
// Program Node representation of a Parser's program structure that
// contains a sequence of statements.
type Program struct {
	Statements []Statement
}

// ----------------------------------------------------------------------------
func (p *Program) TokenLiteral() string {
	if len(p.Statements) > 0 {
		return p.Statements[0].TokenLiteral()
	} else {
		return ""
	}
}
