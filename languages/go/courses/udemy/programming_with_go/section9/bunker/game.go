package main

import (
	"github.com/gdamore/tcell/v2"
)

const (
	PaddleHeight = 5
)

// ----------------------------------------------------------------------------
type Game struct {
	playerX, playerY, opponentX, opponentY int
	ballX, ballY                           int
	midPaddle                              int
	minY, maxY                             int
}

// ----------------------------------------------------------------------------
func NewGame(screen tcell.Screen) *Game {
	width, height := screen.Size()
	midY := height / 2
	midPaddle := PaddleHeight / 2
	return &Game{
		playerX:   0,
		playerY:   midY - midPaddle,
		opponentY: midY - midPaddle,
		opponentX: width - 1,
		minY:      0,
		maxY:      height - PaddleHeight,
		midPaddle: midPaddle,
	}
}

// ----------------------------------------------------------------------------
func (g *Game) MovePlayer(key tcell.Key) {
	switch key {
	case tcell.KeyDown:
		if g.playerY < g.maxY {
			g.playerY += 1
		}
	case tcell.KeyUp:
		if g.playerY > g.minY {
			g.playerY -= 1
		}
	}
}

// ----------------------------------------------------------------------------
func (g *Game) Update() {}

// ----------------------------------------------------------------------------
func (g *Game) Draw(screen tcell.Screen) {
	display(screen, g.playerX, g.playerY, 1, PaddleHeight, ' ')
	display(screen, 1, g.playerY+g.midPaddle, 1, 1, '>')
	display(screen, g.opponentX, g.opponentY, 1, PaddleHeight, ' ')
	display(screen, g.opponentX-1, g.opponentY+g.midPaddle, 1, 1, '<')
}
