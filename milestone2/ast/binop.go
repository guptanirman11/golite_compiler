package ast

import (
	"bytes"
	"golite/token"
)

type BinOpExpr struct {
	*token.Token            // The token information
	operator     Operator   // The operator for the binary expression
	right        Expression // The right operand expression
	left         Expression // The left operand expression
}

func NewBinOp(op Operator, right, left Expression, token *token.Token) *BinOpExpr {
	return &BinOpExpr{token, op, right, left}
}

func (binOp *BinOpExpr) String() string {

	var out bytes.Buffer

	out.WriteString("(")
	out.WriteString(binOp.left.String())
	out.WriteString(" " + OpSring(binOp.operator) + " ")
	out.WriteString(binOp.right.String())
	out.WriteString(")")

	return out.String()
}
