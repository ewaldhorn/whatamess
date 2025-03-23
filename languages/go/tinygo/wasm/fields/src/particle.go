package main

import (
	"fmt"
	"math"
	"math/rand"

	"github.com/ewaldhorn/dommie/dom"
	"github.com/ewaldhorn/tinycanvas/colour"
)

var colours = []colour.Colour{
	*colour.NewColour(32, 32, 128, 255),
	*colour.NewColour(32, 32, 150, 255),
	*colour.NewColour(128, 128, 255, 128),
	*colour.NewColour(90, 90, 215, 255),
	*colour.NewColour(90, 90, 215, 128),
	*colour.NewColour(155, 155, 245, 255),
	*colour.NewColour(64, 64, 128, 128),
	*colour.NewColour(64, 64, 175, 255),
}

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
func (p *Particle) reset() {
	p.x = 20 + rand.Float64()*float64(effect.width-25)
	p.y = 20 + rand.Float64()*float64(effect.height-25)
	p.angle = 0.0
	p.speedX = 1.0
	p.speedY = 1.0
	p.history = []Point{{x: p.x, y: p.y}}
	p.timer = p.maxLength * 2
}

// ----------------------------------------------------------------------------
func NewParticle(effect *Effect, size int) *Particle {
	newParticle := Particle{effect: effect, size: size}

	newParticle.x = 20 + rand.Float64()*float64(effect.width-25)
	newParticle.y = 20 + rand.Float64()*float64(effect.height-25)
	newParticle.angle = 0.0
	newParticle.speedMod = (rand.Float64() * 5) + 1.25
	newParticle.speedX = 1.0
	newParticle.speedY = 1.0
	newParticle.colour = &colours[rand.Intn(len(colours)-1)]
	newParticle.history = []Point{}
	newParticle.addPoint(newParticle.x, newParticle.y)
	newParticle.maxLength = 40 + rand.Intn(80)
	newParticle.timer = newParticle.maxLength * 2

	return &newParticle
}
