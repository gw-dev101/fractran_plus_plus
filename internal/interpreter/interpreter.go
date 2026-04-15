package interpreter

import (
	"fmt"

	"github.com/gw-dev101/fractran_plus_plus/internal/ast"
	"github.com/gw-dev101/fractran_plus_plus/internal/frac_math"
)

// Result describes the outcome of interpreting a program.
type Result struct {
	Steps  int
	Halted bool
}

func (r Result) String() string {
	return "Result{Steps: " + fmt.Sprintf("%d", r.Steps) + ", Halted: " + fmt.Sprintf("%t", r.Halted) + "}"
}

// Interpreter executes FRACTRAN++ programs.
type Interpreter struct{}

// New returns a new interpreter instance.
func New() *Interpreter {
	return &Interpreter{}
}

func (i *Interpreter) Execute(program *ast.Program, input *frac_math.MyInt) (Result, error) {
	state := input.Clone()
	steps := 0
	for {
		newState, changed, err := i.Step(program, state)
		if err != nil {
			return Result{}, err
		}
		if !changed {
			return Result{Steps: steps, Halted: true}, nil
		}
		state.Set(newState)
		steps++
	}

}
func (i *Interpreter) Step(program *ast.Program, state *frac_math.MyInt) (*frac_math.MyInt, bool, error) {
	for _, stmt := range program.Statements {
		switch s := stmt.(type) {
		case ast.Fraction:
			// Multiply state by the fraction and check if it's an integer
			numerator := s.Numerator
			denominator := s.Denominator

			newState := state.Clone()
			newState.Multiply(numerator)
			if !newState.Divide(denominator) {
				continue // not an integer, try next statement
			}
			return newState, true, nil

		default:
			return nil, false, fmt.Errorf("unknown statement type: %T", stmt)
		}
	}
	return state, false, nil
}