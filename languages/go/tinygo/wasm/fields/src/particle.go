package main

import (
	"math"
	"math/rand"

	"github.com/ewaldhorn/tinycanvas/colour"
	"github.com/ewaldhorn/tinycanvas/tinycanvas"
)

// ----------------------------------------------------------------------------
type Particle struct {
	x, y, size     int
	speedX, speedY int
	angle          float64
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
	p.angle += 0.5
	p.x += int(float64(p.speedX) + math.Sin(float64(p.angle))*10)
	p.y += int(float64(p.speedY) + math.Cos(float64(p.angle))*7)
	p.addPoint(p.x, p.y)
}

// ----------------------------------------------------------------------------
func (p *Particle) addPoint(x, y int) {
	p.history = append(p.history, tinycanvas.Point{X: x, Y: y})
	if len(p.history) > p.maxLength {
		// slice off the first entry, have a max length to observe
		p.history = p.history[1:]
	}
}

// ----------------------------------------------------------------------------
func NewParticle(effect *Effect, size int) *Particle {
	newParticle := Particle{effect: effect, size: size}

	newParticle.x = rand.Intn(effect.width)
	newParticle.y = rand.Intn(effect.height)
	newParticle.angle = 0.0
	newParticle.speedX = notZero(rand.Intn(5)-3, 1)
	newParticle.speedY = notZero(rand.Intn(5)-3, 1)
	newParticle.col = colour.NewColourWhite()
	newParticle.history = []tinycanvas.Point{}
	newParticle.addPoint(newParticle.x, newParticle.y)
	newParticle.maxLength = 40 + rand.Intn(60)

	return &newParticle
}
