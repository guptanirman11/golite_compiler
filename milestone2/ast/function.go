package ast

import (
	"bytes"
	"golite/token"
	"golite/types"
)

type Function struct {
	*token.Token
	Name       string
	rType      types.Type
	stmts      []Statement
	parameters []*Decl
	Decls      []*Declaration
}

func NewFunction(name string, params []*Decl, retTy types.Type, decls []*Declaration, stmts []Statement, token *token.Token) *Function {
	return &Function{token, name, retTy, stmts, params, decls}
}

func (f *Function) String() string {

	var out bytes.Buffer
	out.WriteString("Node Function")

	out.WriteString("func" + f.Name + "(")

	for _, param := range f.parameters {
		out.WriteString(param.Name)
		out.WriteString(", ")
	}
	out.WriteString(")" + f.rType.String() + "{")

	for _, declaration := range f.Decls {
		out.WriteString(declaration.String())
	}

	for _, st := range f.stmts {
		out.WriteString(st.String())
	}
	out.WriteString("}")

	return out.String()
}
