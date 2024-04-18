package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

// -------------------------------------------------------------------------------------------------
type Game struct {
	runes []rune
}

// -------------------------------------------------------------------------------------------------
func (g *Game) Update() error {
	return nil
}

// -------------------------------------------------------------------------------------------------
func (g *Game) Draw(screen *ebiten.Image) {
	ebitenutil.DebugPrint(screen, "And here we are!")
}

// -------------------------------------------------------------------------------------------------
func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return windowWidth, windowHeight
}
