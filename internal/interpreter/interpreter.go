package interpreter

import (
	"errors"

	"github.com/gw-dev101/fractran_plus_plus/internal/ast"
)

// ErrNotImplemented is returned by the stub interpreter until execution exists.
var ErrNotImplemented = errors.New("interpreter not implemented")

// Result describes the outcome of interpreting a program.
type Result struct {
	Steps  int
	Halted bool
}

// Interpreter executes FRACTRAN++ programs.
type Interpreter struct{}

// New returns a new interpreter instance.
func New() *Interpreter {
	return &Interpreter{}
}

// Execute is a stub placeholder for the real interpreter implementation.
func (i *Interpreter) Execute(program *ast.Program) (Result, error) {
	return Result{}, ErrNotImplemented
}
