package main

import (
	"bytes"
	"image/color"
	"log"
	"math/rand"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/examples/resources/fonts"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
)

var (
	fontSize         float64 = 48.0
	mPlusNormalFont  *text.GoTextFaceSource
	backgroundColour = color.White
	lightGrey        = color.RGBA{194, 197, 198, 255}
	grey             = color.RGBA{119, 124, 126, 255}
	yellow           = color.RGBA{205, 179, 93, 255}
	green            = color.RGBA{96, 166, 101, 255}
	fontColour       = color.Black
	edge             = false
	grid             [rows * columns]string
	dictionary       []string
	check            [rows * columns]int
	location         = 0
	winner           = false
	isPlaying        = true
	answer           string
)

// -------------------------------------------------------------------------------------------------
func pickRandomWordFromDictionary() string {
	return dictionary[rand.Intn(len(dictionary)-1)]
}

// -------------------------------------------------------------------------------------------------
func setupNewGame() {
	answer = pickRandomWordFromDictionary()
}

// =========================================================================================== MAIN
func main() {
	parseDictionary(loadDictionaryFromDisk())
	setupNewGame()

	// load the font
	fontSource, fontErr := text.NewGoTextFaceSource(bytes.NewReader(fonts.MPlus1pRegular_ttf))
	if fontErr != nil {
		log.Fatal(fontErr)
	}
	mPlusNormalFont = fontSource

	ebiten.SetWindowSize(windowWidth, windowHeight)
	ebiten.SetWindowTitle(windowTitle)
	if err := ebiten.RunGame(&Game{}); err != nil {
		log.Fatal(err)
	}
}
