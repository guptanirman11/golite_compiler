package ast

import (
	"bytes"
	"go/token"
)

type Program struct {
	stmts []Statement
	*token.Token
	// never do []*Statement
	// but always use bin = &BinaryExoression{....} as this a statement itself
}

func newProgram(stmts []Statement, token *token.Token) *Program {
	// return &Program{nil}
	return &Program{stmts, token}
}

func (p *Program) String() string {

	var out bytes.Buffer
	for _, stmt := range p.stmts {
		out.WriteString(stmt.String())
		out.WriteString("\n")
	}
	return out.String()
}
