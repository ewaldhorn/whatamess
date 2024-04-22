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
	edge        bool
	location    int
	grid        [rows * columns]string
	answer      string
	check       [rows * columns]int
	isWinner    bool
	isPlaying   bool
}

// ----------------
func NewGame() *Game {
	game := Game{}
	game.resetGame()
	return &game
}

// -------------------------------------------------------------------------------------------------
func (g *Game) reviewPositions() bool {
	hasWon := true
	tmp := strings.Split(g.answer, "")

	for i := range 5 {
		if g.grid[g.location-i] == tmp[4-i] {
			g.check[g.location-i] = 1
		} else if strings.Contains(g.answer, g.grid[g.location-i]) {
			g.check[g.location-i] = 2
			hasWon = false
		} else {
			g.check[g.location-i] = 3
			hasWon = false
		}
	}

	return hasWon
}

// -------------------------------------------------------------------------------------------------
func (g *Game) resetGame() {
	g.grid = [rows * columns]string{}
	g.location = 0
	g.answer = pickRandomWordFromDictionary()
	g.check = [rows * columns]int{}
	g.isWinner = false
	g.isPlaying = true

	println(g.answer)
}

// -------------------------------------------------------------------------------------------------
func (g *Game) Update() error {
	if !g.isWinner && g.isPlaying {
		g.pressedKeys = inpututil.AppendJustPressedKeys(g.pressedKeys[:0])

		g.edge = (g.location+1)%5 == 0

		if g.isPlaying && len(g.pressedKeys) > 0 {
			pressedKeyString := g.pressedKeys[0].String()

			if strings.Contains(alphabet, pressedKeyString) && pressedKeyString != "" && g.location >= 0 && g.location < (rows*columns) {
				g.grid[g.location] = pressedKeyString
				if !g.edge {
					g.location += 1
				}
			}

			if g.edge && pressedKeyString == "Enter" && g.location < (blockCount-1) && g.grid[g.location] != "" {
				g.isWinner = g.reviewPositions()
				g.location += 1
			} else if pressedKeyString == "Enter" && g.location == (blockCount-1) {
				println("At", g.location, " all done!")
				g.isWinner = g.reviewPositions()
				g.location += 1
				g.isPlaying = false
			}
		}
	} else {
		if g.isWinner {
			g.isPlaying = false
		} else {
			g.isPlaying = false
		}
		g.pressedKeys = inpututil.AppendJustPressedKeys(g.pressedKeys[:0])
		if len(g.pressedKeys) > 0 {
			pressedKeyString := g.pressedKeys[0].String()
			if pressedKeyString == "Enter" {
				g.resetGame()
			}
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
			switch g.check[x+(y*columns)] {
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
			if g.isPlaying && x+(y*columns) == g.location {
				vector.StrokeRect(screen, float32(x*85+10), float32(y*85+10), float32(blockSize), float32(blockSize), 2.0, grey, true)
			}

			if g.grid[x+(y*columns)] != "" {
				drawText(screen,
					fmt.Sprintf("%s", strings.ToUpper(g.grid[x+(y*columns)])),
					float64(x*85+32),
					float64(y*85+15),
					fontColour,
					fontSize)
			}
		}
	}

	if !g.isPlaying {
		if g.isWinner {
			tries := g.location / 5
			var tryWord string
			if tries > 1 {
				tryWord = "tries"
			} else {
				tryWord = "try"
			}

			drawText(screen, fmt.Sprintf("Done in %d %s!", tries, tryWord), float64(10), float64(520), color.Black, messageFontSize)
		} else {
			drawText(screen, "Oops. Try again?", float64(10), float64(520), color.Black, messageFontSize)
		}
		drawText(screen, "Press ENTER to play again.", float64(10), float64(550), color.Black, messageFontSize)
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
// -------------------------------------------------------------------------------------------------
// -------------------------------------------------------------------------------------------------
// -------------------------------------------------------------------------------------------------
