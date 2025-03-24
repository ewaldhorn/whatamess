package main

import (
	"fmt"
	"math"
	"math/rand"

	"github.com/ewaldhorn/dommie/dom"
	"github.com/ewaldhorn/tinycanvas/colour"
)

// ----------------------------------------------------------------------------
type Particle struct {
	x, y           float64
	size           int
	speedX, speedY float64
	speedMod       float64
	angle          float64
	effect         *Effect
	colour         *colour.Colour
	history        []Point
	colourRange    int
	maxLength      int
	timer          int
}

// ----------------------------------------------------------------------------
func (p *Particle) draw() {
	// useful for debugging
	// canvasOne.ColourFilledRectangle(int(p.x), int(p.y), p.size, p.size, *p.col)

	canvasOne.SetColour(*p.colour)
	for i := 1; i < len(p.history); i++ {
		canvasOne.Line(int(p.history[i-1].x), int(p.history[i-1].y), int(p.history[i].x), int(p.history[i].y))
	}
}

// ----------------------------------------------------------------------------
func (p *Particle) update() {
	p.timer -= 1

	if p.timer >= 1 {
		x := int(math.Floor(p.x / float64(p.effect.cellSize)))
		y := int(math.Floor(p.y / float64(p.effect.cellSize)))

		if x < 0 {
			x = 0
		}

		if y < 0 {
			y = 0
		}

		if x >= p.effect.cols-1 {
			x = p.effect.cols - 1
		}

		if y >= p.effect.rows-1 {
			y = p.effect.rows - 1
		}

		idx := y*p.effect.cols + x

		if idx > p.effect.rows*p.effect.cols-1 {
			dom.Log(fmt.Sprintf("Asked for %d, can only go to %d (%d,%d) (%d)", idx, p.effect.rows*p.effect.cols, x, y, len(p.effect.flowField)-1))
			idx = 1
		}

		p.angle = p.effect.flowField[idx]

		p.speedX = math.Cos(p.angle)
		p.speedY = math.Sin(p.angle)

		p.x += p.speedX * p.speedMod
		p.y += p.speedY * p.speedMod

		p.addPoint(p.x, p.y)
	} else {
		if len(p.history) > 1 {
			p.history = p.history[1:]
		} else {
			p.reset()
		}
	}
}

// ----------------------------------------------------------------------------
func (p *Particle) addPoint(x, y float64) {
	p.history = append(p.history, Point{x: x, y: y})
	if len(p.history) >= p.maxLength {
		// slice off the first entry, have a max length to observe
		p.history = p.history[1:]
	}
}

// ----------------------------------------------------------------------------
func initParticle(p *Particle, effect *Effect) {
	p.x = 20 + rand.Float64()*float64(effect.width-25)
	p.y = 20 + rand.Float64()*float64(effect.height-25)
	p.angle = 0.0
	p.speedX = 1.0
	p.speedY = 1.0
	p.speedMod = (rand.Float64() * 5) + 0.25
	p.maxLength = 30 + rand.Intn(100)
	p.timer = p.maxLength * 2
	p.colour = &colours[effect.colourRange][rand.Intn(colour_count)]
	p.history = []Point{{x: p.x, y: p.y}}
}

// ----------------------------------------------------------------------------
func (p *Particle) reset() {
	initParticle(p, p.effect)
}

// ----------------------------------------------------------------------------
func NewParticle(effect *Effect, size int) *Particle {
	newParticle := Particle{effect: effect, size: size}
	newParticle.colourRange = effect.colourRange
	initParticle(&newParticle, effect)

	return &newParticle
}
