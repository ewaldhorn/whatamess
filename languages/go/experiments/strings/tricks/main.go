package main

import (
	"fmt"
	"unicode"
	"unicode/utf8"
)

// ----------------------------------------------------------------------------
// A 15 character length string can be bigger in bytes, some runes take more
// than a single byte to represent.
func runesVsCharacters() {
	str := "🚀🕺(Rocket Man!)"

	fmt.Println("Len in bytes     :", len(str))
	fmt.Println("Len in characters:", utf8.RuneCountInString(str))
}

// ----------------------------------------------------------------------------
// Being able to distinguish between digits and characters is handy.
func charOrDigit() {
	str := "This 1 contains 2 digits."

	for _, r := range str {
		if unicode.IsDigit(r) {
			fmt.Printf("\t\t> %c is a digit.\n", r)
		} else {
			fmt.Printf("I don't care about '%c'.\n", r)
		}
	}
}

// ----------------------------------------------------------------------------
func main() {
	runesVsCharacters()
	charOrDigit()
}
