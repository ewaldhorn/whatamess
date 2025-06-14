package main

import (
	"github.com/gdamore/tcell/v2"
	"github.com/mattn/go-runewidth"
)

// ----------------------------------------------------------------------------
func display(screen tcell.Screen, x, y, w, h int, ch rune) {
	for r := range w {
		for c := range h {
			screen.SetContent(x+r, y+c, ch, nil, tcell.StyleDefault.Foreground(tcell.ColorYellow).Background(tcell.ColorBlue))
		}
	}
}

// ----------------------------------------------------------------------------
func displayText(screen tcell.Screen, x, y int, style tcell.Style, text string) {
	for _, character := range text {
		var comb []rune
		width := runewidth.RuneWidth(character)
		if width == 0 {
			comb = []rune{character}
			character = ' '
			width = 1
		}
		screen.SetContent(x, y, character, comb, style)
		x += width
	}
}
