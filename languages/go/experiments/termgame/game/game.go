package game

import (
	"bytes"
	"fmt"
	"math/rand"
	"time"
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
	hasCrashed     bool
	iteration      int
	speed          int
	pathWidth      int
	obstacleX      int
	statistics     *statistics
}

type statistics struct {
	start  time.Time
	frames int
	fps    float32
}

// update checks whether we have enough frames to do a FPS calculation
func (s *statistics) update() {
	s.frames += 1
	if s.frames == 30 {
		s.fps = float32(s.frames) / float32(time.Since(s.start).Seconds())
		s.frames = 0
		s.start = time.Now()
	}
}

func newStatistics() *statistics {
	return &statistics{
		start: time.Now(),
	}
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
			default:
				buf.WriteByte(g.level[h][w])
			}

		}
		buf.WriteString("\n")
	}

	fmt.Println(ansiEscapeSeq)
	fmt.Println(buf.String())
	fmt.Printf("Iteration %d at %.0f FPS", g.iteration, g.statistics.fps)

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
		g.statistics.update()
		time.Sleep(time.Millisecond * time.Duration(g.speed))
		if g.hasCrashed {
			break
		}
	}

	if g.hasCrashed {
		fmt.Printf("\n\nKABOOM!\n\n")
	}
}
func (g *game) stop() {
	g.isPaused = true
}
func (g *game) update() {
	g.iteration += 1

	// create new layer
	newLayer := g.makeNewLayer()

	// get ready to combine the new layer with the old ones
	combineThemAll := make([][]byte, g.height)
	combineThemAll[0] = newLayer

	// 'move' the old layers one down, so it looks like the screen is moving
	for i := 1; i < g.height; i++ {
		combineThemAll[i] = g.level[i-1]
	}

	if combineThemAll[g.height-2][g.playerPosition] != ROAD {
		g.hasCrashed = true
	}
	combineThemAll[g.height-1][g.playerPosition] = CAR

	g.level = combineThemAll
}

func (g *game) makeNewLayer() []byte {
	newLayer := make([]byte, g.width)

	chance := rand.Intn(100)

	if chance > 60 {
		g.pathWidth += 1

		if g.pathWidth > g.width/2 {
			g.pathWidth = g.width / 2
		}
	} else if chance < 50 {
		g.pathWidth -= 1

		if g.pathWidth < 3 {
			g.pathWidth = 3
		}
	}

	shift := 0

	if chance < 20 {
		shift -= 1
	}

	if chance > 80 {
		shift += 1
	}

	pathXStart := g.obstacleX + shift

	if pathXStart < 5 {
		pathXStart = 5
	}

	if pathXStart > g.width/2 {
		pathXStart = g.width / 2
	}

	pathXEnd := pathXStart + g.pathWidth

	g.obstacleX = pathXStart

	for w := 0; w < g.width; w++ {
		if w >= pathXStart && w <= pathXEnd {
			newLayer[w] = ROAD
		} else {
			newLayer[w] = OBSTACLE
		}

	}

	return newLayer
}

func (g *game) makeNewLevel(width, height int) {
	level := make([][]byte, height)

	g.pathWidth = width / 6

	pathXStart := width/2 - g.pathWidth
	pathXEnd := width/2 + g.pathWidth

	for h := 0; h < height; h++ {
		level[h] = make([]byte, width)
		for w := 0; w < width; w++ {
			if w >= pathXStart && w <= pathXEnd {
				level[h][w] = ROAD
			} else {
				level[h][w] = OBSTACLE
			}
		}
	}

	g.level = level
	g.width = width
	g.height = height
	g.hasCrashed = false
	g.playerPosition = width / 2
	g.obstacleX = pathXStart
	g.speed = 250
}

func NewGame() {
	game := game{
		statistics: newStatistics(),
	}
	game.makeNewLevel(80, 20)
	game.start()
}
