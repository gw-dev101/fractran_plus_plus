package main

//import everything we used to test the compiler, interpreter, lexer, and parser
import (
	"fmt"

	"github.com/gw-dev101/fractran_plus_plus/internal/frac_math"
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
	//	c := compiler.New()
	//	compiledProgram, err := c.Compile(program)
	//	if err != nil {
	//		panic(err)
	//	}
	//	fmt.Println("Compiled Program:", compiledProgram)
	i := interpreter.New()
	// primegame
	n := frac_math.FromFactors(map[int]int{2: 1}) // start with 2^1

	for step := 0; step < 100000; step++ {
		var ok bool
		var err error

		n, ok, err = i.Step(program, n)
		if err != nil {
			panic(err)
		}
		if !ok {
			fmt.Println("halted")
			break
		}

		if n.IsPowerOfTwo() {
			fmt.Println("prime:", exponent(n))
		}
	}
	if err != nil {
		panic(err)
	}
	finalResult, ok, err := i.Step(program, n)
	if err != nil {
		panic(err)
	}
	if !ok {
		fmt.Println("Program halted with final result:", finalResult)
	} else {
		fmt.Println("Program did not halt, final state after 100000 steps:", finalResult)
	}
	fmt.Println("Program executed successfully")
}

// Helper function to get the exponent of 2 in a MyInt, assuming it's a power of two
func exponent(n *frac_math.MyInt) int {
	if exp, ok := n.Factors()[2]; ok {
		return exp
	}
	return 0
}
