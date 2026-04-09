package compiler

import (
	"errors"

	"github.com/gw-dev101/fractran_plus_plus/internal/ast"
)

// ErrNotImplemented is returned by the stub compiler until code generation exists.
var ErrNotImplemented = errors.New("compiler not implemented")

// Compiler turns an AST into a target representation.
type Compiler struct{}

// New returns a new compiler instance.
func New() *Compiler {
	return &Compiler{}
}

// Compile is a stub placeholder for the real compiler implementation.
func (c *Compiler) Compile(program *ast.Program) ([]byte, error) {
	return nil, ErrNotImplemented
}
