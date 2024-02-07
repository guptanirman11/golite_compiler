package types

// Type includes information about working on Types
type Type interface {
	String() string
}

type IntTy struct{}

func (intTy *IntTy) String() string {
	return "int"
}

type BoolTy struct{}

func (boolTy *BoolTy) String() string {
	return "bool"
}
func StringToType(ty string) Type {
	if ty == "int" {
		return IntTySig
	} else if ty == "bool" {
		return BoolTySig
	}
	panic("Not found type")
}

var IntTySig *IntTy
var BoolTySig *BoolTy

// The init() function will only be called once per package. This is where you can setup singletons for types
func init() {
	IntTySig = &IntTy{}
	BoolTySig = &BoolTy{}
}
