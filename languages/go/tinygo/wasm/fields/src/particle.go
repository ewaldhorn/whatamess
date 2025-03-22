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
	speedX, speedY float64
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
	x := p.x / p.effect.cellSize
	y := p.y / p.effect.cellSize
	idx := y*p.effect.cols + x
	p.angle = p.effect.flowField[idx]

	p.speedX = math.Cos(p.angle) * 2.0
	p.speedY = math.Sin(p.angle) * 2.0

	p.x += int(p.speedX)
	p.y += int(p.speedX)

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
	newParticle.speedX = 0.0
	newParticle.speedY = 0.0
	newParticle.col = colour.NewColourWhite()
	newParticle.history = []tinycanvas.Point{}
	newParticle.addPoint(newParticle.x, newParticle.y)
	newParticle.maxLength = 40 + rand.Intn(60)

	return &newParticle
}
