package main

import (
	"fmt"

	"github.com/hajimehoshi/ebiten/v2"
)

func main() {
	fmt.Println("Startup")
	g := &Game{}
	err := ebiten.RunGame(g)
	if err != nil {
		fmt.Println("Error running game:", err)
	}
}
