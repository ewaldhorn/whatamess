package main

import (
	"math/rand"

	"github.com/ewaldhorn/tinycanvas/colour"
	"github.com/ewaldhorn/tinycanvas/tinycanvas"
)

// ----------------------------------------------------------------------------
type Particle struct {
	x, y, size     int
	speedX, speedY int
	effect         *Effect
	col            *colour.Colour
	history        []tinycanvas.Point
	maxLength      int
}

// ----------------------------------------------------------------------------
func (p *Particle) draw() {
	canvasOne.ColourFilledRectangle(p.x, p.y, p.size, p.size, *p.col)

	canvasOne.SetColour(white)
	for i := 1; i < len(p.history); i++ {
		canvasOne.Line(p.history[i-1].X, p.history[i-1].Y, p.history[i].X, p.history[i].Y)
	}
}

// ----------------------------------------------------------------------------
func (p *Particle) update() {
	p.x += p.speedX + (rand.Intn(15) - 7)
	p.y += p.speedY + (rand.Intn(15) - 7)
	p.addPoint(p.x, p.y)
}

// ----------------------------------------------------------------------------
func (p *Particle) addPoint(x, y int) {
	p.history = append(p.history, tinycanvas.Point{X: x, Y: y})
	if len(p.history) > p.maxLength {
		// slice off the first entry
		p.history = p.history[1:]
	}
}

// ----------------------------------------------------------------------------
func NewParticle(effect *Effect, size int) *Particle {
	newParticle := Particle{effect: effect, size: size}

	newParticle.x = rand.Intn(effect.width)
	newParticle.y = rand.Intn(effect.height)
	newParticle.speedX = notZero(rand.Intn(5)-3, 1)
	newParticle.speedY = notZero(rand.Intn(5)-3, 1)
	newParticle.col = colour.NewColourWhite()
	newParticle.history = []tinycanvas.Point{}
	newParticle.addPoint(newParticle.x, newParticle.y)
	newParticle.maxLength = 30

	return &newParticle
}
