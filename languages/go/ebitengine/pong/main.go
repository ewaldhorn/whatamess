package main

import (
	"log"

	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil" // For simpler key state checking
)

const (
	screenWidth  = 600
	screenHeight = 400
	paddleWidth  = 10
	paddleHeight = 80
	ballSize     = 10
	paddleSpeed  = 5
	ballSpeed    = 3
)

type Game struct {
	leftPaddleY  float64
	rightPaddleY float64
	ballX        float64
	ballY        float64
	ballDX       float64
	ballDY       float64
	scoreLeft    int
	scoreRight   int
}

func (g *Game) Update() error {
	// Game logic (input, movement, collisions) goes here
	// Ebiten handles the game loop automatically
	if ebiten.IsKeyPressed(ebiten.KeyW) {
		g.leftPaddleY -= paddleSpeed
	}
	if ebiten.IsKeyPressed(ebiten.KeyS) {
		g.leftPaddleY += paddleSpeed
	}
	// ... (similar for right paddle and ball movement/collision)

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	// Drawing goes here
	screen.Fill(color.Black) // Clear screen

	// Draw paddles
	paddleColor := color.White
	ebitenutil.DrawRect(screen, 0, g.leftPaddleY, paddleWidth, paddleHeight, paddleColor)
	ebitenutil.DrawRect(screen, screenWidth-paddleWidth, g.rightPaddleY, paddleWidth, paddleHeight, paddleColor)

	// Draw ball
	ebitenutil.DrawRect(screen, g.ballX, g.ballY, ballSize, ballSize, paddleColor)

	// Draw scores (you'd need to use a text drawing utility, e.g., from ebiten/v2/text)
	// text.Draw(screen, fmt.Sprintf("%d", g.scoreLeft), someFont, screenWidth/4, 40, color.White)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}

func main() {
	game := &Game{
		// Initialize game state (paddle positions, ball position, etc.)
		leftPaddleY:  float64(screenHeight/2 - paddleHeight/2),
		rightPaddleY: float64(screenHeight/2 - paddleHeight/2),
		ballX:        float64(screenWidth / 2),
		ballY:        float64(screenHeight / 2),
		ballDX:       ballSpeed, // Or random initial
		ballDY:       ballSpeed, // Or random initial
	}
	ebiten.SetWindowSize(screenWidth, screenHeight)
	ebiten.SetWindowTitle("Ebiten Pong")
	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
