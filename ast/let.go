package ast

import (
	"bytes"

	"github.com/rizasghari/ari/token"
)

// LetStatement has the fields we need:
// Name to hold the identifier of the binding and
// Value for the expression that produces the value.
// The two methods statementNode and TokenLiteral
// satisfy the Statement and Node interfaces respectively.
type LetStatement struct {
	Token token.Token // the token.LET token Name *Identifier
	Name  *Identifier
	Value Expression
}

func (ls *LetStatement) statementNode() {}

func (ls *LetStatement) TokenLiteral() string {
	return ls.Token.Literal
}

func (ls *LetStatement) String() string {
	var out bytes.Buffer

	out.WriteString(ls.TokenLiteral() + " ")
	out.WriteString(ls.Name.String())
	out.WriteString(" = ")

	if ls.Value != nil {
		out.WriteString(ls.Value.String())
	}

	out.WriteString(";")

	return out.String()
}