package main

import (
	"fmt"
	"math/rand"
	"time"
)

const roundDelay = 400

// ----------------------------------------------------------------------------
func main() {
	player1 := createPlayer1()
	player2 := createPlayer2()

	for player1.IsAlive() && player2.IsAlive() {
		switch rand.Intn(5) {
		case 0:
			fmt.Printf("%s attacks %s! ", player1.name, player2.name)
			playerAttacks(player1, player2)
		case 1:
			fmt.Printf("%s attacks %s! ", player2.name, player1.name)
			playerAttacks(player2, player1)
		case 2:
			if player1.NeedsHealth() {
				fmt.Printf("%s gains health!\n", player1.name)
				player1.RegenerateHealth()
			} else {
				fmt.Println(".")
			}
		case 3:
			if player2.NeedsHealth() {
				fmt.Printf("%s gains health!\n", player2.name)
				player2.RegenerateHealth()
			} else {
				fmt.Println(".")
			}
		default:
			fmt.Println(".")
		}

		time.Sleep(roundDelay * time.Millisecond)
	}

	fmt.Println()

	if !player1.IsAlive() {
		fmt.Printf("%s is DEAD!!\n", player1.name)
	}

	if !player2.IsAlive() {
		fmt.Printf("%s is DEAD!!\n", player2.name)
	}

}
