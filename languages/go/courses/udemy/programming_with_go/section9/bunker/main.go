package main

import (
	"os"

	"github.com/gdamore/tcell/v2"
)

var xPos = 4
var yPos = 4

// ----------------------------------------------------------------------------
func displayHelloWorld(screen tcell.Screen) {
	w, h := screen.Size()
	style := tcell.StyleDefault.Foreground(tcell.ColorCadetBlue).Background(tcell.ColorOrangeRed)
	displayText(screen, w/2-7, h/2, style, "Hello, World!")
	displayText(screen, w/2-9, h/2+1, tcell.StyleDefault, "Press ESC to exit.")
}

// ----------------------------------------------------------------------------
func renderGameScreen(screen tcell.Screen, game *Game) {
	screen.Clear()

	displayHelloWorld(screen)
	displayText(screen, 10, 20, tcell.StyleDefault.Foreground(tcell.ColorCadetBlue).Background(tcell.ColorOrangeRed), "Packo")
	display(screen, xPos, yPos, 10, 5, '=')
	game.Draw(screen)
	display(screen, 100, 30, 5, 5, 'o')

	screen.Show()
}

// ----------------------------------------------------------------------------
func main() {
	screen := initScreen()
	w, _ := screen.Size()
	game := NewGame(screen)
	game.ToggleDebugging()

	renderGameScreen(screen, game) // not sure if this is needed, I almost think not?

	for {
		shouldRender := false

		switch ev := screen.PollEvent().(type) {
		case *tcell.EventResize:
			screen.Sync()
			w, _ = screen.Size()
			shouldRender = true
		case *tcell.EventKey:
			switch ev.Key() {
			case tcell.KeyEscape:
				screen.Fini()
				os.Exit(0)
			case tcell.KeyRight:
				if xPos < w-10 {
					xPos += 1
					shouldRender = true
				}
			case tcell.KeyLeft:
				if xPos > 0 {
					xPos -= 1
					shouldRender = true
				}
			case tcell.KeyDown:
				game.MovePlayer(tcell.KeyDown)
				shouldRender = true
			case tcell.KeyUp:
				game.MovePlayer(tcell.KeyUp)
				shouldRender = true
			}
		}

		if shouldRender {
			renderGameScreen(screen, game)
		}
	}
}
