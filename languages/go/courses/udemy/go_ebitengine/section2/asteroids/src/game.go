package main

import "github.com/hajimehoshi/ebiten/v2"

// ------------------------------------------------------------------------------------------------
const (
	GAME_VERSION = "0.01"
	GAME_NAME    = "PEW PEWSTER"
	GAME_WIDTH   = 800
	GAME_HEIGHT  = 800
)

// ------------------------------------------------------------------------------------------------
type Game struct {
	sceneManager *SceneManager
	input        Input
}

// ------------------------------------------------------------------------------------------------
func (g *Game) Update() error {
	if g.sceneManager == nil {
		g.sceneManager = &SceneManager{}
		g.sceneManager.GoToScene(NewGameScene())
	}

	g.input.Update()
	if err := g.sceneManager.Update(&g.input); err != nil {
		return err
	}
	return nil
}

// ------------------------------------------------------------------------------------------------
func (g *Game) Draw(screen *ebiten.Image) {
	g.sceneManager.Draw(screen)
}

// ------------------------------------------------------------------------------------------------
func (g *Game) Layout(_, _ int) (screenWidth, screenHeight int) {
	return GAME_WIDTH, GAME_HEIGHT
}
