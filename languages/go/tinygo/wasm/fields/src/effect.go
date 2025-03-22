package main

// ----------------------------------------------------------------------------
type Effect struct {
	width, height int
	particles     []Particle
}

// ----------------------------------------------------------------------------
func (e *Effect) init() {
	e.particles = append(e.particles, *NewParticle(e))
}

// ----------------------------------------------------------------------------
func (e *Effect) render() {
	for _, p := range e.particles {
		p.draw()
	}
}

// ----------------------------------------------------------------------------
func NewEffect(width, height int) *Effect {
	newEffect := Effect{
		width:     width,
		height:    height,
		particles: []Particle{},
	}

	newEffect.init()

	return &newEffect
}
