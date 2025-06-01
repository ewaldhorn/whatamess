package main

import (
	"fmt"
	"os"

	"github.com/gdamore/tcell/v2"
	"github.com/mattn/go-runewidth"
)

// ----------------------------------------------------------------------------
func emitStr(s tcell.Screen, x, y int, style tcell.Style, str string) {
	for _, c := range str {
		var comb []rune
		w := runewidth.RuneWidth(c)
		if w == 0 {
			comb = []rune{c}
			c = ' '
			w = 1
		}
		s.SetContent(x, y, c, comb, style)
		x += w
	}
}

// ----------------------------------------------------------------------------
func displayHelloWorld(s tcell.Screen) {
	w, h := s.Size()
	s.Clear()
	style := tcell.StyleDefault.Foreground(tcell.ColorCadetBlue).Background(tcell.ColorOrangeRed)
	emitStr(s, w/2-7, h/2, style, "Hello, World!")
	emitStr(s, w/2-9, h/2+1, tcell.StyleDefault, "Press ESC to exit.")
	s.Show()
}

// ----------------------------------------------------------------------------
func getNewScreen() tcell.Screen {
	newScreen, e := tcell.NewScreen()
	if e != nil {
		fmt.Fprintf(os.Stderr, "%v\n", e)
		os.Exit(1)
	}

	if e := newScreen.Init(); e != nil {
		fmt.Fprintf(os.Stderr, "%v\n", e)
		os.Exit(1)
	}

	return newScreen
}

// ----------------------------------------------------------------------------
func setupStyle(screen tcell.Screen) {
	defStyle := tcell.StyleDefault.Background(tcell.ColorBlack).Foreground(tcell.ColorWhite)
	screen.SetStyle(defStyle)
}

// ----------------------------------------------------------------------------
func main() {
	s := getNewScreen()
	setupStyle(s)

	displayHelloWorld(s)

	for {
		switch ev := s.PollEvent().(type) {
		case *tcell.EventResize:
			s.Sync()
			displayHelloWorld(s)
		case *tcell.EventKey:
			if ev.Key() == tcell.KeyEscape {
				s.Fini()
				os.Exit(0)
			}
		}
	}
}
