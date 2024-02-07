package ast

import (
	"fmt"
	"golite/token"
)

type IntLiteral struct {
	*token.Token       // The token for the integer literal
	Value        int64 // The value for the integer literal
}

func NewIntLit(value int64, token *token.Token) *IntLiteral {
	return &IntLiteral{token, value}
}
func (il *IntLiteral) String() string {
	return fmt.Sprintf("%d", il.Value)
}
