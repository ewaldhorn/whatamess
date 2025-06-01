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
func renderGameScreen(screen tcell.Screen) {
	screen.Clear()

	displayHelloWorld(screen)
	display(screen, xPos, yPos, 10, 5, '*')

	screen.Show()
}

// ----------------------------------------------------------------------------
func main() {
	screen := initScreen()
	w, h := screen.Size()
	renderGameScreen(screen) // not sure if this is needed, I almost think not?

	for {
		switch ev := screen.PollEvent().(type) {
		case *tcell.EventResize:
			screen.Sync()
			w, h = screen.Size()
			renderGameScreen(screen)
		case *tcell.EventKey:
			switch ev.Key() {
			case tcell.KeyEscape:
				screen.Fini()
				os.Exit(0)
			case tcell.KeyRight:
				if xPos < w-10 {
					xPos += 1
				}
				renderGameScreen(screen)
			case tcell.KeyLeft:
				if xPos > 0 {
					xPos -= 1
				}
				renderGameScreen(screen)
			case tcell.KeyDown:
				if yPos < h-5 {
					yPos += 1
				}
				renderGameScreen(screen)
			case tcell.KeyUp:
				if yPos > 0 {
					yPos -= 1
				}
				renderGameScreen(screen)
			}
		}
	}
}
