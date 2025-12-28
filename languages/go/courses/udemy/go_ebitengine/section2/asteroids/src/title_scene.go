package main

import (
	"asteroids/src/assets"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
)

// ------------------------------------------------------------------------------------------------
type TitleScene struct{}

// ------------------------------------------------------------------------------------------------
func (t *TitleScene) Draw(screen *ebiten.Image) {
	textToDraw := "1 coin 1 play"
	op := &text.DrawOptions{}
	op.GeoM.Translate(GAME_WIDTH/2, GAME_HEIGHT/2)
	op.LineSpacing = assets.TitleFont.Size * 1.5
	text.Draw(screen, textToDraw, assets.TitleFont, op)
}

// ------------------------------------------------------------------------------------------------
func (t *TitleScene) Update(state *State) error {
	if inpututil.IsKeyJustPressed(ebiten.KeySpace) {
		state.SceneManager.GoToScene(NewGameScene())
		return nil
	}

	return nil
}

// ------------------------------------------------------------------------------------------------
func widthOfText(f *text.GoTextFace, t string) float64 {
	_, w := text.Measure(t, f, f.Size*1.5)
	return w
}
