package main

import "fmt"

// ----------------------------------------------------------------------------
func main() {
	player1 := "Asmongold"
	printPlayerInformation(player1, 30)

	player2 := "Lowko"
	printPlayerInformation(player2, 30)
}

// ----------------------------------------------------------------------------
func printPlayerInformation(player string, hps int) {
	fmt.Printf("Player %s has %d hps\n", player, hps)
	fmt.Println("The battle is about to begin!")
}
