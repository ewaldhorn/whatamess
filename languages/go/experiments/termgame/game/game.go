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
	level         [][]byte
	width, height int
}

func (g *game) render() {
	buf := new(bytes.Buffer)
	for h := 0; h < g.height; h++ {
		for w := 0; w < g.width; w++ {
			switch g.level[h][w] {
			case ROAD:
				buf.WriteString(" ")
			case OBSTACLE:
				buf.WriteString("*")
			case CAR:
				buf.WriteString("B")
			}
		}
		buf.WriteString("\n")
	}

	fmt.Println(buf.String())

}

func (g *game) update() {}

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
}

func NewGame() {
	game := game{}
	game.makeNewLevel(80, 20)
	game.render()
}
