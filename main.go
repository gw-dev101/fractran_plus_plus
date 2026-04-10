package main
//import everything we used to test the compiler, interpreter, lexer, and parser
import (
    "github.com/gw-dev101/fractran_plus_plus/internal/compiler"
    "github.com/gw-dev101/fractran_plus_plus/internal/interpreter"
    "github.com/gw-dev101/fractran_plus_plus/internal/lexer"
    "github.com/gw-dev101/fractran_plus_plus/internal/parser"
)

func main() {
    //try to run some tests
    compiler.TestCompile(nil)
    interpreter.TestExecute(nil)
    lexer.TestTokenize(nil)
    parser.TestParse(nil)   
}