package main

import (
	"image/color"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/examples/resources/fonts"
	"golang.org/x/image/font"
	"golang.org/x/image/font/opentype"
)

var (
	fontSize        float64 = 24.0
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

// =========================================================================================== MAIN
func main() {
	loadDictionary()

	tt, fontErr := opentype.Parse(fonts.MPlus1pRegular_ttf)
	if fontErr != nil {
		log.Fatal("unable to parse font", fontErr)
	}

	myFont, fontErr := opentype.NewFace(tt, &opentype.FaceOptions{Size: fontSize, DPI: screenDpi, Hinting: font.HintingFull})
	if fontErr != nil {
		log.Fatal("unable to load font", fontErr)
	}
	_ = myFont

	ebiten.SetWindowSize(windowWidth, windowHeight)
	ebiten.SetWindowTitle(windowTitle)
	if err := ebiten.RunGame(&Game{}); err != nil {
		log.Fatal(err)
	}
}
