package ast

import (
	"bytes"
	"golite/token"
)

type Typedecl struct {
	*token.Token
	name   string
	fields []*Decl
}

func NewtypeDecl(token *token.Token, name string, fields []*Decl) *Typedecl {
	return &Typedecl{token, name, fields}
}

func (td *Typedecl) String() {
	var out bytes.Buffer
	out.WriteString("[NODE TYPEDECLARATION] ")
	out.WriteString("type" + td.name + "struct" + "{")
	for _, field := range td.fields {
		out.WriteString(field.String())
		out.WriteString("; ")

	}
	out.WriteString("}" + ";\n")
}
