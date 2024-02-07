package ast

import (
	"bytes"
	"golite/token"
	"golite/types"
)

type Declaration struct {
	*token.Token
	Ty  types.Type
	Ids []string
}

func NewDeclarationStm(ids []string, ty types.Type, token *token.Token) *Declaration {
	return &Declaration{token, ty, ids}
}
func (d *Declaration) String() string {

	var out bytes.Buffer
	out.WriteString("Node Declaration")
	out.WriteString("var")
	for _, id := range d.Ids {
		out.WriteString(id)
		out.WriteString(", ")
	}
	out.WriteString(d.Ty.String() + ";\n")

	return out.String()

}
