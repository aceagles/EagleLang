package parser

import (
	"strconv"

	"github.com/aceagles/EagleLang/ast"
	"github.com/aceagles/EagleLang/lexer"
	"github.com/aceagles/EagleLang/token"
)

type Parser struct {
	l *lexer.Lexer

	currToken token.Token
	peekToken token.Token
}

func New(l *lexer.Lexer) *Parser {
	p := &Parser{l: l}

	p.nextToken()
	p.nextToken()

	return p
}

func (p *Parser) nextToken() {
	p.currToken = p.peekToken
	p.peekToken = p.l.NextToken()
}

func (p *Parser) expectPeek(t string) bool {
	if string(p.peekToken.Type) == t {
		p.nextToken()
		return true
	} else {
		return false
	}
}

func (p *Parser) ParseProgram() *ast.Program {
	program := &ast.Program{}
	statements := []ast.Statement{}

	for p.currToken.Type != token.EOF {
		switch p.currToken.Type {
		case token.LET:
			stmt := p.parseLetStatement()
			statements = append(statements, stmt)
		}
		p.nextToken()
	}
	program.Statements = statements
	return program
}

func (p *Parser) parseLetStatement() *ast.LetStatement {
	stmt := ast.LetStatement{Token: p.currToken}

	p.nextToken()
	stmt.Name = &ast.Identifier{
		Token: p.currToken,
		Value: p.currToken.Literal,
	}

	if !p.expectPeek(token.ASSIGN) {
		return nil
	}

	stmt.Value = p.parseIntegerLiteral()

	for p.currToken.Type != token.SEMICOLON {
		p.nextToken()
	}

	return &stmt
}

func (p *Parser) parseIntegerLiteral() *ast.IntegerLiteral {
	p.nextToken()
	value, _ := strconv.Atoi(p.currToken.Literal)
	return &ast.IntegerLiteral{
		Token: p.currToken,
		Value: int64(value),
	}
}
