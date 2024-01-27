package token

import "fmt"

// type TokenType string

const (
	SEMICOLON = "SEMICOLON"
	PLUS      = "PLUS"
	MINUS     = "MINUS"
	DIVIDE    = "DIVIDE"
	MULTIPLY  = "MULTIPLY"
	EQUAL     = "EQUAL"

	PRINT   = "PRINT"
	LET     = "LET"
	SPACE   = "SPACE"
	NEWLINE = "NEWLINE"
	ILLEGAL = "ILLEGAL"

	INT_LIT = "INT"
	ID      = "ID"
)

type Token struct {
	Type    string
	Literal string
	Line    int
}

func (t *Token) Print() {
	if t.Type == "EOF" || t.Type == "INVALID" || t.Type == "NEWLINE" || t.Type == "SPACE" {
		return
	} else if t.Type == "ID" {
		fmt.Printf("%v(%d, \"%v\")\n", t.Type, t.Line, t.Literal)
		return
	} else if t.Type == "INT" {
		fmt.Printf("%v(%d, %v)\n", t.Type, t.Line, t.Literal)
		return
	}
	fmt.Printf("%v(%d)\n", t.Type, t.Line)
}
