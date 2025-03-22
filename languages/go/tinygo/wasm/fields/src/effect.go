package main

const ParticleSize = 50

// ----------------------------------------------------------------------------
type Effect struct {
	width, height   int
	particlesWanted int
	particles       []Particle
}

// ----------------------------------------------------------------------------
func (e *Effect) init() {
	for _ = range e.particlesWanted {
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
func NewEffect(width, height int) *Effect {
	newEffect := Effect{
		width:           width,
		height:          height,
		particlesWanted: 50,
		particles:       []Particle{},
	}

	newEffect.init()

	return &newEffect
}
