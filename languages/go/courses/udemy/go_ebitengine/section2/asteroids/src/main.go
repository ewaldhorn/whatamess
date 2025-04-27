package main

import (
	"fmt"

	"github.com/hajimehoshi/ebiten/v2"
)

// ----------------------------------------------------------------------------
func main() {
	fmt.Print("Startup ... ")

	g := NewGame()

	err := ebiten.RunGame(g)
	if err != nil {
		fmt.Println("Error running game:", err)
	}

	fmt.Println("Done")
}
