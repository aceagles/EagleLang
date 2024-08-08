package parser

import (
	"fmt"
	"testing"

	"github.com/aceagles/EagleLang/ast"
	"github.com/aceagles/EagleLang/lexer"
	"github.com/aceagles/EagleLang/token"
	"github.com/stretchr/testify/assert"
)

func TestLetStatement(t *testing.T) {
	input := `
	let x = 5;
	let y = 10;
	let foobar = 838383;
	`
	l := lexer.New(input)
	p := New(&l)

	program := p.ParseProgram()

	want := ast.Program{
		Statements: []ast.Statement{
			&ast.LetStatement{
				Token: token.Token{Type: token.LET, Literal: "let"},
				Name: &ast.Identifier{
					Token: token.Token{Type: token.IDENT, Literal: "x"},
					Value: "x",
				},
				Value: &ast.IntegerLiteral{
					Token: token.Token{Type: token.INT, Literal: "5"},
					Value: 5,
				},
			},
			&ast.LetStatement{
				Token: token.Token{Type: token.LET, Literal: "let"},
				Name: &ast.Identifier{
					Token: token.Token{Type: token.IDENT, Literal: "y"},
					Value: "y",
				},
				Value: &ast.IntegerLiteral{
					Token: token.Token{Type: token.INT, Literal: "10"},
					Value: 10,
				},
			},
			&ast.LetStatement{
				Token: token.Token{Type: token.LET, Literal: "let"},
				Name: &ast.Identifier{
					Token: token.Token{Type: token.IDENT, Literal: "foobar"},
					Value: "foobar",
				},
				Value: &ast.IntegerLiteral{
					Token: token.Token{Type: token.INT, Literal: "838383"},
					Value: 838383,
				},
			},
		},
	}
	assert.Equal(t, want, *program)

	input = `let = 5;
	let y = ;`
	l = lexer.New(input)
	p = New(&l)
	program = p.ParseProgram()
	fmt.Println(program.Statements)
	assert.Len(t, program.Statements, 2)
	assert.Len(t, p.errors, 2)
	assert.Equal(t, "expected next token to be IDENT but got =", p.errors[0])
	assert.Equal(t, "expected next token to be INT but got ;", p.errors[1])
}

func TestString(t *testing.T) {
	program := &ast.Program{
		Statements: []ast.Statement{
			&ast.LetStatement{
				Token: token.Token{Type: token.LET, Literal: "let"},
				Name: &ast.Identifier{
					Token: token.Token{Type: token.IDENT, Literal: "myVar"},
					Value: "myVar",
				},
				Value: &ast.Identifier{
					Token: token.Token{Type: token.IDENT, Literal: "anotherVar"},
					Value: "anotherVar",
				},
			},
		},
	}
	assert.Equal(t, "let myVar = anotherVar;", program.String())
}
