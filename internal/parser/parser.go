package parser

import (
	"errors"

	"github.com/gw-dev101/fractran_plus_plus/internal/ast"
	"github.com/gw-dev101/fractran_plus_plus/internal/lexer"
)

// ErrNotImplemented is returned by the stub parser until parsing exists.
var ErrNotImplemented = errors.New("parser not implemented")

// Parser turns tokens into an AST.
type Parser struct{}

// New returns a new parser instance.
func New() *Parser {
	return &Parser{}
}

// Parse is a stub placeholder for the real parser implementation.
func (p *Parser) Parse(tokens []lexer.Token) (*ast.Program, error) {
	return nil, ErrNotImplemented
}
