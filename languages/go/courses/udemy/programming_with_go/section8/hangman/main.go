package main

var game Game

func main() {
	// prepare assets
	loadDictionary()
	loadHangmanGraphics()

	game = NewGame()
	game.printGameStatus()
	game.letUserGuess()
}
