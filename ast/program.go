package ast

// This Program node is going to be the root node of every AST our parser produces.
// Every valid Arı program is a series of statements.
// These statements are contained in the Program.
// Statements, which is just a slice of AST nodes that implement the Statement interface.
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