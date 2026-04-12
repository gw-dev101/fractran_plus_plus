package parser

import (
	"math/big"

	"github.com/gw-dev101/fractran_plus_plus/internal/ast"
	"github.com/gw-dev101/fractran_plus_plus/internal/lexer"
)

// Parser turns tokens into an AST.
type Parser struct{}

// New returns a new parser instance.
func New() *Parser {
	return &Parser{}
}

func (p *Parser) Parse(tokens []lexer.Token) (*ast.Program, error) {
	//build an AST (array of fractions) from the int and slash tokens
	program := &ast.Program{}
	for i := 0; i < len(tokens); {
		if tokens[i].Kind == "Integer" {
			if i+2 < len(tokens) && tokens[i+1].Kind == "Slash" && tokens[i+2].Kind == "Integer" {
				//parse the numerator and denominator as big.Int
				numerator := new(big.Int)
				denominator := new(big.Int)
				numerator.SetString(tokens[i].Lexeme, 10)
				denominator.SetString(tokens[i+2].Lexeme, 10)
				program.Statements = append(program.Statements, ast.Fraction{
					Numerator:   numerator,
					Denominator: denominator,
				})
				i += 3
			} else {
				return nil, &ParseError{Message: "Expected a slash and another integer after an integer"}
			}
		} else {
			return nil, &ParseError{Message: "Expected an integer"}
		}
	}
	return program, nil
}

// ParseError represents an error that occurred during parsing.
type ParseError struct {
	Message string
}

func (e *ParseError) Error() string {
	return e.Message
}
