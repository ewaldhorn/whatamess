package main

import (
	"embed"
	"image"
	_ "image/png"
	"math"

	"github.com/hajimehoshi/ebiten/v2"
)

//go:embed assets/*
var assets embed.FS

var PlayerSprite = mustLoadImage("assets/player.png")

type Vector2D struct {
	X float64
	Y float64
}

type Game struct {
	playerPosition Vector2D
}

// -------------------------------------------------------------- mustLoadImage
func mustLoadImage(name string) *ebiten.Image {
	f, err := assets.Open(name)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	img, _, err := image.Decode(f)
	if err != nil {
		panic(err)
	}

	return ebiten.NewImageFromImage(img)
}

// ------------------------------------------------------------------ toRadians
func toRadians(degrees float64) float64 {
	return degrees * math.Pi / 180.0
}

// --------------------------------------------------------------------- Update
func (g *Game) Update() error {

	speed := float64(120 / ebiten.TPS())
	g.playerPosition.X += speed

	return nil
}

// ----------------------------------------------------------------------- Draw
func (g *Game) Draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(g.playerPosition.X, g.playerPosition.Y)
	screen.DrawImage(PlayerSprite, op)
}

// --------------------------------------------------------------------- Layout
func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return outsideWidth, outsideHeight
}

// ======================================================================= main
func main() {
	g := &Game{playerPosition: Vector2D{X: 100, Y: 100}}

	err := ebiten.RunGame(g)
	if err != nil {
		panic(err)
	}
}
