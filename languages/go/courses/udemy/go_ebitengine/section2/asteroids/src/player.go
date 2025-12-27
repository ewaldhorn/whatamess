package main

import (
	"asteroids/src/assets"
	"math"

	"github.com/hajimehoshi/ebiten/v2"
)

const (
	ROTATION_PER_SECOND = math.Pi
	MAX_ACCELERATION    = 8.0
)

var currentAcceleration float64

// ------------------------------------------------------------------------------------------------
type Player struct {
	gameScene      *GameScene
	sprite         *ebiten.Image
	rotation_angle float64
	position       Vector
	velocity       float64
}

// ------------------------------------------------------------------------------------------------
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

// ------------------------------------------------------------------------------------------------
func (p *Player) Update() {
	speed := ROTATION_PER_SECOND / float64(ebiten.TPS())

	if ebiten.IsKeyPressed(ebiten.KeyLeft) {
		p.rotation_angle -= speed
	} else if ebiten.IsKeyPressed(ebiten.KeyRight) {
		p.rotation_angle += speed
	}

	p.keepOnScreen()
	if ebiten.IsKeyPressed(ebiten.KeyUp) {
		p.accelerate()
	}
}

// ------------------------------------------------------------------------------------------------
func (p *Player) accelerate() {
	// figure out acceleration
	if currentAcceleration < MAX_ACCELERATION {
		currentAcceleration = p.velocity + 2
	}

	if currentAcceleration > MAX_ACCELERATION {
		currentAcceleration = MAX_ACCELERATION
	}

	p.velocity = currentAcceleration

	// now move in the right direction
	dx := math.Sin(p.rotation_angle) * currentAcceleration
	dy := math.Cos(p.rotation_angle) * -currentAcceleration

	p.position.X += dx
	p.position.Y += dy
}

// ------------------------------------------------------------------------------------------------
func (p *Player) keepOnScreen() {
	if p.position.X >= GAME_WIDTH {
		p.position.X = 0
	} else if p.position.X < 0 {
		p.position.X = GAME_WIDTH
	}

	if p.position.Y >= GAME_HEIGHT {
		p.position.Y = 0
	} else if p.position.Y < 0 {
		p.position.Y = GAME_HEIGHT
	}
}

// ------------------------------------------------------------------------------------------------
func NewPlayer(gameScene *GameScene) *Player {
	bounds := assets.PlayerSprite.Bounds()
	halfW := float64(bounds.Dx()) / 2
	halfH := float64(bounds.Dy()) / 2

	return &Player{
		gameScene: gameScene,
		sprite:    assets.PlayerSprite,
		position:  Vector{X: GAME_WIDTH/2 - halfW, Y: GAME_HEIGHT/2 - halfH},
	}
}
