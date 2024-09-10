package main

import (
	"fmt"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

// ----------------------------------------------------------------------------
type Game struct {
	count int
}

func updateCounter(g *Game) {
	g.count += 1

	if g.count > 250 {
		g.count = 1
		ebiten.SetTPS(1)
	}

	if g.count == 10 {
		ebiten.SetTPS(10)
	}

	if g.count == 40 {
		ebiten.SetTPS(25)
	}

	if g.count == 140 {
		ebiten.SetTPS(60)
	}

}

// ----------------------------------------------------------------------------
func (g *Game) Update() error {
	updateCounter(g)
	return nil
}

// ----------------------------------------------------------------------------
func (g *Game) Draw(screen *ebiten.Image) {
	str := fmt.Sprintf("Hello Ebiten. We are at %d now.", g.count)
	ebitenutil.DebugPrint(screen, str)
}

// ----------------------------------------------------------------------------
func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return SCREEN_WIDTH, SCREEN_HEIGHT
}
