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

func TestParse_MultipleFractions(t *testing.T) {
	program := parseTokens(t, []lexer.Token{
		{Kind: "Integer", Lexeme: "1"},
		{Kind: "Slash", Lexeme: "/"},
		{Kind: "Integer", Lexeme: "2"},
		{Kind: "Integer", Lexeme: "3"},
		{Kind: "Slash", Lexeme: "/"},
		{Kind: "Integer", Lexeme: "4"},
	})

	if len(program.Statements) != 2 {
		t.Fatalf("expected 2 statements, got %d", len(program.Statements))
	}

	fraction1, ok := program.Statements[0].(ast.Fraction)
	if !ok {
		t.Fatalf("expected ast.Fraction, got %T", program.Statements[0])
	}
	if fraction1.Numerator.String() != "1" {
		t.Errorf("expected numerator 1, got %s", fraction1.Numerator)
	}
	if fraction1.Denominator.String() != "2" {
		t.Errorf("expected denominator 2, got %s", fraction1.Denominator)
	}

	fraction2, ok := program.Statements[1].(ast.Fraction)
	if !ok {
		t.Fatalf("expected ast.Fraction, got %T", program.Statements[1])
	}
	if fraction2.Numerator.String() != "3" {
		t.Errorf("expected numerator 3, got %s", fraction2.Numerator)
	}
	if fraction2.Denominator.String() != "4" {
		t.Errorf("expected denominator 4, got %s", fraction2.Denominator)
	}
}

func TestParse_InvalidTokens(t *testing.T) {
	expectParseError(t, []lexer.Token{
		{Kind: "Slash", Lexeme: "/"},
	})
	expectParseError(t, []lexer.Token{
		{Kind: "Integer", Lexeme: "1"},
	})
	expectParseError(t, []lexer.Token{
		{Kind: "Integer", Lexeme: "1"},
		{Kind: "Slash", Lexeme: "/"},
	})
	expectParseError(t, []lexer.Token{
		{Kind: "Slash", Lexeme: "/"},
		{Kind: "Integer", Lexeme: "2"},
	})
	expectParseError(t, []lexer.Token{
		{Kind: "Integer", Lexeme: "1"},
		{Kind: "Integer", Lexeme: "2"},
	})
	expectParseError(t, []lexer.Token{
		{Kind: "Slash", Lexeme: "/"},
		{Kind: "Slash", Lexeme: "/"},
	})
}
