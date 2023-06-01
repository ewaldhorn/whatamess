package game

import (
	"bytes"
	"fmt"
)

const (
	ROAD     = 0
	OBSTACLE = 1
	CAR      = 9
)

type game struct {
	level          [][]byte
	width, height  int
	playerPosition int
	isPaused       bool
	iteration      int
}

func (g *game) render() {
	buf := new(bytes.Buffer)
	for h := 0; h < g.height; h++ {
		for w := 0; w < g.width; w++ {
			switch g.level[h][w] {
			case ROAD:
				buf.WriteString("\u2591")
			case OBSTACLE:
				buf.WriteString("\u2593")
			case CAR:
				buf.WriteString("\u2588")
			}
		}
		buf.WriteString("\n")
	}

	fmt.Println(ansiEscapeSeq)
	fmt.Println(buf.String())
	fmt.Printf("Iteration %d", g.iteration)

}

func (g *game) start() {
	g.isPaused = false
	g.loop()
}
func (g *game) pause() {
	g.isPaused = true
}
func (g *game) loop() {
	for !g.isPaused {
		g.update()
		g.render()
	}
}
func (g *game) stop() {
	g.isPaused = true
}
func (g *game) update() {
	g.iteration += 1
}

func (g *game) makeNewLevel(width, height int) {
	level := make([][]byte, height)

	for h := 0; h < height; h++ {
		level[h] = make([]byte, width)
		for w := 0; w < width; w++ {
			level[h][w] = OBSTACLE
		}
	}

	g.level = level
	g.width = width
	g.height = height
	g.playerPosition = width / 2
}

func NewGame() {
	game := game{}
	game.makeNewLevel(80, 20)
	game.start()
}
