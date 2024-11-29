package parser

import (
	"interpreter/src/ast"
	"interpreter/src/lexer"
	"testing"
)

// ----------------------------------------------------------------------------
func Test_LetStatements(t *testing.T) {
	input := `
	let x = 5;
	let y = 10;
	let z = 2;
	`
	parser := New(lexer.New(input))

	program := parser.ParseProgram()
	if program == nil {
		t.Fatalf("ParseProgram returned nil")
	}

	if len(program.Statements) != 3 {
		t.Fatalf("program.Statements should have 3 entries, it has %d", len(program.Statements))
	}

	tests := []struct {
		expectedIdentifier string
	}{
		{"x"}, {"y"}, {"z"},
	}

	for i, test := range tests {
		statement := program.Statements[i]
		if !testLetStatement(t, statement, test.expectedIdentifier) {
			return
		}
	}
}

// ----------------------------------------------------------------------------
func testLetStatement(t *testing.T, statement ast.Statement, name string) bool {
	if statement.TokenLiteral() != "let" {
		t.Errorf("TokenLiteral was %q, not 'let'", statement.TokenLiteral())
		return false
	}

	letStatement, ok := statement.(*ast.LetStatement)
	if !ok {
		t.Errorf("statement is not *ast.LetStatement, it is %T", statement)
		return false
	}

	if letStatement.Name.Value != name {
		t.Errorf("letStatement.Name is not '%s', it is '%s'", name, letStatement.Name.Value)
		return false
	}

	if letStatement.Name.TokenLiteral() != name {
		t.Errorf("letStatement.Name.TokenLiteral is not '%s', it is '%s'", name, letStatement.Name.TokenLiteral())
		return false
	}

	return true
}
