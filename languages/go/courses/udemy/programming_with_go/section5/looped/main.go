package main

import (
	"fmt"
	"math/rand"
)

func main() {
	// for like a while loop
	counter := 10

	fmt.Print("Done in: ")
	for counter > 0 {
		fmt.Print(counter, " ")
		counter -= 1
	}

	fmt.Println()

	counter = rand.Intn(950)
	tries := 0
	for counter != 10 {
		counter = rand.Intn(950)
		tries += 1
	}
	fmt.Println("It took", tries, "tries to get a 10!")

	counter = rand.Intn(950)
	for attempts := 0; counter != 10; attempts++ {
		counter = rand.Intn(950)
	}
	fmt.Println("Done!")
}
