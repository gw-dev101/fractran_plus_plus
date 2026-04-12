package main

//import everything we used to test the compiler, interpreter, lexer, and parser
import (
	"fmt"
	"github.com/gw-dev101/fractran_plus_plus/internal/compiler"
	"github.com/gw-dev101/fractran_plus_plus/internal/interpreter"
	"github.com/gw-dev101/fractran_plus_plus/internal/lexer"
	"github.com/gw-dev101/fractran_plus_plus/internal/parser"
)

func main() {
	//try to run some tests
	sampleCode := `
		# This is a sample FRACTRAN++ proram
		# Its based on the PRIMEGAME example from the FRACTRAN wikipedia page
		17/91
		78/85
		19/51
		23/38
		29/33
		77/29
		95/23
		77/19
		1/17
		11/13
		13/11
		15/14
		15/2
		55/1
	`

	l := lexer.New()
	tokens, err := l.Tokenize(sampleCode)
	if err != nil {
		panic(err)
	}
	fmt.Println("Tokens:", tokens)
	p := parser.New()
	program, err := p.Parse(tokens)
	if err != nil {
		panic(err)
	}
	fmt.Println("Parsed Program:", program)
	c := compiler.New()
	compiledProgram, err := c.Compile(program)
	if err != nil {
		panic(err)
	}
	fmt.Println("Compiled Program:", compiledProgram)
	i := interpreter.New()
	result, err := i.Execute(program)
	if err != nil {
		panic(err)
	}
	fmt.Println("Execution Result:", result)
	fmt.Println("Program executed successfully")
}
