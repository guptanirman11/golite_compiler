package main

import (
	"fmt"
	"golite/lexer"
	"golite/parser"
	"os"
)

func main() {
	inputSourcePath := os.Args[1]
	lexer := lexer.NewLexer(inputSourcePath)
	// fmt.Println(lexer.GetTokenStream())
	parser := parser.NewParser(lexer)
	fmt.Println(parser.Parse())
}
