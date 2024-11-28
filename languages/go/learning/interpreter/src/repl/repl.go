package repl

import (
	"bufio"
	"fmt"
	"interpreter/src/lexer"
	"interpreter/src/token"
	"io"
)

// ----------------------------------------------------------------------------
const PROMPT = ">: "

// ----------------------------------------------------------------------------
func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)

	for {
		fmt.Fprintf(out, PROMPT)
		scanned := scanner.Scan()

		if !scanned {
			return
		}

		if scanner.Text() == "\\q" {
			break
		} else if scanner.Text() == "\\h" {
			displayHelp()
		} else {
			l := lexer.New(scanner.Text())

			for tok := l.NextToken(); tok.Type != token.EOF; tok = l.NextToken() {
				fmt.Fprintf(out, "%+v\n", tok)
			}
		}
	}
}

// ----------------------------------------------------------------------------
func displayHelp() {
	fmt.Println("\nCommands:\n\\h - display help.")
	fmt.Println("\\q - quit REPL.")
	fmt.Println()
}
