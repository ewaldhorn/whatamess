package game

import "fmt"

type game struct {
	level         [][]byte
	width, height int
}

func (g *game) render() {
	for h := 0; h < g.height; h++ {
		for w := 0; w < g.width; w++ {
			fmt.Printf("%c", g.level[h][w])
		}
		fmt.Println()
	}

}

func (g *game) update() {}

func (g *game) makeNewLevel(width, height int) {
	level := make([][]byte, height)

	for h := 0; h < height; h++ {
		level[h] = make([]byte, width)
		for w := 0; w < width; w++ {
			level[h][w] = '.'
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
