package lexer

import (
	"github.com/aceagles/EagleLang/token"
)

type Lexer struct {
	input      string
	currentPos int
	readPos    int
	char       byte
}

func New(i string) Lexer {
	l := Lexer{}
	l.input = i
	l.readChar()
	return l
}

func (l *Lexer) readChar() {
	if l.readPos >= len(l.input) {
		l.char = 0
	} else {
		l.char = l.input[l.readPos]
	}
	l.currentPos = l.readPos
	l.readPos++
}

func (l *Lexer) NextToken() token.Token {
	var tok token.Token
	l.removeWhitespace()
	switch l.char {
	case '{':
		tok = NewToken(token.LBRACE, l.char)
	case '}':
		tok = NewToken(token.RBRACE, l.char)
	case '(':
		tok = NewToken(token.LPAREN, l.char)
	case ')':
		tok = NewToken(token.RPAREN, l.char)
	case ',':
		tok = NewToken(token.COMMA, l.char)
	case ';':
		tok = NewToken(token.SEMICOLON, l.char)
	case '+':
		tok = NewToken(token.PLUS, l.char)
	case '=':
		tok = NewToken(token.ASSIGN, l.char)
	case 0:
		tok.Literal = ""
		tok.Type = token.EOF
	default:
		if isLetter(l.char) {
			identifier := l.readIdent()
			tok = lookupIdentifier(identifier)
			return tok
		} else if isNumber(l.char) {
			numeral := l.readNumber()
			tok := token.Token{token.INT, numeral}
			return tok
		} else {
			tok = NewToken(token.ILLEGAL, ' ')
		}
	}

	l.readChar()
	return tok
}

func isWhitespace(s byte) bool {
	return s == ' ' || s == '\t' || s == '\r' || s == '\n'
}

func isNumber(s byte) bool {
	return s >= '0' && s <= '9'
}

func (l *Lexer) removeWhitespace() {
	for isWhitespace(l.char) {
		l.readChar()
	}
}

func NewToken(t token.TokenType, l byte) token.Token {
	return token.Token{Type: t, Literal: string(l)}
}

func isLetter(c byte) bool {
	return c >= 'a' && c <= 'z' || c >= 'A' && c <= 'Z' || c == '_'
}

func (l *Lexer) readIdent() string {
	position := l.currentPos

	for isLetter(l.char) {
		l.readChar()
	}
	return l.input[position:l.currentPos]
}

func lookupIdentifier(s string) token.Token {
	switch s {
	case "fn":
		return token.Token{token.FUNCTION, "fn"}
	case "let":
		return token.Token{token.LET, "let"}
	default:
		return token.Token{token.IDENT, s}
	}
}

func (l *Lexer) readNumber() string {
	position := l.currentPos
	for isNumber(l.char) {
		l.readChar()
	}
	return l.input[position:l.currentPos]
}
