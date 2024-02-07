package ast

import (
	"bytes"
	"golite/token"
)

type Block struct {
	*token.Token
	stmts []Statement
}

func NewBlock(token *token.Token, stmts []Statement) *Block {
	return &Block{token, stmts}
}

func (b *Block) String() string {
	var out bytes.Buffer
	out.WriteString("[NODE BLOCK] ")

	out.WriteString("{\n")
	for _, st := range b.stmts {
		out.WriteString(st.String() + "\n")
	}
	out.WriteString("}\n")
	return out.String()
}
