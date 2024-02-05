package symboltable

import (
	"fmt"
	"go/types"
)

type varScope int

const (
	GLOBAL varScope = iota
	LOCAL
)

type VarEntry struct {
	Name  string
	Ty    types.Type
	Scope varScope
}

// type FuncEntry struct{
// 	*token.Token
// 	Name string
// 	RetTy types.Type
// 	Parameters []*VarEntry
// 	Variables *SymbolTable[*&VarEntry]
// }

// type SymbolTables struct {
// 	Funcs *SymbolTable[*FuncEntry]
// 	Globals *SymbolTable[*VarEntry]
// }

// func NewSymbolTAbles() *SymbolTables{
// 	return &SymbolTables{NewSymbolTAble[*VarEntry](nil)}
// }

type SymbolTable[T fmt.Stringer] struct {
	parent *SymbolTable[T]
	table  map[string]T
}
