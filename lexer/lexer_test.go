package lexer

import (
	"testing"

	"github.com/MangoSteen0903/go-interpreter/token"
)

type testLexerConfig struct {
	expectedType    token.TokenType
	expectedLiteral string
}

func TestNextToken(t *testing.T) {
	input := `=+(){},;`

	tests := []testLexerConfig{
		{token.ASSIGN, "="},
		{token.PLUS, "+"},
		{token.LAPREN, "("},
		{token.RAPREN, ")"},
		{token.LBRACE, "{"},
		{token.RBRACE, "}"},
		{token.COMMA, ","},
		{token.SEMICOLON, ";"},
		{token.EOF, ""},
	}

	l := New(input)

	for i, tt := range tests {
		tok := l.NextToken()
		if tok.Type != tt.expectedType {
			t.Fatalf("tests[%d] - tokentype wrong, expected = %q, got = %q", i, tt.expectedType, tok.Type)
		}

		if tok.Literal != tt.expectedLiteral {
			t.Fatalf("tests[%d] - tokentype wrong, expected = %q, got = %q", i, tt.expectedLiteral, tok.Literal)
		}
	}
}
