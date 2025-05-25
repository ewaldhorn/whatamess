package main

import (
	"fmt"
	"math/rand"
	"syscall/js"
)

const (
	canvasWidth  = 600
	canvasHeight = 400
	paddleWidth  = 10
	paddleHeight = 80
	ballSize     = 10
	paddleSpeed  = 5
	ballSpeed    = 3
)

// Game state
type Game struct {
	ctx          js.Value
	leftPaddleY  float64
	rightPaddleY float64
	ballX        float64
	ballY        float64
	ballDX       float64
	ballDY       float64
	scoreLeft    int
	scoreRight   int
	keysPressed  map[string]bool
}

func main() {
	// Get canvas and context
	doc := js.Global().Get("document")
	canvas := doc.Call("getElementById", "gameCanvas")
	ctx := canvas.Call("getContext", "2d")

	game := &Game{
		ctx:          ctx,
		leftPaddleY:  float64(canvasHeight/2 - paddleHeight/2),
		rightPaddleY: float64(canvasHeight/2 - paddleHeight/2),
		ballX:        float64(canvasWidth / 2),
		ballY:        float64(canvasHeight / 2),
		ballDX:       ballSpeed * (2*rand.Float64() - 1), // Random initial X direction
		ballDY:       ballSpeed * (2*rand.Float64() - 1), // Random initial Y direction
		scoreLeft:    0,
		scoreRight:   0,
		keysPressed:  make(map[string]bool),
	}

	// Set up event listeners for keyboard input
	doc.Call("addEventListener", "keydown", js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		key := args[0].Get("key").String()
		game.keysPressed[key] = true
		return nil
	}))
	doc.Call("addEventListener", "keyup", js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		key := args[0].Get("key").String()
		delete(game.keysPressed, key)
		return nil
	}))

	// Start the game loop
	var renderFrame js.Func
	renderFrame = js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		game.update()
		game.draw()
		js.Global().Call("requestAnimationFrame", renderFrame)
		return nil
	})
	defer renderFrame.Release() // Release the JS function when main exits

	js.Global().Call("requestAnimationFrame", renderFrame)

	// Keep the Go program running indefinitely (otherwise WASM exits)
	<-make(chan bool)
}

func (g *Game) update() {
	// Paddle movement based on keysPressed
	if g.keysPressed["w"] {
		g.leftPaddleY -= paddleSpeed
	}
	if g.keysPressed["s"] {
		g.leftPaddleY += paddleSpeed
	}
	if g.keysPressed["ArrowUp"] {
		g.rightPaddleY -= paddleSpeed
	}
	if g.keysPressed["ArrowDown"] {
		g.rightPaddleY += paddleSpeed
	}

	// Clamp paddles to canvas bounds
	if g.leftPaddleY < 0 {
		g.leftPaddleY = 0
	}
	if g.leftPaddleY > canvasHeight-paddleHeight {
		g.leftPaddleY = canvasHeight - paddleHeight
	}
	if g.rightPaddleY < 0 {
		g.rightPaddleY = 0
	}
	if g.rightPaddleY > canvasHeight-paddleHeight {
		g.rightPaddleY = canvasHeight - paddleHeight
	}

	// Ball movement
	g.ballX += g.ballDX
	g.ballY += g.ballDY

	// Ball collision with top/bottom walls
	if g.ballY <= 0 || g.ballY >= canvasHeight-ballSize {
		g.ballDY *= -1
	}

	// Ball collision with paddles
	// Left paddle
	if g.ballX <= paddleWidth &&
		g.ballY+ballSize >= g.leftPaddleY &&
		g.ballY <= g.leftPaddleY+paddleHeight {
		g.ballDX *= -1
		g.ballX = paddleWidth // Prevent sticking
	}
	// Right paddle
	if g.ballX >= canvasWidth-paddleWidth-ballSize &&
		g.ballY+ballSize >= g.rightPaddleY &&
		g.ballY <= g.rightPaddleY+paddleHeight {
		g.ballDX *= -1
		g.ballX = canvasWidth - paddleWidth - ballSize // Prevent sticking
	}

	// Ball out of bounds (scoring)
	if g.ballX < 0 { // Right scores
		g.scoreRight++
		g.resetBall()
	} else if g.ballX > canvasWidth-ballSize { // Left scores
		g.scoreLeft++
		g.resetBall()
	}
}

func (g *Game) resetBall() {
	g.ballX = float64(canvasWidth / 2)
	g.ballY = float64(canvasHeight / 2)
	g.ballDX = ballSpeed * (2*rand.Float64() - 1)
	g.ballDY = ballSpeed * (2*rand.Float64() - 1)
	println("Score: %d - %d\n", g.scoreLeft, g.scoreRight) // Log to console
}

func (g *Game) draw() {
	// Clear canvas
	g.ctx.Call("clearRect", 0, 0, canvasWidth, canvasHeight)

	// Draw paddles
	g.ctx.Set("fillStyle", "white")
	g.ctx.Call("fillRect", 0, g.leftPaddleY, paddleWidth, paddleHeight)
	g.ctx.Call("fillRect", canvasWidth-paddleWidth, g.rightPaddleY, paddleWidth, paddleHeight)

	// Draw ball
	g.ctx.Call("fillRect", g.ballX, g.ballY, ballSize, ballSize)

	// Draw scores
	g.ctx.Set("font", "30px Arial")
	g.ctx.Set("textAlign", "center")
	g.ctx.Call("fillText", fmt.Sprintf("%d", g.scoreLeft), canvasWidth/4, 40)
	g.ctx.Call("fillText", fmt.Sprintf("%d", g.scoreRight), canvasWidth*3/4, 40)
}
