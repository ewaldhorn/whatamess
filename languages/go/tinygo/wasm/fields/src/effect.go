package main

const ParticleSize = 10

// ----------------------------------------------------------------------------
type Effect struct {
	width, height   int
	particlesWanted int
	cellSize        int
	rows, cols      int
	particles       []Particle
}

// ----------------------------------------------------------------------------
func (e *Effect) init() {
	for range e.particlesWanted {
		e.particles = append(e.particles, *NewParticle(e, ParticleSize))
	}
}

// ----------------------------------------------------------------------------
func (e *Effect) render() {
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
		particlesWanted: 50,
		cellSize:        cellSize,
		rows:            width / cellSize,
		cols:            height / cellSize,
		particles:       []Particle{},
	}

	newEffect.init()

	return &newEffect
}
