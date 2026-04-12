package main

//import everything we used to test the compiler, interpreter, lexer, and parser
import (
	"fmt"
	"math/big"

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
	n := big.NewInt(2)

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

		if isPowerOfTwo(n) {
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
func isPowerOfTwo(n *big.Int) bool {
	if n.Sign() <= 0 {
		return false
	}

	tmp := new(big.Int).Set(n)
	two := big.NewInt(2)
	zero := big.NewInt(0)

	for {
		mod := new(big.Int).Mod(tmp, two)
		if mod.Cmp(zero) != 0 {
			break
		}
		tmp.Div(tmp, two)
	}

	return tmp.Cmp(big.NewInt(1)) == 0
}

func exponent(n *big.Int) int {
	tmp := new(big.Int).Set(n)
	two := big.NewInt(2)

	exp := 0
	for tmp.Cmp(big.NewInt(1)) > 0 {
		tmp.Div(tmp, two)
		exp++
	}
	return exp
}
