package main

import (
	"github.com/gdamore/tcell"
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
func main() {}
