package ast

import "github.com/rizasghari/ari/token"

type Identifier struct {
	Token token.Token // the token.IDENT token Value string
	Value string
}
func (i *Identifier) expressionNode() {}

func (i *Identifier) TokenLiteral() string {
	return i.Token.Literal
}

func (i *Identifier) String() string {
	return i.Value
}