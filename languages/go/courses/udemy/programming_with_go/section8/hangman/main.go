package main

var game Game

func main() {
	loadDictionary()
	game = NewGame()
	game.printGameStatus()
}
