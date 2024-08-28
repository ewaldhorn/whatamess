package main

import "github.com/hajimehoshi/ebiten/v2"

type Player struct {
	position Vector2D
	sprite   *ebiten.Image
}

func NewPlayer() *Player {
	return &Player{
		position: Vector2D{X: 100, Y: 100},
		sprite:   PlayerSprite,
	}
}

func (p *Player) Update() {}

func (p *Player) Draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(p.position.X, p.position.Y)
	screen.DrawImage(p.sprite, op)
}
