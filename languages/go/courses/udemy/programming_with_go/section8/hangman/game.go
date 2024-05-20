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
		if ch == ' ' {
			fmt.Print("  ")
		} else if game.guessedLetters[ch] == true {
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
		fmt.Println(hangmanGraphics[game.guesses])
		fmt.Println(game.guesses, "guesses used.")
		fmt.Println()
	}

	fmt.Print("Target word: ")
	game.printRevealedLetters()
	fmt.Println()
}

// ----------------------------------------------------------------------------------- letUserGuess
func (game *Game) letUserGuess() bool {
	fmt.Print("\nTake a guess: ")

	var userGuess string
	_, err := fmt.Scanln(&userGuess)

	userGuess = strings.ToLower(strings.TrimSpace(userGuess)) // remove spaces, make lower case

	if err != nil || len(userGuess) != 1 && userGuess != "quit" && userGuess[0] >= 'a' && userGuess[0] <= 'z' {
		fmt.Println("Please enter one letter from a..z, or 'quit' to, well, quit!")
	} else {
		// ok we have a potential guess
		if userGuess == "quit" {
			return true
		}

		guess := userGuess[0]

		if strings.Contains(game.hiddenWord, string(guess)) {
			game.guessedLetters[rune(guess)] = true
		} else {
			game.guesses += 1
		}
	}

	return false
}

// ----------------------------------------------------------------------------------- checkForAWin
func (game Game) checkForAWin() bool {
	hasWon := true

	for _, ch := range game.hiddenWord {
		if game.guessedLetters[ch] == false {
			hasWon = false
			break
		}
	}

	if hasWon {
		fmt.Println("\nYowsers! You made it in", game.guesses, "guesses!")
	}

	return hasWon
}

// ---------------------------------------------------------------------------------- checkForALoss
func (game Game) checkForALoss() bool {
	if game.guesses >= 9 {
		game.printGameStatus()
		fmt.Println("\nOh no, you didn't make it!")
		return true
	}
	return false
}
