package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Buffer(make([]byte, 1000000), 1000000)

	var N int
	scanner.Scan()
	fmt.Sscan(scanner.Text(), &N)

	for i := 0; i < N; i++ {
		scanner.Scan()
		line := scanner.Text()
		fmt.Fprintln(os.Stderr, line)

		choices := strings.Split(line[1+strings.Index(line, "("):strings.Index(line, ")")], "|")
		fmt.Fprintln(os.Stdout, fmt.Sprintf("%v%v%v", line[:strings.Index(line, "(")], choices[i], line[1+strings.Index(line, ")"):]))
	}
}

// ----------------------------------------------------------------------------
func applyTemplate(inputs ...string) []string {
	result := make([]string, len(inputs))

	for i := range inputs {

	}

	return result
}

// ----------------------------------------------------------------------------
func getChoicesFromString(string) []string {

}
