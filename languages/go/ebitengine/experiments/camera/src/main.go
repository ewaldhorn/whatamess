package main

import (
	"image/color"
	"math/rand"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

// ----------------------------------------------------------------------------
const (
	screenWidth  = 640
	screenHeight = 480
	gameWidth    = 1920
	gameHeight   = 1080
)

// ----------------------------------------------------------------------------
var (
	cameraX float64 = 0
	cameraY float64 = 0
)

// ----------------------------------------------------------------------------
type Game struct {
	world  *ebiten.Image
	rect   *ebiten.Image
	x, y   float64
	xd, yd float64
}

// ----------------------------------------------------------------------------
func NewGame() *Game {
	game := &Game{}

	game.rect = ebiten.NewImage(50, 50)
	game.rect.Fill(color.RGBA{255, 0, 0, 255})

	game.world = ebiten.NewImage(gameWidth, gameHeight)
	game.world.Fill(color.RGBA{0, 255, 0, 255})
	vector.DrawFilledCircle(game.world, 800, 600, 100, color.White, true)
	vector.DrawFilledRect(game.world, gameWidth-200, gameHeight-200, 150, 150, color.RGBA{255, 255, 0, 255}, false)

	game.x = float64(100 + rand.Intn(500))
	game.y = float64(100 + rand.Intn(500))
	game.xd = 1.0
	game.yd = 1.0

	return game
}

// ----------------------------------------------------------------------------
func (g *Game) Update() error {
	if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
		// Handle mouse click
	}

	if ebiten.IsKeyPressed(ebiten.KeyLeft) {
		cameraX -= 5
		if cameraX < 0 {
			cameraX = 0
		}
	}
	if ebiten.IsKeyPressed(ebiten.KeyRight) {
		cameraX += 5
		if cameraX > float64(gameWidth-screenWidth) {
			cameraX = float64(gameWidth - screenWidth)
		}
	}
	if ebiten.IsKeyPressed(ebiten.KeyUp) {
		cameraY -= 5
		if cameraY < 0 {
			cameraY = 0
		}
	}
	if ebiten.IsKeyPressed(ebiten.KeyDown) {
		cameraY += 5
		if cameraY > float64(gameHeight-screenHeight) {
			cameraY = float64(gameHeight - screenHeight)
		}
	}

	g.x += g.xd
	g.y += g.yd

	if g.x < 10 || g.x > gameWidth-60 {
		g.xd *= -1
	}

	if g.y < 10 || g.y > gameHeight-60 {
		g.yd *= -1
	}

	return nil
}

// ----------------------------------------------------------------------------
func (g *Game) Draw(screen *ebiten.Image) {
	// Draw game world
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(-cameraX, -cameraY)
	screen.DrawImage(g.world, op)

	// Draw a rectangle to demonstrate camera movement
	op = &ebiten.DrawImageOptions{}
	op.GeoM.Translate(g.x-cameraX, g.y-cameraY)
	screen.DrawImage(g.rect, op)
}

// ----------------------------------------------------------------------------
func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}

// ----------------------------------------------------------------------------
func main() {
	ebiten.SetWindowSize(screenWidth, screenHeight)
	ebiten.SetWindowTitle("Camera Example")
	if err := ebiten.RunGame(NewGame()); err != nil {
		panic(err)
	}
}
