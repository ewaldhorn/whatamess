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
	return &Ball{
		id:     id,
		xPos:   rand.Float32() * float32(SCREEN_WIDTH),
		yPos:   rand.Float32() * float32(SCREEN_HEIGHT),
		xDelta: rand.Float32() * 2.5,
		yDelta: rand.Float32() * 2.5,
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
