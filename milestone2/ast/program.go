package ast

import (
	"bytes"
	"golite/token"
)

type Program struct {
	*token.Token
	stmts []Statement
}

func NewProgram(stmts []Statement, token *token.Token) *Program {
	return &Program{token, stmts}
}
func (p *Program) String() string {

	var out bytes.Buffer
	for _, stmt := range p.stmts {
		out.WriteString(stmt.String())
		out.WriteString("\n")
	}
	return out.String()
}
