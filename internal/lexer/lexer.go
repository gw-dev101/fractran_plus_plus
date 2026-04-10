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
type Lexer struct{
	input string
	pos int
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
        if isWhitespace(l.peek()) || l.peek() == ',' {
            l.advance()
            continue
        }

        if isDigit(l.peek()) {
            left := l.readInteger()

            if l.peek() == '/' {
                l.advance()

                if !isDigit(l.peek()) {
                    return nil, fmt.Errorf("expected digit after '/'")
                }

                right := l.readInteger()

                tokens = append(tokens, Token{
                    Kind:   "Fraction",
                    Lexeme: left + "/" + right,
                })
            } else {
                tokens = append(tokens, Token{
                    Kind:   "Integer",
                    Lexeme: left,
                })
            }

            continue
        }

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