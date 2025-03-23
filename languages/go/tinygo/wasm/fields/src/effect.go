package main

import (
	"math"

	"github.com/ewaldhorn/tinycanvas/colour"
)

const ParticleSize = 10

// ----------------------------------------------------------------------------
type Effect struct {
	width, height   int
	particlesWanted int
	cellSize        int
	rows, cols      int
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

	for range e.particlesWanted {
		e.particles = append(e.particles, *NewParticle(e, ParticleSize))
	}
}

// ----------------------------------------------------------------------------
func (e *Effect) drawGrid() {
	canvasOne.SetColour(*colour.NewColour(128, 0, 0, 255))

	for col := range e.cols {
		canvasOne.Line(e.cellSize*col, 0, e.cellSize*col, canvasOne.Height())
	}

	for row := range e.rows {
		canvasOne.Line(0, e.cellSize*row, canvasOne.Width(), e.cellSize*row)
	}
}

// ----------------------------------------------------------------------------
func (e *Effect) render() {
	e.drawGrid()

	for idx, p := range e.particles {
		p.draw()
		p.update()
		e.particles[idx] = p
	}
}

// ----------------------------------------------------------------------------
func NewEffect(width, height, cellSize int) *Effect {
	newEffect := Effect{
		width:           width,
		height:          height,
		particlesWanted: 100,
		cellSize:        cellSize,
		curve:           0.5,
		zoom:            0.2,
		rows:            height / cellSize,
		cols:            width / cellSize,
		particles:       []Particle{},
	}

	newEffect.flowField = []float64{}
	newEffect.init()

	return &newEffect
}
