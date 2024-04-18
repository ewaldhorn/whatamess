package main

import (
	"image/color"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"golang.org/x/image/font"
)

var (
	fontSize        int = 32
	mPlusNormalFont font.Face
	bkg             = color.White
	lightGrey       = color.RGBA{194, 197, 198, 255}
	grey            = color.RGBA{119, 124, 126, 255}
	yellow          = color.RGBA{205, 179, 93, 255}
	green           = color.RGBA{96, 166, 101, 255}
	fontColour      = color.Black
	edge            = false
	grid            [rows * columns]string
	dictionary      []string
	check           [rows * columns]int
	location        = 0
	winner          = false
	answer          string
)

type Game struct{}

func (g *Game) Update() error {
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	ebitenutil.DebugPrint(screen, "Hello, World!")
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 320, 240
}

func main() {
	ebiten.SetWindowSize(640, 480)
	ebiten.SetWindowTitle("Hello, World!")
	if err := ebiten.RunGame(&Game{}); err != nil {
		log.Fatal(err)
	}
}
