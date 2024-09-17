package main

import (
	"fmt"
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
)

const (
	screenWidth  = 640
	screenHeight = 480
)

type Game struct {
	box1           *ebiten.Image
	box2           *ebiten.Image
	mouseX, mouseY int
	boxX, boxY     int
}

func (g *Game) Update() error {
	// Move box around to simulate collision
	g.mouseX, g.mouseY = ebiten.CursorPosition()

	// Check for collision
	if g.checkCollision() {
		fmt.Println("Collision at", g.mouseX, g.mouseY)
	}

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	mv := &ebiten.DrawImageOptions{}
	mv.GeoM.Translate(float64(g.mouseX), float64(g.mouseY))
	screen.DrawImage(g.box1, mv)

	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(float64(g.boxX), float64(g.boxY))
	screen.DrawImage(g.box2, op)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}

func (g *Game) checkCollision() bool {
	return g.mouseX+50 >= g.boxX && g.mouseX <= g.boxX+50 &&
		g.mouseY+50 >= g.boxY && g.mouseY <= g.boxY+50
}

func main() {
	ebiten.SetWindowSize(screenWidth, screenHeight)
	ebiten.SetWindowTitle("Collision Detection")

	box1 := ebiten.NewImage(50, 50)
	box1.Fill(color.RGBA{R: 255, G: 0, B: 0, A: 255}) // Red

	box2 := ebiten.NewImage(50, 50)
	box2.Fill(color.RGBA{R: 0, G: 255, B: 0, A: 255}) // Green

	game := &Game{box1: box1, box2: box2, boxX: screenWidth/2 - 25, boxY: screenHeight/2 - 25}
	ebiten.RunGame(game)
}
