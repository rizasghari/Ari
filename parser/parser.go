package parser

import (
	"github.com/rizasghari/ari/ast"
	"github.com/rizasghari/ari/lexer"
	"github.com/rizasghari/ari/token"
)

// Recursive Descent Parser

// The Parser has three fields: l, curToken and peekToken.
// l is a pointer to an instance of the lexer, on which we
// repeatedly call NextToken() to get the next token in the input.
// curToken and peekToken act exactly like the two “pointers” our
// lexer has: position and peekPosition. But instead of pointing to a
// character in the input, they point to the current and the next token.
// Both are important: we need to look at the curToken, which is the current
// token under examination, to decide what to do next, and we also need peekToken
// for this decision if curToken doesn’t give us enough information.
// Think of a single line only containing 5;. Then curToken is a token.INT
// and we need peekToken to decide whether we are at the end of the line or
// if we are at just the start of an arithmetic expression.
type Parser struct {
	l         *lexer.Lexer
	curToken  token.Token
	peekToken token.Token
}

func New(l *lexer.Lexer) *Parser {
	p := &Parser{l: l}
	p.nextToken()
	p.nextToken()
	return p
}

func (p *Parser) nextToken() {
	p.curToken = p.peekToken
	p.peekToken = p.l.NextToken()
}

func (p *Parser) ParseProgram() *ast.Program {
	// The first thing ParseProgram does is construct the root node of the AST, an *ast.Program.
	program := &ast.Program{}
	program.Statements = []ast.Statement{}

	// It then iterates over every token in the input until it encounters an token.EOF token.
	// It does this by repeatedly calling nextToken, which advances both p.curToken and p.peekToken.
	// In every iteration it calls parseStatement, whose job it is to parse a statement.
	// If parseStatement returned something other than nil, a ast.Statement, its return value is added
	// to Statements slice of the AST root node. When nothing is left to parse the *ast.Program root node is returned.
	for p.curToken.Type != token.EOF {
		stmt := p.parseStatement()
		if stmt != nil {
			program.Statements = append(program.Statements, stmt)
		}
		p.nextToken()
	}

	return program
}

func (p *Parser) parseStatement() ast.Statement {
	switch p.curToken.Type {
	case token.LET:
		return p.parseLetStatement()
	default:
		return nil
	}
}

func (p *Parser) parseLetStatement() *ast.LetStatement {
	stmt := &ast.LetStatement{Token: p.curToken}
	if !p.expectPeek(token.IDENT) {
		return nil
	}
	stmt.Name = &ast.Identifier{Token: p.curToken, Value: p.curToken.Literal}
	if !p.expectPeek(token.ASSIGN) {
		return nil
	}
	// TODO: We're skipping the expressions until we encounter a semicolon
	for !p.curTokenIs(token.SEMICOLON) {
		p.nextToken()
	}
	return stmt
}

func (p *Parser) curTokenIs(t token.TokenType) bool {
	return p.curToken.Type == t
}

func (p *Parser) peekTokenIs(t token.TokenType) bool {
	return p.peekToken.Type == t
}

func (p *Parser) expectPeek(t token.TokenType) bool {
	if p.peekTokenIs(t) {
		p.nextToken()
		return true
	} else {
		return false
	}
}
