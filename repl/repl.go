// Package repl is Read-Eval-Print Loop
package repl

import (
	"bufio"
	"fmt"
	"io"

	"monkey_lang/lexer"
	"monkey_lang/token"
)

const PROMPT = ">>>"

func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)
	for {
		_, _ = fmt.Fprintf(out, PROMPT)
		scanned := scanner.Scan()
		if !scanned {
			return
		}

		line := scanner.Text()
		l := lexer.New(line)

		for tok := l.NextToken(); tok.Type != token.EOF; tok = l.NextToken() {
			fmt.Fprintf(out, "%+v\n", tok)
		}
	}
}
