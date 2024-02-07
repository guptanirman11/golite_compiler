package ast

import (
	"bytes"
	"golite/token"
)

type Print struct {
	*token.Token              // The token for this print statement
	str          string       //string to print
	expr         []Expression // The value being printed out in the string
}

func NewPrintStm(expr []Expression, token *token.Token, str string) *Print {
	return &Print{token, str, expr}
}
func (p *Print) String() string {

	var out bytes.Buffer

	out.WriteString("printf(" + p.str)
	for _, exp := range p.expr {
		out.WriteString(exp.String())
	}

	out.WriteString(");")
	return out.String()

}
