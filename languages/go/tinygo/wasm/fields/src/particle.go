package main

import (
	"math"
	"math/rand"

	"github.com/ewaldhorn/tinycanvas/colour"
)

// ----------------------------------------------------------------------------
type Particle struct {
	x, y           float64
	size           int
	speedX, speedY float64
	angle          float64
	effect         *Effect
	col            *colour.Colour
	history        []Point
	maxLength      int
}

// ----------------------------------------------------------------------------
func (p *Particle) draw() {
	canvasOne.ColourFilledRectangle(int(p.x), int(p.y), p.size, p.size, *p.col)

	canvasOne.SetColour(white)
	for i := 1; i < len(p.history); i++ {
		canvasOne.Line(int(p.history[i-1].x), int(p.history[i-1].y), int(p.history[i].x), int(p.history[i].y))
	}
}

// ----------------------------------------------------------------------------
func (p *Particle) update() {
	x := int(p.x) / p.effect.cellSize
	y := int(p.y) / p.effect.cellSize
	idx := y*p.effect.rows + x
	p.angle = p.effect.flowField[idx]

	p.speedX = math.Cos(p.angle) * 2.0
	p.speedY = math.Sin(p.angle) * 2.0

	p.x += p.speedX
	p.y += p.speedX

	p.addPoint(p.x, p.y)
}

// ----------------------------------------------------------------------------
func (p *Particle) addPoint(x, y float64) {
	p.history = append(p.history, Point{x: x, y: y})
	if len(p.history) > p.maxLength {
		// slice off the first entry, have a max length to observe
		p.history = p.history[1:]
	}
}

// ----------------------------------------------------------------------------
func NewParticle(effect *Effect, size int) *Particle {
	newParticle := Particle{effect: effect, size: size}

	newParticle.x = rand.Float64() * float64(effect.width)
	newParticle.y = rand.Float64() * float64(effect.height)
	newParticle.angle = 0.0
	newParticle.speedX = 0.0
	newParticle.speedY = 0.0
	newParticle.col = colour.NewColourWhite()
	newParticle.history = []Point{}
	newParticle.addPoint(newParticle.x, newParticle.y)
	newParticle.maxLength = 40 + rand.Intn(60)

	return &newParticle
}
