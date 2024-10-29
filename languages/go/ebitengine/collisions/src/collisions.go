package main

import (
	"image/color"

	"github.com/hajimehoshi/bitmapfont/v3"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
)

// ----------------------------------------------------------------------------
var ops = &ebiten.DrawImageOptions{}
var fontFace = text.NewGoXFace(bitmapfont.Face)

// ----------------------------------------------------------------------------
type Collision struct {
	balls       []Ball
	ebitenImage *ebiten.Image
}

// ----------------------------------------------------------------------------
func NewCollision() *Collision {
	newBalls := make([]Ball, 0)

	for i := range 10 {
		newBalls = append(newBalls, *NewBall(i))
	}

	return &Collision{
		balls:       newBalls,
		ebitenImage: ebiten.NewImage(int(SCREEN_WIDTH), int(SCREEN_HEIGHT)),
	}
}

// ----------------------------------------------------------------------------
func (s *Collision) Update() error {
	for i := range s.balls {
		s.balls[i].update()
		s.balls[i].checkCollisions(s.balls)
	}

	return nil
}

// ----------------------------------------------------------------------------
func (s *Collision) Draw(screen *ebiten.Image) {
	s.ebitenImage.Fill(color.Black)

	for _, b := range s.balls {
		b.draw(s.ebitenImage)
	}

	screen.DrawImage(s.ebitenImage, ops)

	if !ebiten.IsFocused() {
		textOp := &text.DrawOptions{}

		str := "...Click me to interact..."
		tw, th := text.Measure(str, fontFace, textOp.LineSpacing)
		textOp.GeoM.Translate(float64(SCREEN_WIDTH)/2-(tw/2), float64(SCREEN_HEIGHT)/2-(th/2))
		text.Draw(screen, str, fontFace, textOp)
	}

}

// ----------------------------------------------------------------------------
func (s *Collision) Layout(outsideWidth, outsideHeight int) (int, int) {
	return int(SCREEN_WIDTH), int(SCREEN_HEIGHT)
}
