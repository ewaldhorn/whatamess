package main

import (
	"fmt"
	"os"

	"github.com/gdamore/tcell/v2"
)

// ----------------------------------------------------------------------------
// Initialize the screen for the application.
func initScreen() tcell.Screen {
	screen := getNewScreen()
	setupStyle(screen)

	return screen
}

// ----------------------------------------------------------------------------
func getNewScreen() tcell.Screen {
	newScreen, err := tcell.NewScreen()
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}

	if err := newScreen.Init(); err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(2)
	}

	return newScreen
}

// ----------------------------------------------------------------------------
// Initialize the style for the application.
func setupStyle(screen tcell.Screen) {
	defStyle := tcell.StyleDefault.Background(tcell.ColorBlack).Foreground(tcell.ColorWhite)
	screen.SetStyle(defStyle)
}
