package ast

import (
	"bytes"
	"golite/token"
)

type Conditional struct {
	*token.Token
	block     *Block
	elseBlock *Block
	exp       Expression
}

func NewConditional(token *token.Token, block *Block, elseBlock *Block, expression Expression) *Conditional {
	return &Conditional{token, block, elseBlock, expression}
}

func (c *Conditional) String() string {
	var out bytes.Buffer
	out.WriteString("[NODE CONDITIONAL] ")
	out.WriteString("if(")
	out.WriteString(c.exp.String())
	out.WriteString(")" + c.block.String())

	if c.elseBlock != nil {
		out.WriteString("else" + c.elseBlock.String())
	}

	return out.String()
}
