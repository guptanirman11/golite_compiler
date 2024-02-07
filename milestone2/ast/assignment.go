package ast

import (
	"bytes"
	"golite/token"
)

type Assignment struct {
	*token.Token
	variable *Lvalue    // The lvalue for the assignment statement
	right    Expression // The value assigned to the variable of this statements
}

func NewAssignStm(variable *Lvalue, right Expression, token *token.Token) *Assignment {
	return &Assignment{token, variable, right}
}
func (a *Assignment) String() string {

	var out bytes.Buffer
	out.WriteString("[NODE ASSIGNMENT] ")

	out.WriteString(a.variable.String())
	out.WriteString("=")
	out.WriteString(a.right.String())
	out.WriteString(";")

	return out.String()

}
