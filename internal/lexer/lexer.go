package lexer

import (
	"fmt"
)

// ErrNotImplemented is returned by the stub lexer until tokenization exists.
//var ErrNotImplemented = errors.New("lexer not implemented")

// Token represents a lexical token produced by the lexer.
type Token struct {
	Kind   string
	Lexeme string
	Line   int
	Column int
}

// Lexer tokenizes FRACTRAN++ source text.
type Lexer struct {
	input string
	pos   int
}

func (l *Lexer) skipUntilSeparator() {
	for l.pos < len(l.input) && !isWhitespace(l.peek()) && l.peek() != ',' {
		l.advance()
	}
}

func (l *Lexer) peek() rune {
	if l.pos >= len(l.input) {
		return 0
	}
	return rune(l.input[l.pos])
}

func (l *Lexer) advance() {
	l.pos++
}

func (l *Lexer) readInteger() string {
	start := l.pos
	for l.pos < len(l.input) && isDigit(l.peek()) {
		l.advance()
	}
	return l.input[start:l.pos]
}

// New returns a new lexer instance.
func New() *Lexer {
	return &Lexer{}
}

func (l *Lexer) Tokenize(input string) ([]Token, error) {
	l.input = input
	l.pos = 0

	var tokens []Token

	for l.pos < len(l.input) {
		ch := l.peek()
		// Skip separators
		if isWhitespace(ch) || ch == ',' {
			l.advance()
			continue
		}

		// Integer
		if isDigit(ch) {
			value := l.readInteger()

			tokens = append(tokens, Token{
				Kind:   "Integer",
				Lexeme: value,
				Line:   1,
				Column: 1,
			})
			continue
		}

		// Slash
		if ch == '/' {
			l.advance()
			tokens = append(tokens, Token{
				Kind:   "Slash",
				Lexeme: "/",
				Line:   1,
				Column: 1,
			})
			continue
		}

		// Everything else is a comment and we decided to
		l.skipUntilSeparator()
	}

	return tokens, nil
}

func isDigit(r rune) bool {
	return r >= '0' && r <= '9'
}
func isWhitespace(r rune) bool {
	return r == ' ' || r == '\t' || r == '\n' || r == '\r'
}
