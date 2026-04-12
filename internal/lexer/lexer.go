package lexer

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
// Tokenize takes a string of FRACTRAN++ source code and returns a slice of tokens or an error if tokenization fails.
// The lexer takes care of skipping comments(anything that is not part of the grammar) and separators (whitespace and commas).
// the current implementation only supports pure FRACTRAN programs , but will be extended to support the full FRACTRAN++ syntax in the future.
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
