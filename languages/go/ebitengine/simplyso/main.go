package main

import (
	"github.com/hajimehoshi/ebiten/v2"
)

const (
	ScreenWidth  = 800
	ScreenHeight = 600
)

type Game struct {
	player *Player
}

// --------------------------------------------------------------------- Update
func (g *Game) Update() error {

	g.player.Update()

	return nil
}

// ----------------------------------------------------------------------- Draw
func (g *Game) Draw(screen *ebiten.Image) {
	g.player.Draw(screen)
}

// --------------------------------------------------------------------- Layout
func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return ScreenWidth, ScreenHeight
}

// ======================================================================= main
func main() {
	g := &Game{player: NewPlayer()}

	err := ebiten.RunGame(g)
	if err != nil {
		panic(err)
	}
}
