package main

import (
	"math"
	"math/rand"

	"github.com/ewaldhorn/tinycanvas/colour"
)

const ParticleSize = 10
const ParticleCount = 750

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
func (e *Effect) randomise() {
	e.curve = 1.5 + rand.Float64()*24
	e.zoom = 0.2 + rand.Float64()*6
	e.init()
	e.colourRange = rand.Intn(10)
	e.createParticles()
}

// ----------------------------------------------------------------------------
func NewEffect(width, height, cellSize int) *Effect {
	newEffect := Effect{
		colourRange: rand.Intn(10),
		particles:   []Particle{},
	}

	newEffect.particles = make([]Particle, ParticleCount)
	newEffect.init()
	newEffect.randomise()

	return &newEffect
}
