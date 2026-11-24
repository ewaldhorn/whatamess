package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

type EnemyTracker struct {
	possibleCells [15][15]bool
}

// Initialize the tracker with every water cell as a possible starting location
func NewEnemyTracker(game_map []string) EnemyTracker {
	var t EnemyTracker

	for y, row := range game_map {
		for x, char := range row {
			if char == '.' {
				t.possibleCells[y][x] = true
			}
		}
	}

	return t
}

func (t *EnemyTracker) ProcessOrders(orders string, game_map []string) {
	if orders == "NA" {
		return
	}

	actions := strings.Split(orders, "|")
	for _, action := range actions {
		parts := strings.Split(strings.TrimSpace(action), " ")
		command := parts[0]

		switch command {
		case "MOVE":
			t.handleMove(parts[1], game_map)
		case "SURFACE":
			sector, _ := strconv.Atoi(parts[1])
			t.handleSurface(sector)
		case "SILENCE":
			t.handleSilence(game_map)
		}
	}
}

// Shift all possible cells in the direction of movement
func (t *EnemyTracker) handleMove(direction string, game_map []string) {
	var nextPossible [15][15]bool

	// Range over the tracker's current possible cells
	for y, row := range t.possibleCells {
		for x, isPossible := range row {
			// If this cell was a valid position, calculate where it moved
			if isPossible {
				nx, ny := x, y
				switch direction {
				case "N":
					ny--
				case "S":
					ny++
				case "E":
					nx++
				case "W":
					nx--
				}

				// If the new position is within bounds and is valid water, keep it as a possibility
				if nx >= 0 && nx < 15 && ny >= 0 && ny < 15 && game_map[ny][nx] == '.' {
					nextPossible[ny][nx] = true
				}
			}
		}
	}
	t.possibleCells = nextPossible
}

// Filter out any cells that are not in the surfaced sector
func (t *EnemyTracker) handleSurface(sector int) {
	// Calculate the bounds of the 5x5 sector (Sectors 1-9)
	minX := ((sector - 1) % 3) * 5
	minY := ((sector - 1) / 3) * 5
	maxX := minX + 4
	maxY := minY + 4

	for y, row := range t.possibleCells {
		for x, isPossible := range row {
			if isPossible {
				// If a cell is outside the sector, it's no longer a possible location
				if x < minX || x > maxX || y < minY || y > maxY {
					t.possibleCells[y][x] = false
				}
			}
		}
	}
}

// Expand all possible cells by up to 4 spaces in all directions
func (t *EnemyTracker) handleSilence(game_map []string) {
	var nextPossible [15][15]bool

	// Directions: N, S, E, W
	dx := []int{0, 0, 1, -1}
	dy := []int{-1, 1, 0, 0}

	for y, row := range t.possibleCells {
		for x, isPossible := range row {
			if isPossible {
				// The enemy might have moved 0 spaces
				nextPossible[y][x] = true

				// Check moving 1 to 4 spaces in each direction
				for d := 0; d < 4; d++ {
					for step := 1; step <= 4; step++ {
						nx := x + (dx[d] * step)
						ny := y + (dy[d] * step)

						if nx >= 0 && nx < 15 && ny >= 0 && ny < 15 && game_map[ny][nx] == '.' {
							nextPossible[ny][nx] = true
						} else {
							break // Stop looking further in this direction if we hit an island or edge
						}
					}
				}
			}
		}
	}
	t.possibleCells = nextPossible
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func findEmptyStartingCell(game_map []string, width, height int) string {
	for {
		x := rand.Intn(width)
		y := rand.Intn(height)

		if game_map[y][x] == '.' {
			return fmt.Sprintf("%d %d", x, y)
		}
	}
}

func determineNewDirection(x, y, width, height int, game_map []string, visited [][]bool) string {
	type Point struct {
		x, y int
		dir  string
	}

	options := []Point{
		{x, y - 1, "N"},
		{x, y + 1, "S"},
		{x + 1, y, "E"},
		{x - 1, y, "W"},
	}

	rand.Shuffle(len(options), func(i, j int) {
		options[i], options[j] = options[j], options[i]
	})

	for _, opt := range options {
		if opt.x >= 0 && opt.x < width && opt.y >= 0 && opt.y < height {
			if game_map[opt.y][opt.x] == '.' && !visited[opt.y][opt.x] {
				return opt.dir
			}
		}
	}

	return "SURFACE" // trapped, need to clear visited blocks
}

func main() {
	rand.Seed(time.Now().UnixNano())

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Buffer(make([]byte, 1000000), 1000000)

	var width, height, myId int
	scanner.Scan()
	fmt.Sscan(scanner.Text(), &width, &height, &myId)

	game_map := make([]string, height)
	for i := 0; i < height; i++ {
		scanner.Scan()
		game_map[i] = scanner.Text()
	}

	// keep track of where we were
	visited := make([][]bool, height)
	for i := 0; i < height; i++ {
		visited[i] = make([]bool, width)
	}

	fmt.Println(findEmptyStartingCell(game_map, width, height))

	tracker := NewEnemyTracker(game_map)

	for {
		var x, y, myLife, oppLife, torpedoCooldown, sonarCooldown, silenceCooldown, mineCooldown int
		scanner.Scan()
		fmt.Sscan(scanner.Text(), &x, &y, &myLife, &oppLife, &torpedoCooldown, &sonarCooldown, &silenceCooldown, &mineCooldown)

		var sonarResult string
		scanner.Scan()
		fmt.Sscan(scanner.Text(), &sonarResult)

		scanner.Scan()
		opponentOrders := scanner.Text()
		tracker.ProcessOrders(opponentOrders, game_map)

		possibleCount := 0
		bestTargetX, bestTargetY := -1, -1

		for y, row := range tracker.possibleCells {
			for x, isPossible := range row {
				if isPossible {
					possibleCount++
					bestTargetX, bestTargetY = x, y
				}
			}
		}

		visited[y][x] = true // mark where we are, i keep forgetting to

		direction := determineNewDirection(x, y, width, height, game_map, visited)

		if direction != "SURFACE" {
			fmt.Print("MOVE ", direction, " TORPEDO")
		} else {
			fmt.Print("SURFACE")

			// surfaced, reset visited spots
			for i := 0; i < height; i++ {
				for j := 0; j < width; j++ {
					visited[i][j] = false
				}
			}
			visited[y][x] = true // mark our current position, this is the source of the position bug
		}

		if torpedoCooldown == 0 && possibleCount <= 3 { // pretty sure we got them
			fmt.Println("|TORPEDO", bestTargetX, bestTargetY)
		} else {
			fmt.Println()
		}
	}
}
