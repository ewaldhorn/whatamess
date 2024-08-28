package main

import (
	"math/rand"

	"github.com/hajimehoshi/ebiten/v2"
)

type Meteor struct {
	position Vector2D
	sprite   *ebiten.Image
}

func NewMeteor() *Meteor {
	sprite := MeteorSprites[rand.Intn(len(MeteorSprites))]

	return &Meteor{
		position: Vector2D{},
		sprite:   sprite,
	}
}
