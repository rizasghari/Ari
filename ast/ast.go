package ast

// AST: Abstract Syntax Tree

// The AST we are going to construct consists solely of Nodes
// that are connected to each other - it’s a tree after all.
// Some of these nodes implement the Statement and some the Expression interface.

// Every node in our AST has to implement the Node interface,
// meaning it has to provide a TokenLiteral() method that returns
// the literal value of the token it’s associated with.
type Node interface {
	// TokenLiteral() will be used only for debugging and testing.
	TokenLiteral() string
	String() string
}

type Statement interface {
	Node
	statementNode()
}

type Expression interface {
	Node
	expressionNode()
}