package ast

import (
	"fmt"
	"golite/token"
)

type BoolLiteral struct {
	*token.Token      // The token for the integer literal
	Value        bool // The value for the integer literal
}

func NewBoolLit(value bool, token *token.Token) *BoolLiteral {
	return &BoolLiteral{token, value}
}
func (bl *BoolLiteral) String() string {
	return fmt.Sprintf("%v", bl.Value)
}
