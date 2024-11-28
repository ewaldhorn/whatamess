package lexer

import (
	"interpreter/src/token"
	"testing"
)

// ----------------------------------------------------------------------------
// Test helpers for working with the lexer.
// Tests exercise the behavior of tokens and token operations.
func Test_SmokeTest_NextToken(t *testing.T) {
	input := "=+-/*(){},;$!!!<===+-!="

	tests := []struct {
		expectedType    token.TokenType
		expectedLiteral string
	}{
		{token.ASSIGN, "="},
		{token.PLUS, "+"},
		{token.MINUS, "-"},
		{token.SLASH, "/"},
		{token.ASTERISK, "*"},
		{token.OPAREN, "("},
		{token.CPAREN, ")"},
		{token.OBRACE, "{"},
		{token.CBRACE, "}"},
		{token.COMMA, ","},
		{token.SEMICOLON, ";"},
		{token.ILLEGAL, "$"}, // for now, $ is not supported
		{token.BANG, "!"},
		{token.BANG, "!"},
		{token.BANG, "!"},
		{token.LTE, "<="},
		{token.EQUALS, "=="},
		{token.PLUS, "+"},
		{token.MINUS, "-"},
		{token.NOTEQUALS, "!="},
		{token.EOF, ""},
	}

	l := New(input)

	for i, tt := range tests {
		tok := l.NextToken()

		if tok.Type != tt.expectedType {
			t.Fatalf("tests[%d] - tokenType wrong. expected=%q got=%q", i, tt.expectedType, tok.Type)
		}

		if tok.Literal != tt.expectedLiteral {
			t.Fatalf("tests[%d] - literal wrong. expected='%q', got='%q'", i, tt.expectedLiteral, tok.Literal)
		}
	}
}

// ----------------------------------------------------------------------------
func Test_SyntaxTest_NextToken(t *testing.T) {
	input := `let five = 5;
	let ten = 10;
	let add = func(x, y) {
		x + y;
	};

	let result=add(five, ten);
	!-/*5;
	5 < 10 > 5;

	if (5 < 10) {
		return true;
	} else {
		return false;
	}

	10 == 10;
	10 != 9;

	5 >= 6;
	6 <= 5;
	`

	tests := []struct {
		expectedType    token.TokenType
		expectedLiteral string
	}{
		{token.LET, "let"},
		{token.IDENT, "five"},
		{token.ASSIGN, "="},
		{token.INT, "5"},
		{token.SEMICOLON, ";"},
		{token.LET, "let"},
		{token.IDENT, "ten"},
		{token.ASSIGN, "="},
		{token.INT, "10"},
		{token.SEMICOLON, ";"},
		{token.LET, "let"},
		{token.IDENT, "add"},
		{token.ASSIGN, "="},
		{token.FUNCTION, "func"},
		{token.OPAREN, "("},
		{token.IDENT, "x"},
		{token.COMMA, ","},
		{token.IDENT, "y"},
		{token.CPAREN, ")"},
		{token.OBRACE, "{"},
		{token.IDENT, "x"},
		{token.PLUS, "+"},
		{token.IDENT, "y"},
		{token.SEMICOLON, ";"},
		{token.CBRACE, "}"},
		{token.SEMICOLON, ";"},
		{token.LET, "let"},
		{token.IDENT, "result"},
		{token.ASSIGN, "="},
		{token.IDENT, "add"},
		{token.OPAREN, "("},
		{token.IDENT, "five"},
		{token.COMMA, ","},
		{token.IDENT, "ten"},
		{token.CPAREN, ")"},
		{token.SEMICOLON, ";"},
		{token.BANG, "!"},
		{token.MINUS, "-"},
		{token.SLASH, "/"},
		{token.ASTERISK, "*"},
		{token.INT, "5"},
		{token.SEMICOLON, ";"},
		{token.INT, "5"},
		{token.LT, "<"},
		{token.INT, "10"},
		{token.GT, ">"},
		{token.INT, "5"},
		{token.SEMICOLON, ";"},
		{token.IF, "if"},
		{token.OPAREN, "("},
		{token.INT, "5"},
		{token.LT, "<"},
		{token.INT, "10"},
		{token.CPAREN, ")"},
		{token.OBRACE, "{"},
		{token.RETURN, "return"},
		{token.TRUE, "true"},
		{token.SEMICOLON, ";"},
		{token.CBRACE, "}"},
		{token.ELSE, "else"},
		{token.OBRACE, "{"},
		{token.RETURN, "return"},
		{token.FALSE, "false"},
		{token.SEMICOLON, ";"},
		{token.CBRACE, "}"},
		{token.INT, "10"},
		{token.EQUALS, "=="},
		{token.INT, "10"},
		{token.SEMICOLON, ";"},
		{token.INT, "10"},
		{token.NOTEQUALS, "!="},
		{token.INT, "9"},
		{token.SEMICOLON, ";"},
		{token.INT, "5"},
		{token.GTE, ">="},
		{token.INT, "6"},
		{token.SEMICOLON, ";"},
		{token.INT, "6"},
		{token.LTE, "<="},
		{token.INT, "5"},
		{token.SEMICOLON, ";"},
		{token.EOF, ""},
	}

	l := New(input)

	for i, tt := range tests {
		tok := l.NextToken()

		if tok.Type != tt.expectedType {
			t.Fatalf("syntax tests[%d] - tokenType wrong. expected=%q got=%q", i, tt.expectedType, tok.Type)
		}

		if tok.Literal != tt.expectedLiteral {
			t.Fatalf("syntax tests[%d] - literal wrong. expected='%q', got='%q'", i, tt.expectedLiteral, tok.Literal)
		}
	}
}
