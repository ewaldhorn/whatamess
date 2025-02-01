package main

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

const (
	screenWidth  = 640
	screenHeight = 480
	gameWidth    = 1920
	gameHeight   = 1080
)

var (
	cameraX float64 = 0
	cameraY float64 = 0
)

// ----------------------------------------------------------------------------
type Game struct{}

// ----------------------------------------------------------------------------
func NewGame() *Game {
	game := &Game{}

	return game
}

// ----------------------------------------------------------------------------
func (g *Game) Update() error {
	if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
		// Handle mouse click
	}

	if ebiten.IsKeyPressed(ebiten.KeyLeft) {
		cameraX -= 5
	}
	if ebiten.IsKeyPressed(ebiten.KeyRight) {
		cameraX += 5
	}
	if ebiten.IsKeyPressed(ebiten.KeyUp) {
		cameraY -= 5
	}
	if ebiten.IsKeyPressed(ebiten.KeyDown) {
		cameraY += 5
	}

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	// Draw game world
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(-cameraX, -cameraY)
	screen.DrawImage(ebiten.NewImage(gameWidth, gameHeight).Fill(color.RGBA{0, 255, 0, 255}), op)

	// Draw a rectangle to demonstrate camera movement
	op = &ebiten.DrawImageOptions{}
	op.GeoM.Translate(100-cameraX, 100-cameraY)
	rect := ebiten.NewImage(50, 50).Fill(color.RGBA{255, 0, 0, 255})
	screen.DrawImage(rect, op)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}

func main() {
	ebiten.SetWindowSize(screenWidth, screenHeight)
	ebiten.SetWindowTitle("Camera Example")
	if err := ebiten.RunGame(NewGame()); err != nil {
		panic(err)
	}
}
