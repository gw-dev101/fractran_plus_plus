package ast

import (
	"math/big"
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
	Numerator   *big.Int
	Denominator *big.Int
}

func (Fraction) isStatement() {}
