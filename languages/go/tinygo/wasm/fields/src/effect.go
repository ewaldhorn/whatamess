package main

import (
	"math"
	"math/rand"

	"github.com/ewaldhorn/tinycanvas/colour"
)

// ----------------------------------------------------------------------------
const (
	ParticleSize  = 6
	ParticleCount = 750
)

// ----------------------------------------------------------------------------
type Effect struct {
	colourRange int
	isDebugging bool
	curve, zoom float64
	flowField   [ROWS * COLS]float64
	particles   []Particle
}

// ----------------------------------------------------------------------------
func (e *Effect) init() {
	// configure flow field angles
	pos := 0
	for row := range ROWS {
		for col := range COLS {
			angle := (math.Cos(float64(col)*e.zoom) + math.Sin(float64(row)*e.zoom)) * e.curve
			e.flowField[pos] = angle
			pos++
		}
	}
}

// ----------------------------------------------------------------------------
func (e *Effect) toggleDebugging() {
	e.isDebugging = !e.isDebugging
}

// ----------------------------------------------------------------------------
func (e *Effect) drawGrid() {
	canvasOne.SetColour(*colour.NewColour(64, 64, 64, 128))

	for col := range COLS {
		canvasOne.Line(CELL_SIZE*col, 0, CELL_SIZE*col, canvasOne.Height())
	}

	for row := range ROWS {
		canvasOne.Line(0, CELL_SIZE*row, canvasOne.Width(), CELL_SIZE*row)
	}
}

// ----------------------------------------------------------------------------
func (e *Effect) render() {
	if e.isDebugging {
		e.drawGrid()
	}

	for idx := range e.particles {
		e.particles[idx].draw()
		e.particles[idx].update()
	}
}

// ----------------------------------------------------------------------------
func (e *Effect) createParticles() {
	for idx := range ParticleCount {
		e.particles[idx] = *NewParticle(e, ParticleSize)
	}
}

// ----------------------------------------------------------------------------
func (e *Effect) randomise(changeColour bool) {
	e.curve = 1.0 + rand.Float64()*30
	e.zoom = 0.10 + rand.Float64()*8
	e.init()
	if changeColour {
		e.colourRange = rand.Intn(10)
	}
	e.createParticles()
}

// ----------------------------------------------------------------------------
// Only changes the colours, nothing else
func (e *Effect) colourChange() {
	e.colourRange = rand.Intn(10)
	for idx := range ParticleCount {
		e.particles[idx].colourChange()
	}
}

// ----------------------------------------------------------------------------
// Changes everything, except the colours
func (e *Effect) patternChange() {
	e.randomise(false)
}

// ----------------------------------------------------------------------------
func NewEffect(width, height, cellSize int) *Effect {
	newEffect := Effect{
		colourRange: rand.Intn(10),
		particles:   []Particle{},
	}

	newEffect.particles = make([]Particle, ParticleCount)
	newEffect.init()
	newEffect.randomise(true)

	return &newEffect
}
