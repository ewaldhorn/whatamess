package main

import (
	"asteroids/src/assets"
	"math"
	"math/rand/v2"

	"github.com/hajimehoshi/ebiten/v2"
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
	}

	return m
}
