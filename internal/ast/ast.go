package ast

import (
	"github.com/gw-dev101/fractran_plus_plus/internal/frac_math"
)

// Program is the root syntax tree for a FRACTRAN++ source file.
type Program struct {
	Statements []Statement
}

// Statement is the common interface for all AST statement nodes.
type Statement interface {
	isStatement()
}

type Fraction struct {
	Numerator   *frac_math.MyInt
	Denominator *frac_math.MyInt
}

func (Fraction) isStatement() {}
