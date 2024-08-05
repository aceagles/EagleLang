package repl

import (
	"bufio"
	"fmt"
	"io"

	"github.com/aceagles/EagleLang/lexer"
	"github.com/aceagles/EagleLang/token"
)

func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)

	for {
		io.WriteString(out, `>>> `)
		if scanned := scanner.Scan(); !scanned {
			return
		}
		line := scanner.Text()
		l := lexer.New(line)
		for tok := l.NextToken(); tok.Type != token.EOF; tok = l.NextToken() {
			io.WriteString(out, fmt.Sprintf("%+v\n", tok))
		}

	}
}
