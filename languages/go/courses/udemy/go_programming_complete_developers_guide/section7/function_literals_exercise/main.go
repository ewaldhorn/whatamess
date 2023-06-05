package main

import (
	"fmt"
	"unicode"
)

type LineCallback func(line string)

func lineIterator(lines []string, callback LineCallback) {
	for _, val := range lines {
		callback(val)
	}
}

func main() {
	fmt.Printf("\nFunction Literals Exercise.\n")

	lines := []string{"There are", "68 letters,", "five digits,", "12 spaces,", "and 4 punctuation marks in these lines of text!"}

	letters := 0
	numbers := 0
	punctuation := 0
	spaces := 0

	lineFunction := func(line string) {
		for _, r := range line {
			if unicode.IsLetter(r) {
				letters += 1
			} else if unicode.IsNumber(r) {
				numbers += 1
			} else if unicode.IsPunct(r) {
				punctuation += 1
			} else {
				spaces += 1
			}
		}
	}

	lineIterator(lines, lineFunction)

	fmt.Printf("There are:\n%d Letters\n%d Numbers\n%d Punctuation\n%d Spaces\n", letters, numbers, punctuation, spaces)

	fmt.Println()
}
