package main

import (
	"fmt"
	"goltext/gol"
)

func main() {
	fmt.Printf("\nGame of Life - Console Version.\n")

	gol.PlayGOL(20, 20)

	fmt.Printf("\nDone.\n")
}
