package main

import (
	"image/color"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
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

func (p *Paddle) MoveOnKeyPress() {
	if ebiten.IsKeyPressed(ebiten.KeyArrowDown) {
		p.Y += paddleSpeed
	} else if ebiten.IsKeyPressed(ebiten.KeyArrowUp) {
		p.Y -= paddleSpeed
	}
}

// ----------------------------------------------------------------------------
type Ball struct {
	Object
	dxdt int // velocity change per tic
	dydt int
}

func (b *Ball) Move() {
	b.X += b.dxdt
	b.Y += b.dydt
}

// ----------------------------------------------------------------------------
type Game struct {
	paddle    Paddle
	ball      Ball
	score     int
	highScore int
}

func (g Game) Layout(outsideW, outsideH int) (int, int) {
	return SCREEN_WIDTH, SCREEN_HEIGHT
}

func (g Game) Draw(screen *ebiten.Image) {
	vector.DrawFilledRect(screen, float32(g.paddle.X), float32(g.paddle.Y), float32(g.paddle.W), float32(g.paddle.H), color.White, true)

	vector.DrawFilledRect(screen, float32(g.ball.X), float32(g.ball.Y), float32(g.ball.W), float32(g.ball.H), color.White, true)

	// TODO display score and highscore
	// scoreStr := fmt.Sprintf("Score: %d", g.score)
	// highscoreStr := fmt.Sprintf("High Score: %d", g.highScore)
	// ops := &text.DrawOptions{}

}

func (g *Game) Update() error {
	g.paddle.MoveOnKeyPress()
	g.ball.Move()

	g.CollideWithWall()
	g.CollideWithPaddle()

	return nil
}

func (g *Game) Reset() {
	g.ball.X = 0
	g.ball.Y = 0
	if g.score > g.highScore {
		g.highScore = g.score
	}

	g.score = 0
}

func (g *Game) CollideWithWall() {
	if g.ball.X >= SCREEN_WIDTH {
		g.Reset()
	} else if g.ball.X <= 0 {
		g.ball.dxdt = ballSpeed
	} else if g.ball.Y <= 0 || g.ball.Y >= SCREEN_HEIGHT {
		g.ball.dydt *= -1
	}
}

func (g *Game) CollideWithPaddle() {
	if g.ball.X >= g.paddle.X && g.ball.Y >= g.paddle.Y && g.ball.Y <= g.paddle.Y+g.paddle.H {
		g.ball.dxdt *= -1
	}
}

// ----------------------------------------------------------------------------
func main() {
	ebiten.SetWindowTitle("Pong in Ebiten")
	ebiten.SetWindowSize(SCREEN_WIDTH, SCREEN_HEIGHT)

	g := &Game{}
	g.paddle = Paddle{Object: Object{X: 600, Y: 200, W: 15, H: 100}}
	g.ball = Ball{Object: Object{X: 0, Y: 0, W: 15, H: 15}, dxdt: ballSpeed, dydt: ballSpeed}

	err := ebiten.RunGame(g)
	if err != nil {
		log.Fatal(err)
	}
}
