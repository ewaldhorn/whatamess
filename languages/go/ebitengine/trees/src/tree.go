package main

import (
	"math"

	"github.com/hajimehoshi/bitmapfont/v3"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

// ----------------------------------------------------------------------------
const (
	TREE_HEIGHT = 180
)

// ----------------------------------------------------------------------------
var ops = &ebiten.DrawImageOptions{}
var fontFace = text.NewGoXFace(bitmapfont.Face)

// ----------------------------------------------------------------------------
type Tree struct {
	ebitenImage *ebiten.Image
}

// ----------------------------------------------------------------------------
func NewTree() *Tree {
	tmp := Tree{
		ebitenImage: ebiten.NewImage(int(SCREEN_WIDTH), int(SCREEN_HEIGHT)),
	}

	tmp.ebitenImage.Fill(COLOUR_BACKGROUND)
	tmp.branch(SCREEN_WIDTH/2, SCREEN_HEIGHT-50, SCREEN_WIDTH/2, SCREEN_HEIGHT-50-TREE_HEIGHT, TREE_HEIGHT)

	return &tmp
}

// ----------------------------------------------------------------------------
func (s *Tree) Update() error {

	if ebiten.IsKeyPressed(ebiten.KeyR) {
		s.ebitenImage.Fill(COLOUR_BACKGROUND)
		s.branch(SCREEN_WIDTH/2, SCREEN_HEIGHT-50, SCREEN_WIDTH/2, SCREEN_HEIGHT-50-TREE_HEIGHT, TREE_HEIGHT)
	}

	return nil
}

// ----------------------------------------------------------------------------
func rotateBy(x float32, y float32, len float32, degrees float32) (float32, float32) {
	newX := x + (len * float32(math.Cos(float64(degrees*math.Pi/180))))
	newY := y - (len * float32(math.Sin(float64(degrees*math.Pi/180))))

	return newX, newY
}

// ----------------------------------------------------------------------------
func (s *Tree) branch(startX float32, startY float32, endX float32, endY float32, len float32) {
	if len > 10 {
		// need to work on stroke width
		vector.StrokeLine(s.ebitenImage, startX, startY, endX, endY, len/10, COLOUR_BROWN, true)

		newLen := len * randomBetween(0.5, 0.7)
		newAngle := randomBetween(0, 60)

		newX, newY := rotateBy(endX, endY, newLen, newAngle)
		s.branch(endX, endY, newX, newY, newLen)

		newX, newY = rotateBy(endX, endY, newLen, newAngle+100)
		s.branch(endX, endY, newX, newY, newLen)

	} else {
		vector.DrawFilledCircle(s.ebitenImage, startX, startY, len/3, COLOUR_GREEN, true)
	}
}

// ----------------------------------------------------------------------------
func (s *Tree) Draw(screen *ebiten.Image) {

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
func (s *Tree) Layout(outsideWidth, outsideHeight int) (int, int) {
	return int(SCREEN_WIDTH), int(SCREEN_HEIGHT)
}
