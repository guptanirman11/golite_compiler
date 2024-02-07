package ast

import (
	"bytes"
	"golite/token"
)

type Read struct {
	*token.Token
	Lvalue *Lvalue
}

func NewRead(token *token.Token, lvalues *Lvalue) *Read {
	return &Read{token, lvalues}

}

func (r *Read) String() string {
	var out bytes.Buffer
	out.WriteString("[NODE READ] ")

	out.WriteString("scan" + r.Lvalue.String() + ";")
	return out.String()
}
