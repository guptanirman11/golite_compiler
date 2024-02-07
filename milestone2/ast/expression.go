package ast

type Operator int

const (
	ADD Operator = iota
	SUB
	GT
	GE
	LT
	LE
)

func StrToOp(op string) Operator {
	switch op {
	case "+":
		return ADD
	case "-":
		return SUB
	case ">":
		return GT
	case ">=":
		return GE
	case "<":
		return LT
	case "<=":
		return LE
	}
	panic("Error: Could not determine operator")
}
func OpSring(op Operator) string {

	switch op {
	case ADD:
		return "+"
	case SUB:
		return "-"
	case GT:
		return ">"
	case GE:
		return ">="
	case LT:
		return "<"
	case LE:
		return "<="
	}
	panic("Error: Could not determine operator")
}

// Expression is the base node for all expression specific nodes (i.e., every expression sub node implements this
// interface)
type Expression interface {
	String() string
}
