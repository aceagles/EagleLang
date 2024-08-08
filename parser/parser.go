package parser

import (
	"strconv"

	"github.com/aceagles/EagleLang/ast"
	"github.com/aceagles/EagleLang/lexer"
	"github.com/aceagles/EagleLang/token"
)

type (
	prefixParseFn func() ast.Expression
	infixParseFn  func(ast.Expression) ast.Expression
)
type Parser struct {
	l *lexer.Lexer

	currToken token.Token
	peekToken token.Token
	errors    []string

	prefixParseFns map[token.TokenType]prefixParseFn
	infixParseFns  map[token.TokenType]infixParseFn
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
		p.errors = append(p.errors, "expected next token to be "+t+" but got "+string(p.peekToken.Type))
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

	if !p.expectPeek(token.IDENT) {
		return nil
	}

	stmt.Name = &ast.Identifier{
		Token: p.currToken,
		Value: p.currToken.Literal,
	}

	if !p.expectPeek(token.ASSIGN) {
		return nil
	}

	p.expectPeek(token.INT)
	stmt.Value = p.parseIntegerLiteral()

	for p.currToken.Type != token.SEMICOLON {
		p.nextToken()
	}

	return &stmt
}

func (p *Parser) parseIntegerLiteral() *ast.IntegerLiteral {
	value, _ := strconv.Atoi(p.currToken.Literal)
	return &ast.IntegerLiteral{
		Token: p.currToken,
		Value: int64(value),
	}
}
