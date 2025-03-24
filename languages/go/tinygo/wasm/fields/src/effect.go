package main

import (
	"math"
	"math/rand"

	"github.com/ewaldhorn/tinycanvas/colour"
)

const ParticleSize = 10

// ----------------------------------------------------------------------------
type Effect struct {
	width, height   int
	particlesWanted int
	cellSize        int
	rows, cols      int
	colourRange     int
	isDebugging     bool
	curve, zoom     float64
	flowField       []float64
	particles       []Particle
}

// ----------------------------------------------------------------------------
func (e *Effect) init() {
	// configure flow field angles
	for row := range e.rows {
		for col := range e.cols {
			angle := (math.Cos(float64(col)*e.zoom) + math.Sin(float64(row)*e.zoom)) * e.curve
			e.flowField = append(e.flowField, angle)
		}
	}
}

// ----------------------------------------------------------------------------
func (e *Effect) toggleDebugging() {
	e.isDebugging = !e.isDebugging
}

// ----------------------------------------------------------------------------
func (e *Effect) drawGrid() {
	canvasOne.SetColour(*colour.NewColour(64, 64, 64, 255))

	for col := range e.cols {
		canvasOne.Line(e.cellSize*col, 0, e.cellSize*col, canvasOne.Height())
	}

	for row := range e.rows {
		canvasOne.Line(0, e.cellSize*row, canvasOne.Width(), e.cellSize*row)
	}
}

// ----------------------------------------------------------------------------
func (e *Effect) render() {
	if e.isDebugging {
		e.drawGrid()
	}

	for idx, p := range e.particles {
		p.draw()
		p.update()
		e.particles[idx] = p
	}
}

// ----------------------------------------------------------------------------
func (e *Effect) createParticles() {
	e.particles = nil
	e.particles = []Particle{}

	for range e.particlesWanted {
		e.particles = append(e.particles, *NewParticle(e, ParticleSize))
	}
}

// ----------------------------------------------------------------------------
func (e *Effect) randomise() {
	e.flowField = nil
	e.flowField = []float64{}
	e.curve = 2.5 + rand.Float64()*20
	e.zoom = 0.2 + rand.Float64()*5
	e.init()
	e.colourRange = rand.Intn(10)
	e.createParticles()
}

// ----------------------------------------------------------------------------
func NewEffect(width, height, cellSize int) *Effect {
	newEffect := Effect{
		width:           width,
		height:          height,
		particlesWanted: 1000,
		colourRange:     rand.Intn(2),
		cellSize:        cellSize,
		rows:            height / cellSize,
		cols:            width / cellSize,
		particles:       []Particle{},
	}

	newEffect.flowField = []float64{}
	newEffect.particles = []Particle{}
	newEffect.randomise()
	newEffect.init()

	return &newEffect
}
