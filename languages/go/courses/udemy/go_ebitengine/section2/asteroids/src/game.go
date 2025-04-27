package main

import (
	"github.com/hajimehoshi/ebiten/v2"
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
	return nil
}

// ----------------------------------------------------------------------------
func (g *Game) Draw(screen *ebiten.Image) {
	g.player.Draw(screen)
	// w, h := assets.PlayerSprite.Bounds().Size().X, assets.PlayerSprite.Bounds().Size().Y
	// tmp := fmt.Sprintf("We have %d x %d", w, h)
	// ebitenutil.DebugPrint(screen, tmp)
	// options := &ebiten.DrawImageOptions{}
	// screen.DrawImage(assets.PlayerSprite, options)
}

// ----------------------------------------------------------------------------
func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return outsideWidth, outsideHeight
}
