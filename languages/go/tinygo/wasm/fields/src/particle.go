package main

import (
	"math/rand"

	"github.com/ewaldhorn/tinycanvas/colour"
)

// ----------------------------------------------------------------------------
type Particle struct {
	x, y   int
	effect *Effect
}

// ----------------------------------------------------------------------------
func (p *Particle) draw() {
	canvasOne.ColourFilledRectangle(p.x, p.y, 10, 10, *colour.NewColourWhite())
}

// ----------------------------------------------------------------------------
func NewParticle(effect *Effect) *Particle {
	newParticle := Particle{effect: effect}

	newParticle.x = rand.Intn(effect.width)
	newParticle.y = rand.Intn(effect.height)

	return &newParticle
}
