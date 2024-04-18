package main

import (
	"fmt"
	"image/color"
	"strings"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
)

// -------------------------------------------------------------------------------------------------
type Game struct {
	runes []rune
}

// -------------------------------------------------------------------------------------------------
func (g *Game) Update() error {
	return nil
}

// -------------------------------------------------------------------------------------------------
func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(backgroundColour)

	// render grid blocks
	for x := 0; x < columns; x++ {
		for y := 0; y < rows; y++ {
			block := ebiten.NewImage(blockSize, blockSize)
			switch check[x+(y*columns)] {
			case 1:
				block.Fill(green)
				fontColour = color.White
			case 2:
				block.Fill(yellow)
				fontColour = color.White
			case 3:
				block.Fill(grey)
				fontColour = color.White
			default:
				block.Fill(lightGrey)
				fontColour = color.Black
			}
			// todo: show the active block, aka cursor block

			drawOptions := &ebiten.DrawImageOptions{}
			drawOptions.GeoM.Translate(float64(x*85+10), float64(y*85+10))
			screen.DrawImage(block, drawOptions)

			if grid[x+(y*columns)] != "" {
				drawText(screen, fmt.Sprintf("%s", strings.ToUpper(grid[x+(y*columns)])), float64(x*85+32), float64(y*85+15), fontColour, fontSize)
			}
		}
	}

	ebitenutil.DebugPrint(screen, "And here we are!")
}

// -------------------------------------------------------------------------------------------------
func drawText(screen *ebiten.Image, msg string, x float64, y float64, color color.Color, size float64) {
	textOptions := &text.DrawOptions{}
	textOptions.GeoM.Translate(x, y)
	textOptions.ColorScale.ScaleWithColor(color)
	fontFace := &text.GoTextFace{Source: mPlusNormalFont, Size: size}
	text.Draw(screen, msg, fontFace, textOptions)
}

// -------------------------------------------------------------------------------------------------
func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return windowWidth, windowHeight
}

// -------------------------------------------------------------------------------------------------
func suppressRepeatingKey(key ebiten.Key) bool {
	const (
		delay    = 30
		interval = 3
	)

	duration := inpututil.KeyPressDuration(key)

	if duration == 1 {
		return true
	}
	if duration >= delay && (duration-delay&interval == 0) {
		return true
	}

	return false
}

// -------------------------------------------------------------------------------------------------
// -------------------------------------------------------------------------------------------------
// -------------------------------------------------------------------------------------------------
