package main

import (
	"fmt"
	"os"
)

var hangmanGraphics = []string{}

// ---------------------------------------------------------------------------- loadHangmanGraphics
// Read entire file into memory as a single string. Handy!
func loadHangmanGraphics() {
	for i := range 10 {
		input, err := os.ReadFile(fmt.Sprintf("assets/hangman_%d.txt", i))

		if err != nil {
			fmt.Println(err)
		}

		hangmanGraphics = append(hangmanGraphics, string(input))
	}
}
