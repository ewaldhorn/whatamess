package main

import (
	"fmt"
	"os"
	"strings"
)

func toLeet(s string) string {
	replacements := map[string]string{
		"a": "4",
		"A": "4",
		"e": "3",
		"E": "3",
		"i": "1",
		"I": "1",
		"o": "0",
		"O": "0",
		"s": "5",
		"S": "5",
		"t": "7",
		"T": "7",
	}
	for k, v := range replacements {
		s = strings.ReplaceAll(s, k, v)
	}
	return s
}

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage:", os.Args[0], "string")
		return
	}
	fmt.Println(toLeet(os.Args[1]))
}
