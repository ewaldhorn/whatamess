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

// Trim all spaces
// Find last index of " " (space)
// Figure out how far from the NEXT character to the end of the string it is
// Return that value, which is the length of the last word
func moreEfficientLengthLastWord(data string) int {
	tmp := strings.TrimSpace(data)
	return len(tmp[strings.LastIndex(tmp, " ")+1:])
}

func main() {
	str := "This is easy"
	expected := len("easy")
	fmt.Printf("%s - %d (%d)\n", str, lengthLastWord(str), expected)
	fmt.Printf("%s - %d (%d)\n\n", str, moreEfficientLengthLastWord(str), expected)

	str = "Not   much harder"
	expected = len("harder")
	fmt.Printf("%s - %d (%d)\n", str, lengthLastWord(str), expected)
	fmt.Printf("%s - %d (%d)\n\n", str, moreEfficientLengthLastWord(str), expected)

	str = "Now, this probably    a bit  harder now "
	expected = len("now")
	fmt.Printf("%s - %d (%d)\n", str, lengthLastWord(str), expected)
	fmt.Printf("%s - %d (%d)\n\n", str, moreEfficientLengthLastWord(str), expected)

	str = "   Ano ther bi t of a challenge      "
	expected = len("challenge")
	fmt.Printf("%s - %d (%d)\n", str, lengthLastWord(str), expected)
	fmt.Printf("%s - %d (%d)\n\n", str, moreEfficientLengthLastWord(str), expected)
}
