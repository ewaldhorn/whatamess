package main

import "github.com/hajimehoshi/ebiten/v2"

type Player struct {
	position Vector2D
	sprite   *ebiten.Image
}

func NewPlayer() *Player {
	sprite := PlayerSprite
	bounds := sprite.Bounds()
	halfW := float64(bounds.Dx()) / 2
	halfH := float64(bounds.Dy()) / 2

	pos := Vector2D{
		X: ScreenWidth/2 - halfW,
		Y: ScreenHeight/2 - halfH,
	}

	return &Player{
		position: pos,
		sprite:   sprite,
	}
}

func (p *Player) Update() {}

func (p *Player) Draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(p.position.X, p.position.Y)
	screen.DrawImage(p.sprite, op)
}
