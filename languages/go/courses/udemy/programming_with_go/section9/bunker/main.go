package main

import (
	"os"

	"github.com/gdamore/tcell/v2"
)

var pos = 4

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
	display(screen, pos, 5, 10, 5, '*')

	screen.Show()
}

// ----------------------------------------------------------------------------
func main() {
	screen := initScreen()
	renderGameScreen(screen) // not sure if this is needed, I almost think not?

	for {
		switch ev := screen.PollEvent().(type) {
		case *tcell.EventResize:
			screen.Sync()
			renderGameScreen(screen)
		case *tcell.EventKey:
			switch ev.Key() {
			case tcell.KeyEscape:
				screen.Fini()
				os.Exit(0)
			case tcell.KeyRight:
				pos += 1
				renderGameScreen(screen)
			case tcell.KeyLeft:
				pos -= 1
				renderGameScreen(screen)
			}
		}
	}
}
