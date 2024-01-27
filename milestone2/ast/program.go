package ast

type Program struct {
	stmts []Statement
	// never do []*Statement
	// but always use bin = &BinaryExoression{....} as this a statement itself
}

func newProgram( /*..*/ ) *Program {
	return &Program{nil}
}
