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
	gameWidth    = 2048
	gameHeight   = 2048
	cameraJump   = 10
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
		cameraX -= cameraJump
		if cameraX < 0 {
			cameraX = 0
		}
	}
	if ebiten.IsKeyPressed(ebiten.KeyRight) {
		cameraX += cameraJump
		if cameraX > float64(gameWidth-screenWidth) {
			cameraX = float64(gameWidth - screenWidth)
		}
	}
	if ebiten.IsKeyPressed(ebiten.KeyUp) {
		cameraY -= cameraJump
		if cameraY < 0 {
			cameraY = 0
		}
	}
	if ebiten.IsKeyPressed(ebiten.KeyDown) {
		cameraY += cameraJump
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

	// mini-map follows
	// Draw mini-map
	miniMapWidth := 200
	miniMapHeight := 150
	miniMapX := screenWidth - miniMapWidth - 10
	miniMapY := screenHeight - miniMapHeight - 10

	// Draw mini-map background
	miniMapBackground := ebiten.NewImage(miniMapWidth, miniMapHeight)
	miniMapBackground.Fill(color.RGBA{0, 0, 0, 128})
	op = &ebiten.DrawImageOptions{}
	op.GeoM.Translate(float64(miniMapX), float64(miniMapY))
	screen.DrawImage(miniMapBackground, op)

	// Draw mini-map game world
	miniMapGameWorld := ebiten.NewImage(gameWidth, gameHeight)
	miniMapGameWorld.Fill(color.RGBA{0, 0, 0, 128})
	op = &ebiten.DrawImageOptions{}
	op.GeoM.Scale(float64(miniMapWidth)/float64(gameWidth), float64(miniMapHeight)/float64(gameHeight))
	op.GeoM.Translate(float64(miniMapX), float64(miniMapY))
	screen.DrawImage(miniMapGameWorld, op)

	// Draw mini-map camera rectangle
	miniMapCameraRect := ebiten.NewImage(int(float64(screenWidth)*float64(miniMapWidth)/float64(gameWidth)), int(float64(screenHeight)*float64(miniMapHeight)/float64(gameHeight)))
	miniMapCameraRect.Fill(color.RGBA{64, 64, 64, 128})
	op = &ebiten.DrawImageOptions{}
	op.GeoM.Translate(float64(miniMapX)+(float64(cameraX)/float64(gameWidth))*float64(miniMapWidth), float64(miniMapY)+(float64(cameraY)/float64(gameHeight))*float64(miniMapHeight))
	screen.DrawImage(miniMapCameraRect, op)
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
