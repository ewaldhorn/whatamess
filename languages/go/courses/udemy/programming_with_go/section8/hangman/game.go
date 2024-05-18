package main

import (
	"fmt"
	"strings"
)

type Game struct {
	hiddenWord     string
	guessedLetters map[rune]bool
	guesses        int
}

// ---------------------------------------------------------------------------------------- NewGame
func NewGame() Game {
	game := Game{
		hiddenWord:     getRandomWordFromDictionary(),
		guessedLetters: map[rune]bool{},
	}

	return game
}

// --------------------------------------------------------------------------- printRevealedLetters
func (game Game) printRevealedLetters() {
	for _, ch := range game.hiddenWord {
		if game.guessedLetters[ch] == true {
			fmt.Printf("%c ", ch)
		} else {
			fmt.Print("_ ")
		}
	}
}

// -------------------------------------------------------------------------------- printGameStatus
func (game Game) printGameStatus() {
	fmt.Println()
	fmt.Println(strings.Repeat("-_-_-_-_-_-_", 5))
	fmt.Println()
	if game.guesses > 0 {
		fmt.Println(game.guesses, " guesses used.")
	}
	fmt.Print("Target word: ")
	game.printRevealedLetters()
	fmt.Println()
}
