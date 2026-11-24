package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"time"
)

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

func findTorpedoTarget(x, y, width, height int, game_map []string) (int, int) {
	for i := 0; i < 30; i++ {
		tx := x + rand.Intn(9) - 4 // Range: -4 to +4
		ty := y + rand.Intn(9) - 4 // Range: -4 to +4

		dist := abs(tx-x) + abs(ty-y)

		if tx >= 0 && tx < width && ty >= 0 && ty < height {
			if dist >= 2 && dist <= 4 && game_map[ty][tx] == '.' {
				return tx, ty
			}
		}
	}

	// couldn't fire
	return -1, -1
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

	for {
		var x, y, myLife, oppLife, torpedoCooldown, sonarCooldown, silenceCooldown, mineCooldown int
		scanner.Scan()
		fmt.Sscan(scanner.Text(), &x, &y, &myLife, &oppLife, &torpedoCooldown, &sonarCooldown, &silenceCooldown, &mineCooldown)

		var sonarResult string
		scanner.Scan()
		fmt.Sscan(scanner.Text(), &sonarResult)

		scanner.Scan()
		opponentOrders := scanner.Text()
		_ = opponentOrders // I don't know what to do with this yet

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

		if torpedoCooldown == 0 {
			target_x, target_y := findTorpedoTarget(x, y, width, height, game_map)
			if target_x > -1 {
				fmt.Println("|TORPEDO", target_x, target_y)
			} else {
				fmt.Println()
			}
		} else {
			fmt.Println()
		}
	}
}
