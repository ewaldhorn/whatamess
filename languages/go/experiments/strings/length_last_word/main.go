package main

import (
	"fmt"
	"strings"
)

// Trim all spaces
// Splint into words
// Find and return length of last word
func lengthLastWord(data string) int {
	words := strings.Split(strings.TrimSpace(data), " ")
	return len(words[len(words)-1])
}

func main() {
	str := "This is easy"
	expected := len("easy")
	fmt.Printf("%s - %d (%d)\n", str, lengthLastWord(str), expected)

	str = "Not   much harder"
	expected = len("harder")
	fmt.Printf("%s - %d (%d)\n", str, lengthLastWord(str), expected)

	str = "Now, this probably    a bit  harder now "
	expected = len("now")
	fmt.Printf("%s - %d (%d)\n", str, lengthLastWord(str), expected)

	str = "   Ano ther bi t of a challenge      "
	expected = len("challenge")
	fmt.Printf("%s - %d (%d)\n", str, lengthLastWord(str), expected)
}
