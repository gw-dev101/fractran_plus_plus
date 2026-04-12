package parser

import (
	"testing"

	"github.com/gw-dev101/fractran_plus_plus/internal/ast"
	"github.com/gw-dev101/fractran_plus_plus/internal/lexer"
)

// helper to reduce boilerplate
func parseTokens(t *testing.T, tokens []lexer.Token) *ast.Program {
	t.Helper()

	p := New()
	program, err := p.Parse(tokens)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	return program
}

func expectParseError(t *testing.T, tokens []lexer.Token) {
	t.Helper()

	p := New()
	_, err := p.Parse(tokens)
	if err == nil {
		t.Fatal("expected error, got nil")
	}
	if _, ok := err.(*ParseError); !ok {
		t.Fatalf("expected ParseError, got %T", err)
	}
}
func TestParse_SingleFraction(t *testing.T) {
	program := parseTokens(t, []lexer.Token{
		{Kind: "Integer", Lexeme: "1"},
		{Kind: "Slash", Lexeme: "/"},
		{Kind: "Integer", Lexeme: "2"},
	})

	if len(program.Statements) != 1 {
		t.Fatalf("expected 1 statement, got %d", len(program.Statements))
	}

	fraction, ok := program.Statements[0].(ast.Fraction)
	if !ok {
		t.Fatalf("expected ast.Fraction, got %T", program.Statements[0])
	}

	if fraction.Numerator.String() != "1" {
		t.Errorf("expected numerator 1, got %s", fraction.Numerator)
	}
	if fraction.Denominator.String() != "2" {
		t.Errorf("expected denominator 2, got %s", fraction.Denominator)
	}
}
