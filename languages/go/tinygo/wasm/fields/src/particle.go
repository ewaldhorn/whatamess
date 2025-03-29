package main

import (
	"math"
	"math/rand"

	"github.com/ewaldhorn/tinycanvas/colour"
)

// ----------------------------------------------------------------------------
const (
	MaxHistory   = 200
	MinHistory   = 50
	BorderOffset = 20
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
	history        [MaxHistory]Point
	currentHistory int
	colourRange    int
	maxLength      int
	timer          int
}

// ----------------------------------------------------------------------------
func (p *Particle) draw() {
	// useful for debugging
	if p.effect.isDebugging {
		canvasOne.ColourFilledRectangle(int(p.x), int(p.y), p.size, p.size, *p.colour)
	}

	canvasOne.SetColour(*p.colour)
	for i := 1; i < p.currentHistory; i++ {
		canvasOne.Line(int(p.history[i-1].x), int(p.history[i-1].y), int(p.history[i].x), int(p.history[i].y))
	}
}

// ----------------------------------------------------------------------------
func (p *Particle) update() {
	p.timer -= 1

	if p.timer >= 1 {
		x := int(math.Floor(p.x / float64(CELL_SIZE)))
		y := int(math.Floor(p.y / float64(CELL_SIZE)))

		x = max(0, min(x, COLS-1))
		y = max(0, min(y, ROWS-1))

		idx := y*COLS + x

		p.angle = p.effect.flowField[idx]

		p.speedX = math.Cos(p.angle)
		p.speedY = math.Sin(p.angle)

		p.x += p.speedX * p.speedMod
		p.y += p.speedY * p.speedMod

		p.addPoint(p.x, p.y)
	} else {
		if p.currentHistory > 1 {
			p.currentHistory -= 1
		} else {
			p.reset()
		}
	}
}

// ----------------------------------------------------------------------------
func (p *Particle) addPoint(x, y float64) {
	if p.currentHistory < p.maxLength-1 {
		p.currentHistory += 1
		p.history[p.currentHistory] = Point{x: x, y: y}
	}

	if p.currentHistory >= p.maxLength-1 {
		// slice off the first entry, have a max length to observe
		// p.history = p.history[1:]
		// TODO: Do something
	}
}

// ----------------------------------------------------------------------------
func initParticle(p *Particle, effect *Effect) {
	p.x = BorderOffset + rand.Float64()*float64(CANVAS_WIDTH-BorderOffset)
	p.y = BorderOffset + rand.Float64()*float64(CANVAS_HEIGHT-BorderOffset)
	p.angle = 0.0
	p.speedX = 1.0
	p.speedY = 1.0
	p.speedMod = (rand.Float64() * 7) + 0.15
	p.maxLength = MinHistory + rand.Intn(MaxHistory-MinHistory-5)
	p.timer = p.maxLength * 2
	p.colour = &colours[effect.colourRange][rand.Intn(colour_count)]
}

// ----------------------------------------------------------------------------
func (p *Particle) reset() {
	initParticle(p, p.effect)
	p.currentHistory = 0
	p.history[0] = Point{x: p.x, y: p.y}
}

// ----------------------------------------------------------------------------
func (p *Particle) colourChange() {
	p.colour = &colours[p.effect.colourRange][rand.Intn(colour_count)]
}

// ----------------------------------------------------------------------------
func NewParticle(effect *Effect, size int) *Particle {
	newParticle := Particle{
		effect:      effect,
		size:        size,
		colourRange: effect.colourRange,
	}

	initParticle(&newParticle, effect)
	newParticle.history[0] = Point{x: newParticle.x, y: newParticle.y}

	return &newParticle
}
