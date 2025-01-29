package main

import "fmt"

// ----------------------------------------------------------------------------
type Player struct {
	name            string
	health          int
	attack, defense int
	luck            int
}

// ----------------------------------------------------------------------------
func main() {
	player1 := Player{name: "Asmongold", health: 30, attack: 8, defense: 5, luck: 4}
	printPlayerInformation(player1)

	player2 := Player{name: "Lowko", health: 30, attack: 8, defense: 5, luck: 3}
	printPlayerInformation(player2)
}

// ----------------------------------------------------------------------------
func printPlayerInformation(player Player) {
	fmt.Printf("Player %s has %d hps\n", player.name, player.health)
	fmt.Println("The battle is about to begin!")
}
