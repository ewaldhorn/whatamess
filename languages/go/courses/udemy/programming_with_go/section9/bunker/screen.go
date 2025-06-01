package main

import (
	"fmt"
	"os"

	"github.com/gdamore/tcell/v2"
)

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
