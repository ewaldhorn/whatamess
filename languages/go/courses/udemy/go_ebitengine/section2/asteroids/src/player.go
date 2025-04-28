package main

import (
	"asteroids/src/assets"
	"math"

	"github.com/hajimehoshi/ebiten/v2"
)

const (
	rotationPerSecond = math.Pi
)

// ----------------------------------------------------------------------------
type Player struct {
	sprite         *ebiten.Image
	rotation_angle float64
}

// ----------------------------------------------------------------------------
func (p *Player) Draw(screen *ebiten.Image) {
	bounds := p.sprite.Bounds()
	halfW := float64(bounds.Dx() / 2)
	halfH := float64(bounds.Dy() / 2)

	options := &ebiten.DrawImageOptions{}

	options.GeoM.Translate(-halfW, -halfH)
	options.GeoM.Rotate(p.rotation_angle)
	options.GeoM.Translate(halfW, halfH)

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
}

// ----------------------------------------------------------------------------
func NewPlayer(game *Game) *Player {
	return &Player{
		sprite: assets.PlayerSprite,
	}
}
