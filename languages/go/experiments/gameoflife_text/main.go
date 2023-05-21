package main

import (
	"fmt"
	"goltext/gol"
)

func main() {
	fmt.Printf("\nGame of Life - Console Version.\n")

	gol.PlayGOL(30, 30)

	fmt.Printf("\nDone.\n")
}
