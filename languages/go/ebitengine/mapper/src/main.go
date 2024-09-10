package main

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"

	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

// ----------------------------------------------------------------------------
const (
	SCREEN_WIDTH  int = 1024
	SCREEN_HEIGHT int = 768
)

// ----------------------------------------------------------------------------
type Game struct{}

// ----------------------------------------------------------------------------
func (g *Game) Update() error {
	return nil
}

// ----------------------------------------------------------------------------
func (g *Game) Draw(screen *ebiten.Image) {
	ebitenutil.DebugPrint(screen, "Hello Ebiten!")
}

// ----------------------------------------------------------------------------
func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return SCREEN_WIDTH, SCREEN_HEIGHT
}

// ----------------------------------------------------------------------------
func main() {
	ebiten.SetWindowSize(1024, 768)
	ebiten.SetWindowTitle("Mapper Experiment")

	err := ebiten.RunGame(&Game{})

	if err != nil {
		log.Fatal(err)
	}
}
