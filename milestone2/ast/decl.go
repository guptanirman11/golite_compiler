package ast

import (
	"bytes"
	"go/types"
	"golite/token"
	// "golite/types"
)

type Decl struct {
	*token.Token
	Name string // The lvalue for the assignment statement
	ty   types.Type
	// right    Expression // The value assigned to the variable of this statements
}

func NewDeclStm(name string, ty types.Type, right Expression, token *token.Token) *Decl {
	return &Decl{token, name, ty}
}
func (d *Decl) String() string {

	var out bytes.Buffer
	out.WriteString("Node Decl")
	out.WriteString(d.Name + " " + d.ty.String() + "\n")
	out.WriteString(";")

	return out.String()

}
