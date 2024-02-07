package ast

import (
	"bytes"
	"golite/types"

	//"golite/cfg"
	// st "golite/symboltable"
	"golite/token"
	//"fmt"
)

type Lvalue struct {
	*token.Token
	Ids      []string
	Ty       types.Type
	StructTy types.Type
	FieldIdx int
	// Scope    st.VarScope
}

// func NewLvalue(ids []string, token *token.Token) *Lvalue {
// 	return &Lvalue{token, ids, nil, nil, 0, st.GLOBAL}
// }

func (l *Lvalue) String() string {
	var out bytes.Buffer
	out.WriteString("[NODE LVALUE] ")

	for _, id := range l.Ids[:len(l.Ids)-1] {
		out.WriteString(id + ".")
	}

	out.WriteString(l.Ids[len(l.Ids)-1])

	return out.String()
}
