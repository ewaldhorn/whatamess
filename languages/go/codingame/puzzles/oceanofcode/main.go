package main

import (
	"fmt"
	"math/rand"
	"time"
)

// Find ourselves a random empty cell
func find_empty_cell(game_map []string) string {
	x := 0
	y := 0

	fmt.Println("Y", len(game_map))
	fmt.Println("X", len(game_map[0]))

	for {
		random_x := rand.Intn(len(game_map))
		random_y := rand.Intn(len(game_map[0]))

		if game_map[random_y][random_x] == '.' {
			x = random_x
			y = random_y
			break
		}
	}

	return fmt.Sprintf("%d %d", x, y)
}

func main() {
	rand.Seed(time.Now().UnixNano())
	game_map := make([]string, 3)

	game_map[0] = ".xx"
	game_map[1] = "x.x"
	game_map[2] = "xx."

	fmt.Println(game_map)

	fmt.Println("[", game_map[0], "]")

	fmt.Println(find_empty_cell(game_map))

	//     scanner := bufio.NewScanner(os.Stdin)
	//     scanner.Buffer(make([]byte, 1000000), 1000000)

	//     var width, height, myId int
	//     scanner.Scan()
	//     fmt.Sscan(scanner.Text(),&width, &height, &myId)

	//     game_map := make([]string, height)

	//     for i := 0; i < height; i++ {
	//         scanner.Scan()
	//         line := scanner.Text()
	//         game_map = append(game_map, line)
	//     }

	//     // fmt.Fprintln(os.Stderr, "Debug messages...")
	//     fmt.Println(find_empty_cell(game_map))// Write action to stdout
	//     for {
	//         var x, y, myLife, oppLife, torpedoCooldown, sonarCooldown, silenceCooldown, mineCooldown int
	//         scanner.Scan()
	//         fmt.Sscan(scanner.Text(),&x, &y, &myLife, &oppLife, &torpedoCooldown, &sonarCooldown, &silenceCooldown, &mineCooldown)
	//         var sonarResult string
	//         scanner.Scan()
	//         fmt.Sscan(scanner.Text(),&sonarResult)

	//         scanner.Scan()
	//         opponentOrders := scanner.Text()
	//         _ = opponentOrders // to avoid unused error

	//	        // fmt.Fprintln(os.Stderr, "Debug messages...")
	//	        fmt.Println("MSG just waiting")// Write action to stdout
	//	    }
	//	}
}
