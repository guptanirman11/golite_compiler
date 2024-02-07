package ast

import (
	"bytes"
	"golite/token"
)

type Delete struct {
	*token.Token
	Expression Expression
}

func NewDelete(token *token.Token, expression Expression) *Delete {
	return &Delete{token, expression}

}

func (d *Delete) String() string {
	var out bytes.Buffer
	out.WriteString("[NODE DELETE] ")

	out.WriteString("delete " + d.Expression.String() + ";")

	return out.String()
}
