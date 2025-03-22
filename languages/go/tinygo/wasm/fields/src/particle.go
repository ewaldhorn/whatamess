package main

import (
	"math/rand"

	"github.com/ewaldhorn/tinycanvas/colour"
)

// ----------------------------------------------------------------------------
type Particle struct {
	x, y   int
	effect *Effect
	col    *colour.Colour
}

// ----------------------------------------------------------------------------
func (p *Particle) draw() {
	canvasOne.ColourFilledRectangle(p.x, p.y, 10, 10, *p.col)
}

// ----------------------------------------------------------------------------
func NewParticle(effect *Effect) *Particle {
	newParticle := Particle{effect: effect}

	newParticle.x = rand.Intn(effect.width)
	newParticle.y = rand.Intn(effect.height)
	newParticle.col = colour.NewColourWhite()

	return &newParticle
}
