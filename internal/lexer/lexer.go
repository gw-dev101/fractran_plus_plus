package lexer

import "errors"

// ErrNotImplemented is returned by the stub lexer until tokenization exists.
var ErrNotImplemented = errors.New("lexer not implemented")

// Token represents a lexical token produced by the lexer.
type Token struct {
	Kind   string
	Lexeme string
	Line   int
	Column int
}

// Lexer tokenizes FRACTRAN++ source text.
type Lexer struct{}

// New returns a new lexer instance.
func New() *Lexer {
	return &Lexer{}
}

// Tokenize is a stub placeholder for the real lexer implementation.
func (l *Lexer) Tokenize(source string) ([]Token, error) {
	return nil, ErrNotImplemented
}
