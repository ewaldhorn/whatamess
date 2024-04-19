package main

import (
	"fmt"
	"image/color"
	"strings"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

// -------------------------------------------------------------------------------------------------
type Game struct {
	pressedKeys []ebiten.Key
	runes       []rune
}

// -------------------------------------------------------------------------------------------------
func reviewPositions() bool {
	hasWon := true
	tmp := strings.Split(answer, "")

	for i := range 5 {
		if grid[location-i] == tmp[4-i] {
			check[location-i] = 1
		} else if strings.Contains(answer, grid[location-i]) {
			check[location-i] = 2
			hasWon = false
		} else {
			check[location-i] = 3
			hasWon = false
		}
	}

	return hasWon
}

// -------------------------------------------------------------------------------------------------
func (g *Game) Update() error {
	if !isWinner && isPlaying {
		g.pressedKeys = inpututil.AppendJustPressedKeys(g.pressedKeys[:0])

		edge = (location+1)%5 == 0

		if isPlaying && len(g.pressedKeys) > 0 {
			pressedKeyString := g.pressedKeys[0].String()

			if strings.Contains(alphabet, pressedKeyString) && pressedKeyString != "" && location >= 0 && location < (rows*columns) {
				grid[location] = pressedKeyString
				if !edge {
					location += 1
				}
			}

			if edge && pressedKeyString == "Enter" && location < (blockCount-1) && grid[location] != "" {
				println("Letters:", grid[location-4], grid[location-3], grid[location-2], grid[location-1], grid[location])
				isWinner = reviewPositions()
				location += 1
			} else if pressedKeyString == "Enter" && location == (blockCount-1) {
				println("At", location, " all done!")
				isWinner = reviewPositions()
				location += 1
				isPlaying = false
			}
		}
	} else {
		if isWinner {
			println("WINNER!!")
			isPlaying = false
		} else {
			println("NOPE!!")
			isPlaying = false
		}
	}
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

			drawOptions := &ebiten.DrawImageOptions{}
			drawOptions.GeoM.Translate(float64(x*85+10), float64(y*85+10))
			screen.DrawImage(block, drawOptions)

			// show active block
			if isPlaying && x+(y*columns) == location {
				vector.StrokeRect(screen, float32(x*85+10), float32(y*85+10), float32(blockSize), float32(blockSize), 2.0, grey, true)
			}

			if grid[x+(y*columns)] != "" {
				drawText(screen,
					fmt.Sprintf("%s", strings.ToUpper(grid[x+(y*columns)])),
					float64(x*85+32),
					float64(y*85+15),
					fontColour,
					fontSize)
			}
		}
	}

	if !isPlaying {
		if isWinner {
			tries := location / 5
			var tryWord string
			if tries > 1 {
				tryWord = "tries"
			} else {
				tryWord = "try"
			}

			drawText(screen, fmt.Sprintf("Done in %d %s!", tries, tryWord), float64(10), float64(520), color.Black, fontSize)
		} else {
			drawText(screen, "Done. Try again!", float64(10), float64(520), color.Black, fontSize)
		}
	}
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
	println("Duration is", duration)
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
