package gol

import (
	"fmt"
	"math/rand"
	"time"
)

func PlayGOL(width, height int) {
	fmt.Println(ansiEscapeSeq)

	gameBoard1 := NewGameBoard(width, height)
	gameBoard1.PreSeed()
	gameBoard2 := NewGameBoard(width, height)

	currentTick := 0
	maxTicks := 250

	for {
		gameBoard1.Display()
		Tick(width, height, gameBoard1, gameBoard2)
		gameBoard1, gameBoard2 = gameBoard2, gameBoard1
		time.Sleep(delayPerTick * time.Millisecond)
		fmt.Println(ansiEscapeSeq)
		fmt.Printf("On tick %d of %d.\n", currentTick, maxTicks)
		currentTick += 1

		if currentTick > maxTicks {
			break
		}
	}
}

// NewGameBoard creates a new game board, with the provided width and height, one bool per entry.
func NewGameBoard(width, height int) Board {
	tmp := make(Board, height)

	for row := range tmp {
		tmp[row] = make([]bool, width)
	}

	return tmp
}

// PreSeed will populate the game board with a random number of active cells
// There's about a 20% chance of the cell being activated (1/5)
func (b Board) PreSeed() {
	for _, row := range b {
		for i := range row {
			if rand.Intn(5) == 1 {
				row[i] = true
			}
		}
	}
}

func (b Board) Display() {
	for _, row := range b {
		for _, cell := range row {
			switch {
			case cell:
				fmt.Printf("%s", aliveCell)
			default:
				fmt.Printf("%s", deadCell)
			}
		}
		fmt.Printf("\n")
	}
}

func (b Board) IsAlive(width, height, x, y int) bool {
	y = (height + y) % height
	x = (width + x) % width
	return b[y][x]
}

func (b Board) CountNeighbours(width, height, x, y int) int {
	var neighbours int

	for i := y - 1; i <= y+1; i++ {
		for j := x - 1; j <= x+1; j++ {
			if i == y && j == x {
				continue
			}
			if b.IsAlive(width, height, j, i) {
				neighbours++
			}
		}
	}
	return neighbours
}

func (b Board) Next(width, height, x, y int) bool {
	n := b.CountNeighbours(width, height, x, y)
	alive := b.IsAlive(width, height, x, y)
	if n < 4 && n > 1 && alive {
		return true
	} else if n == 3 && !alive {
		return true
	} else {
		return false
	}
}

func Tick(width, height int, a, b Board) {
	for i := 0; i < height; i++ {
		for j := 0; j < width; j++ {
			b[i][j] = a.Next(width, height, j, i)
		}
	}
}
