package main

import (
	"asteroids/src/assets"
	"math"
	"math/rand/v2"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/solarlune/resolv"
)

// ------------------------------------------------------------------------------------------------
const (
	rotationSpeedMin = -0.02
	rotationSpeedMax = 0.02
	fragments        = 4
)

// ------------------------------------------------------------------------------------------------
type Meteor struct {
	game          *GameScene
	position      Vector
	rotation      float64
	movement      Vector
	angle         float64
	rotationSpeed float64
	sprite        *ebiten.Image
	meteorObject  *resolv.Circle // resolv object for collision detection
}

// ------------------------------------------------------------------------------------------------
func NewMeteor(baseVelocity float64, g *GameScene, index int) *Meteor {
	target := Vector{X: GAME_WIDTH / 2, Y: GAME_HEIGHT / 2}
	angle := rand.Float64() * 2 * math.Pi
	distanceFromCenter := GAME_WIDTH/2.0 + 500
	pos := Vector{
		X: target.X + math.Cos(angle)*distanceFromCenter,
		Y: target.Y + math.Sin(angle)*distanceFromCenter,
	}
	velocity := baseVelocity + rand.Float64()*1.5
	direction := Vector{
		X: target.X - pos.X,
		Y: target.Y - pos.Y,
	}

	normalisedDirection := direction.Normalize()
	movement := Vector{
		X: normalisedDirection.X * velocity,
		Y: normalisedDirection.Y * velocity,
	}

	sprite := assets.MeteorSprites[rand.IntN(len(assets.MeteorSprites))]

	m := &Meteor{
		game:          g,
		position:      pos,
		angle:         angle,
		movement:      movement,
		rotationSpeed: rotationSpeedMin + rand.Float64()*(rotationSpeedMax-rotationSpeedMin),
		sprite:        sprite,
		meteorObject:  resolv.NewCircle(pos.X, pos.Y, float64(sprite.Bounds().Dx()/2)),
	}

	m.meteorObject.Tags().Set(TagMeteor | TagLarge)
	m.meteorObject.SetData(&ObjectData{index: index})

	return m
}

// ------------------------------------------------------------------------------------------------
func NewMeteorSmall(baseVelocity float64, g *GameScene, index int) *Meteor {
	target := Vector{X: GAME_WIDTH / 2, Y: GAME_HEIGHT / 2}
	angle := rand.Float64() * 2 * math.Pi
	distanceFromCenter := GAME_WIDTH/2.0 + 500
	pos := Vector{
		X: target.X + math.Cos(angle)*distanceFromCenter,
		Y: target.Y + math.Sin(angle)*distanceFromCenter,
	}
	velocity := baseVelocity + rand.Float64()*1.5
	direction := Vector{
		X: target.X - pos.X,
		Y: target.Y - pos.Y,
	}

	normalisedDirection := direction.Normalize()
	movement := Vector{
		X: normalisedDirection.X * velocity,
		Y: normalisedDirection.Y * velocity,
	}

	sprite := assets.MeteorSprites[rand.IntN(len(assets.MeteorSpritesSmall))]

	m := &Meteor{
		game:          g,
		position:      pos,
		angle:         angle,
		movement:      movement,
		rotationSpeed: rotationSpeedMin + rand.Float64()*(rotationSpeedMax-rotationSpeedMin),
		sprite:        sprite,
		meteorObject:  resolv.NewCircle(pos.X, pos.Y, float64(sprite.Bounds().Dx()/2)),
	}

	m.meteorObject.Tags().Set(TagMeteor | TagSmall)
	m.meteorObject.SetData(&ObjectData{index: index})

	return m
}

// ------------------------------------------------------------------------------------------------
func (m *Meteor) Draw(screen *ebiten.Image) {
	bounds := m.sprite.Bounds()
	halfW := float64(bounds.Dx()) / 2
	halfH := float64(bounds.Dy()) / 2

	op := &ebiten.DrawImageOptions{}

	op.GeoM.Translate(-halfW, -halfH)
	op.GeoM.Rotate(m.rotation)
	op.GeoM.Translate(halfW, halfH)

	op.GeoM.Translate(m.position.X, m.position.Y)

	screen.DrawImage(m.sprite, op)
}

// ------------------------------------------------------------------------------------------------
func (m *Meteor) Update() {
	dx, dy := m.movement.X, m.movement.Y

	m.position.X += dx
	m.position.Y += dy
	m.rotation += m.rotationSpeed

	m.keepOnScreen()

	// remember to update collision detection object
	m.meteorObject.SetPosition(m.position.X, m.position.Y)
}

// ------------------------------------------------------------------------------------------------
func (m *Meteor) keepOnScreen() {
	if m.position.X >= GAME_WIDTH_F64 {
		m.position.X = 0
		// m.meteorObject.SetPosition(0, m.position.Y) // Do we really need this?
	}

	if m.position.X < 0 {
		m.position.X = GAME_WIDTH_F64
		// m.meteorObject.SetPosition(GAME_WIDTH_F64, m.position.Y)
	}

	if m.position.Y >= GAME_HEIGHT_F64 {
		m.position.Y = 0
		// m.meteorObject.SetPosition(m.position.X, 0)
	}

	if m.position.Y < 0 {
		m.position.Y = GAME_HEIGHT_F64
		// m.meteorObject.SetPosition(m.position.X, GAME_HEIGHT_F64)
	}
}
