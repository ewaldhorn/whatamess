package game

import (
	"bytes"
	"fmt"
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
	iteration      int
	speed          int
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
	g.speed = 200
}

func NewGame() {
	game := game{
		statistics: newStatistics(),
	}
	game.makeNewLevel(80, 20)
	game.start()
}
