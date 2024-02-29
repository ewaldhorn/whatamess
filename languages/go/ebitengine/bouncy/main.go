package main

import (
	"image/color"
	"log"
	"math/rand"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
)

const (
	screenWidth  = 800
	screenHeight = 600
	count        = 2000
)

type Entity struct {
	xPos, yPos int32
	xDir, yDir bool
	colour     color.RGBA
}

func (s *Entity) Init() {
	s.xPos = int32(rand.Intn(screenWidth))
	s.yPos = int32(rand.Intn(screenHeight))
	s.xDir = rand.Intn(50) > 25
	s.yDir = rand.Intn(50) > 25

	s.colour = color.RGBA{R: uint8(rand.Intn(255)), G: uint8(rand.Intn(255)), B: uint8(rand.Intn(255)), A: 0xff}
}

func (s *Entity) RefreshColour() {
	s.colour = color.RGBA{R: uint8(rand.Intn(255)), G: uint8(rand.Intn(255)), B: uint8(rand.Intn(255)), A: 0xff}
}

func (s *Entity) Update() {
	if s.xDir {
		s.xPos += 1
	} else {
		s.xPos -= 1
	}
	if s.yDir {
		s.yPos += 1
	} else {
		s.yPos -= 1
	}

	if s.xPos < 1 || screenWidth < s.xPos {
		s.xDir = !s.xDir
	}

	if s.yPos < 1 || screenHeight < s.yPos {
		s.yDir = !s.yDir
	}

}

func (s *Entity) Draw(screen []byte) {
	if s.xPos < 1 || s.xPos >= screenWidth || s.yPos < 1 || s.yPos >= screenHeight {
		s.RefreshColour()
	} else {
		offset := s.xPos + (s.yPos * screenWidth)

		screen[4*offset] = s.colour.R
		screen[4*offset+1] = s.colour.G
		screen[4*offset+2] = s.colour.B
		screen[4*offset+3] = s.colour.A
	}
}

type Game struct {
	stars  [count]Entity
	pixels []byte
}

func NewGame() *Game {
	g := &Game{}
	for i := 0; i < count; i++ {
		g.stars[i].Init()
	}
	return g
}

func (g *Game) Update() error {
	for i := 0; i < count; i++ {
		g.stars[i].Update()
	}
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	if g.pixels == nil {
		g.pixels = make([]byte, screenWidth*screenHeight*4)
	}

	for i := 0; i < count; i++ {
		g.stars[i].Draw(g.pixels)
	}

	screen.WritePixels(g.pixels)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}

func main() {
	rand.Seed(time.Now().UnixNano())
	ebiten.SetWindowSize(screenWidth, screenHeight)
	ebiten.SetWindowTitle("Bouncy (Ebitengine Experiment)")
	if err := ebiten.RunGame(NewGame()); err != nil {
		log.Fatal(err)
	}
}
