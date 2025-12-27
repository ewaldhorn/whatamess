package main

import (
	"fmt"

	"github.com/hajimehoshi/ebiten/v2"
)

// ----------------------------------------------------------------------------
func main() {
	fmt.Print("Startup ... ")

	ebiten.SetWindowSize(GAME_WIDTH, GAME_HEIGHT)
	ebiten.SetWindowTitle(fmt.Sprintf("%s v%s", GAME_NAME, GAME_VERSION))

	err := ebiten.RunGame(&Game{})
	if err != nil {
		fmt.Println("Error running game:", err)
	}

	fmt.Println("Done")
}
