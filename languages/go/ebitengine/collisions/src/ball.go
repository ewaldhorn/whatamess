package main

import (
	"image/color"
	"math/rand"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

// ----------------------------------------------------------------------------
type Ball struct {
	id             int
	xPos, yPos     float32
	radius         float32
	xDelta, yDelta float32
	collided       bool
}

// ----------------------------------------------------------------------------
func NewBall(id int) *Ball {

	newXPos := 40 + (rand.Float32() * (SCREEN_WIDTH - 100))
	newYPos := 40 + (rand.Float32() * (SCREEN_HEIGHT - 100))

	return &Ball{
		id:     id,
		xPos:   newXPos,
		yPos:   newYPos,
		xDelta: 0.25 + rand.Float32()*2.0,
		yDelta: 0.25 + rand.Float32()*2.0,
		radius: 20.0,
	}
}

// ----------------------------------------------------------------------------
func (b Ball) draw(screen *ebiten.Image) {
	var colour color.Color = COLOUR_WHITE

	if b.collided {
		colour = COLOUR_RED
	}

	vector.DrawFilledCircle(screen, b.xPos, b.yPos, b.radius, colour, true)
}

// ----------------------------------------------------------------------------
func (b *Ball) update() {
	b.xPos += b.xDelta
	b.yPos += b.yDelta
	b.checkBounds()
}

// ----------------------------------------------------------------------------
func (b *Ball) checkBounds() {
	if b.xPos <= b.radius || b.xPos >= SCREEN_WIDTH-b.radius {
		b.xDelta *= -1
	}

	if b.yPos <= b.radius || b.yPos >= SCREEN_HEIGHT-b.radius {
		b.yDelta *= -1
	}
}
