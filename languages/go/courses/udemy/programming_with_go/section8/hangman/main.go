package main

var game Game

func main() {
	// prepare assets
	loadDictionary()
	loadHangmanGraphics()

	game = NewGame()

	isQuitting := false

	for isQuitting == false {
		game.printGameStatus()
		isQuitting = game.letUserGuess()
		isQuitting = isQuitting || game.checkForAWin()
		isQuitting = isQuitting || game.checkForALoss()
	}
}
