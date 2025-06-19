package main

import rl "vendor:raylib"

centeredText :: proc(txt: cstring, yPos, size: i32, colour: rl.Color) {
	textWidth := rl.MeasureText(txt, size)
	x := i32(SCREEN_WIDTH / 2 - textWidth / 2)
	y := yPos - size / 2
	rl.DrawText(txt, x, y, size, colour)
}
