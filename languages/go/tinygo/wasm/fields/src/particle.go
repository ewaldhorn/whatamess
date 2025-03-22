package main

import (
	"math/rand"

	"github.com/ewaldhorn/tinycanvas/colour"
)

// ----------------------------------------------------------------------------
type Particle struct {
	x, y, size     int
	speedX, speedY int
	effect         *Effect
	col            *colour.Colour
}

// ----------------------------------------------------------------------------
func (p *Particle) draw() {
	canvasOne.ColourFilledRectangle(p.x, p.y, p.size, p.size, *p.col)
}

// ----------------------------------------------------------------------------
func (p *Particle) update() {
	p.x += p.speedX
	p.y += p.speedY
}

// ----------------------------------------------------------------------------
func NewParticle(effect *Effect, size int) *Particle {
	newParticle := Particle{effect: effect, size: size}

	newParticle.x = rand.Intn(effect.width)
	newParticle.y = rand.Intn(effect.height)
	newParticle.speedX = 1
	newParticle.speedY = 1
	newParticle.col = colour.NewColourWhite()

	return &newParticle
}
