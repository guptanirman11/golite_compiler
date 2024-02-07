package ast

import (
	"bytes"
	"golite/token"
)

type Loop struct {
	*token.Token
	expr  Expression
	block *Block
}

func NewLoop(token *token.Token, expression Expression, block *Block) *Loop {
	return &Loop{token, expression, block}
}

func (l *Loop) String() string {
	var out bytes.Buffer
	out.WriteString("[NODE LOOP] ")
	out.WriteString("for(" + l.expr.String() + ")")
	out.WriteString(l.block.String())
	return out.String()
}
