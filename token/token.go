package token

const (
	ILLEGAL = "ILLEGAL"
	EOF     = "EOF"

	IDENT = "IDENT"
	INT   = "INT"

	ASSIGN = "="
	PLUS   = "+"

	COMMA     = ","
	SEMICOLON = ";"

	LPAREN = "("
	RPAREN = ")"
	LBRACE = "{"
	RBRACE = "}"

	FUNCTION = "FUNCTION"
	LET      = "LET"
	IF       = "IF"
	RETURN   = "RETURN"

	LT  = "<"
	GT  = ">"
	LTE = "<="
	GTE = ">="
	EQ  = "=="
	NEQ = "!="

	BANG = "!"
)

var keywords = map[string]TokenType{
	"fn":  FUNCTION,
	"let": LET,
}

type TokenType string

type Token struct {
	Type    TokenType
	Literal string
}

func LookupIdentifier(s string) (tok Token) {
	tok.Literal = s
	if tokType, ok := keywords[s]; ok {
		tok.Type = tokType
	} else {
		tok.Type = IDENT
	}
	return
}
