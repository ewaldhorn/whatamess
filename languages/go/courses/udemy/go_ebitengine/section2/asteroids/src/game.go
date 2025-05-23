package main

import (
	"github.com/hajimehoshi/ebiten/v2"
)

// ----------------------------------------------------------------------------
const (
	GAME_VERSION = "0.01"
	GAME_NAME    = "PEW PEWSTER"
	GAME_WIDTH   = 800
	GAME_HEIGHT  = 800
)

// ----------------------------------------------------------------------------
type Game struct {
	player *Player
}

// ----------------------------------------------------------------------------
func NewGame() *Game {
	g := &Game{}

	g.player = NewPlayer(g)

	return g
}

// ----------------------------------------------------------------------------
func (g *Game) Update() error {
	g.player.Update()

	return nil
}

// ----------------------------------------------------------------------------
func (g *Game) Draw(screen *ebiten.Image) {
	g.player.Draw(screen)
}

// ----------------------------------------------------------------------------
func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return GAME_WIDTH, GAME_HEIGHT
}
