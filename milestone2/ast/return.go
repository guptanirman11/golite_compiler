package ast

import (
	"bytes"
	"golite/token"
)

type Return struct {
	*token.Token
	expr Expression
}

func NewReturn(token *token.Token, expression Expression) *Return {
	return &Return{token, expression}

}

func (r *Return) STring() string {
	var out bytes.Buffer
	if r.expr != nil {
		out.WriteString("return " + r.expr.String() + ";")
	} else {
		out.WriteString("return;")
	}

	return out.String()
}
