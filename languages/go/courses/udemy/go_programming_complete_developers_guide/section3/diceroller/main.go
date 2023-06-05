package main

import (
	"fmt"
	"math/rand"
)

func rollDie(max int) int {
	return 1 + rand.Intn(max)
}

func rollDice(count int) {
	for i := 1; i <= count; i++ {
		fmt.Printf("Die %3d is %d\n", i, rollDie(6))
	}
}

func main() {
	rollDice(10)
	rollAndExplain(10)
}

func rollAndExplain(count int) {
	for j := 0; j < count; j++ {
		die1 := rollDie(6)
		die2 := rollDie(6)

		fmt.Printf("We have %d and %d : ", die1, die2)

		switch tally := die1 + die2; {
		case tally == 2:
			fmt.Println("Snake eyes!")
		case tally == 7:
			fmt.Println("Lucky 7")
		case tally%2 == 0:
			fmt.Println("Even")
		case tally%2 != 0:
			fmt.Println("Odd")
		default:
			fmt.Println("Whatevs!")
		}
	}
}
