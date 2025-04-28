package main

import (
	"asteroids/src/assets"
	"math"

	"github.com/hajimehoshi/ebiten/v2"
)

const (
	rotationPerSecond = math.Pi
	maxAcceleration   = 8.0
)

var currentAcceleration float64

// ----------------------------------------------------------------------------
type Player struct {
	game           *Game
	sprite         *ebiten.Image
	rotation_angle float64
	position       Vector
	velocity       float64
}

// ----------------------------------------------------------------------------
func (p *Player) Draw(screen *ebiten.Image) {
	bounds := p.sprite.Bounds()
	halfW := float64(bounds.Dx() / 2)
	halfH := float64(bounds.Dy() / 2)

	options := &ebiten.DrawImageOptions{}

	// handle rotation
	options.GeoM.Translate(-halfW, -halfH)
	options.GeoM.Rotate(p.rotation_angle)
	options.GeoM.Translate(halfW, halfH)

	// handle movement
	options.GeoM.Translate(p.position.X, p.position.Y)

	screen.DrawImage(p.sprite, options)
}

// ----------------------------------------------------------------------------
func (p *Player) Update() {
	speed := rotationPerSecond / float64(ebiten.TPS())

	if ebiten.IsKeyPressed(ebiten.KeyLeft) {
		p.rotation_angle -= speed
	} else if ebiten.IsKeyPressed(ebiten.KeyRight) {
		p.rotation_angle += speed
	}

	if ebiten.IsKeyPressed(ebiten.KeyUp) {
		p.accelerate()
	}
}

// ----------------------------------------------------------------------------
func (p *Player) accelerate() {
	// figure out acceleration
	if currentAcceleration < maxAcceleration {
		currentAcceleration = p.velocity + 2
	}

	if currentAcceleration > maxAcceleration {
		currentAcceleration = maxAcceleration
	}

	p.velocity = currentAcceleration

	// now move in the right direction
	dx := math.Sin(p.rotation_angle) * currentAcceleration
	dy := math.Cos(p.rotation_angle) * -currentAcceleration

	p.position.X += dx
	p.position.Y += dy
}

// ----------------------------------------------------------------------------
func NewPlayer(game *Game) *Player {
	return &Player{
		game:   game,
		sprite: assets.PlayerSprite,
	}
}
