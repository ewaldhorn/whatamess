package main

import (
	"math"
	"math/rand"

	"github.com/ewaldhorn/tinycanvas/colour"
)

// ----------------------------------------------------------------------------
// Potential future enhancements for this file:
// - Use constants instead of magic numbers for speeds and colors
// - Consider using a buffered channel for particle updates
// - Profile memory usage to manage history slice alloc
// - Add particle lifecycle event hooks
// - Implement concurrency-safe particle updates
// - Add interpolation between history points

// ----------------------------------------------------------------------------
const (
	MaxHistory   = 200
	MinHistory   = 50
	BorderOffset = 20
)

// ----------------------------------------------------------------------------
type Particle struct {
	x, y           float32
	size           int
	speedX, speedY float32
	speedMod       float32
	angle          float32
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
	p.timer--

	if p.timer >= 1 {
		// Calculate grid position
		gridX := int(math.Floor(float64(p.x / float32(CELL_SIZE))))
		gridY := int(math.Floor(float64(p.y / float32(CELL_SIZE))))

		// Clamp to grid bounds
		gridX = max(0, min(gridX, COLS-1))
		gridY = max(0, min(gridY, ROWS-1))

		// Get flow field angle at current position
		flowFieldIdx := gridY*COLS + gridX
		p.angle = p.effect.flowField[flowFieldIdx]

		// Update velocity based on angle
		p.speedX = float32(math.Cos(float64(p.angle)))
		p.speedY = float32(math.Sin(float64(p.angle)))

		// Update position
		p.x += p.speedX * p.speedMod
		p.y += p.speedY * p.speedMod

		p.addPoint(p.x, p.y)
	} else {
		// Fade out or reset particle
		if p.currentHistory > 1 {
			p.currentHistory--
		} else {
			p.reset()
		}
	}
}

// ----------------------------------------------------------------------------
func (p *Particle) addPoint(x, y float32) {
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
	p.x = BorderOffset + rand.Float32()*float32(CANVAS_WIDTH-BorderOffset)
	p.y = BorderOffset + rand.Float32()*float32(CANVAS_HEIGHT-BorderOffset)
	p.angle = 0.0
	p.speedX = 1.0
	p.speedY = 1.0
	p.speedMod = (rand.Float32() * 7) + 0.15
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
