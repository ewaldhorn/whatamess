package main

import (
	"asteroids/src/assets"
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
)

const (
	MAX_METEORS = 10
)

// ------------------------------------------------------------------------------------------------
type TitleScene struct {
	meteors     map[int]*Meteor
	meteorCount int
}

// ------------------------------------------------------------------------------------------------
func (t *TitleScene) Draw(screen *ebiten.Image) {
	// draw meteors first
	for _, m := range t.meteors {
		m.Draw(screen)
	}

	// now the title text
	textToDraw := "press SPACE to play"
	op := &text.DrawOptions{
		LayoutOptions: text.LayoutOptions{
			PrimaryAlign: text.AlignCenter,
		},
	}
	op.ColorScale.ScaleWithColor(color.White)
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

	if len(t.meteors) < MAX_METEORS {
		m := NewMeteor(0.25, &GameScene{}, len(t.meteors)-1)
		t.meteorCount += 1
		t.meteors[t.meteorCount] = m
	}

	for _, m := range t.meteors {
		m.Update()
	}

	return nil
}

// ------------------------------------------------------------------------------------------------
func widthOfText(f *text.GoTextFace, t string) float64 {
	w, _ := text.Measure(t, f, f.Size*1.5)
	return w
}
