package interpreter

import (
	"fmt"
	"math/big"

	"github.com/gw-dev101/fractran_plus_plus/internal/ast"
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

func (i *Interpreter) Execute(program *ast.Program, input *big.Int) (Result, error) {
	state := new(big.Int).Set(input)
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
func (i *Interpreter) Step(program *ast.Program, state *big.Int) (*big.Int, bool, error) {
	for _, stmt := range program.Statements {
		switch s := stmt.(type) {
		case ast.Fraction:
			var tmp big.Int
			tmp.Mod(state, s.Denominator)
			if tmp.Sign() == 0 {
				newState := new(big.Int).Mul(state, s.Numerator)
				newState.Div(newState, s.Denominator)
				return newState, true, nil
			}
		default:
			return nil, false, fmt.Errorf("unknown statement type: %T", stmt)
		}
	}
	return state, false, nil
}