package ast

import "github.com/rizasghari/ari/token"

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