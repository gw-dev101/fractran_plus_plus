package main
//import everything we used to test the compiler, interpreter, lexer, and parser
import (
    "internal/compiler"
    "internal/interpreter"
    "internal/lexer"
    "internal/parser"
)

func main() {
    //try to run some tests
    compiler.TestCompile(nil)
    interpreter.TestExecute(nil)
    lexer.TestTokenize(nil)
    parser.TestParse(nil)   
}