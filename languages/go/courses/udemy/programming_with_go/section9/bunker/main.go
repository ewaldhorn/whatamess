package main

import (
	"os"

	"github.com/gdamore/tcell/v2"
)

// ----------------------------------------------------------------------------
func displayHelloWorld(screen tcell.Screen) {
	w, h := screen.Size()
	screen.Clear()
	style := tcell.StyleDefault.Foreground(tcell.ColorCadetBlue).Background(tcell.ColorOrangeRed)
	displayText(screen, w/2-7, h/2, style, "Hello, World!")
	displayText(screen, w/2-9, h/2+1, tcell.StyleDefault, "Press ESC to exit.")

	screen.Show()
}

// ----------------------------------------------------------------------------
func main() {
	screen := initScreen()

	displayHelloWorld(screen)
	display(screen, 5, 5, 10, 5, '*')
	screen.Show()

	for {
		switch ev := screen.PollEvent().(type) {
		case *tcell.EventResize:
			screen.Sync()
			displayHelloWorld(screen)
			display(screen, 5, 5, 10, 5, '*')
			screen.Show()
		case *tcell.EventKey:
			if ev.Key() == tcell.KeyEscape {
				screen.Fini()
				os.Exit(0)
			}
		}
	}
}
