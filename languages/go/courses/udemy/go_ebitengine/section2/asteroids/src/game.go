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
	sceneManager *SceneManager
	input        Input
	player       *Player
}

// ----------------------------------------------------------------------------
func NewGame() *Game {
	g := &Game{}

	g.player = NewPlayer(g)

	return g
}

// ----------------------------------------------------------------------------
func (g *Game) Update() error {
	if g.sceneManager == nil {
		g.sceneManager = &SceneManager{}
		// g.sceneManager.GoToScene(NewGameScence())
	}

	g.input.Update()
	if err := g.sceneManager.Update(&g.input); err != nil {
		return err
	}

	return nil
}

// ----------------------------------------------------------------------------
func (g *Game) Draw(screen *ebiten.Image) {
	g.sceneManager.Draw(screen)
}

// ----------------------------------------------------------------------------
func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return GAME_WIDTH, GAME_HEIGHT
}
