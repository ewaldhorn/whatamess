package main

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"
)

// ----------------------------------------------------------------------------
const (
	SCREEN_WIDTH  = 640
	SCREEN_HEIGHT = 480
	ballSpeed     = 3
	paddleSpeed   = 6
)

// ----------------------------------------------------------------------------
type Object struct {
	X, Y, W, H int
}

// ----------------------------------------------------------------------------
type Paddle struct {
	Object
}

// ----------------------------------------------------------------------------
type Ball struct {
	Object
	dxdt int // velocity change per tic
	dydt int
}

// ----------------------------------------------------------------------------
type Game struct {
	padded    Paddle
	ball      Ball
	score     int
	highScore int
}

// ----------------------------------------------------------------------------
func (g Game) Layout(outsideW, outsideH int) (int, int) {
	return SCREEN_WIDTH, SCREEN_HEIGHT
}

// ----------------------------------------------------------------------------
func (g Game) Draw(screen *ebiten.Image) {

}

// ----------------------------------------------------------------------------
func (g *Game) Update() error {
	return nil
}

// ----------------------------------------------------------------------------
func main() {
	ebiten.SetWindowTitle("Pong in Ebiten")
	ebiten.SetWindowSize(SCREEN_WIDTH, SCREEN_HEIGHT)

	g := &Game{}

	err := ebiten.RunGame(g)
	if err != nil {
		log.Fatal(err)
	}
}
