package ast

import "github.com/rizasghari/ari/token"

// Abstract Syntax Tree

// The AST we are going to construct consists solely of Nodes
// that are connected to each other - it’s a tree after all.
// Some of these nodes implement the Statement and some the Expression interface.

// Every node in our AST has to implement the Node interface,
// meaning it has to provide a TokenLiteral() method that returns
// the literal value of the token it’s associated with.
type Node interface {
	// TokenLiteral() will be used only for debugging and testing.
	TokenLiteral() string
}

type Statement interface {
	Node
	statementNode()
}

type Expression interface {
	Node
	expressionNode()
}

// This Program node is going to be the root node of every AST our parser produces.
// Every valid Arı program is a series of statements.
// These statements are contained in the Pro- gram.Statements, which is just a slice
// of AST nodes that implement the Statement interface.
type Program struct {
	Statements []Statement
}

func (p *Program) TokenLiteral() string {
	if len(p.Statements) > 0 {
		return p.Statements[0].TokenLiteral()
	} else {
		return ""
	}
}

type Identifier struct {
	Token token.Token // the token.IDENT token Value string
	Value string
}

func (i *Identifier) expressionNode() {}

func (i *Identifier) TokenLiteral() string {
	return i.Token.Literal
}

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
