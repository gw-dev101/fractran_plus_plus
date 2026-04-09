package ast

// Program is the root syntax tree for a FRACTRAN++ source file.
type Program struct {
	Statements []Statement
}

// Statement is the common interface for all AST statement nodes.
type Statement interface {
	isStatement()
}

// PlaceholderStatement exists so the AST package compiles before the language
// model is fully designed.
type PlaceholderStatement struct{}

func (PlaceholderStatement) isStatement() {}
