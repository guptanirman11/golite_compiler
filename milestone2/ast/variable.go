package ast

import (
	"golite/token"
)

type Variable struct {
	*token.Token
	Identifier string // The variable name
}

func NewVariable(identifier string, token *token.Token) *Variable {
	return &Variable{token, identifier}
}
func (v *Variable) String() string { return v.Identifier }
