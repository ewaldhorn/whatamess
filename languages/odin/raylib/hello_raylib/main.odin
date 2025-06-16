package main

import "core:fmt"
import rl "vendor:raylib"
// ----------------------------------------------------------------------------
SCREEN_WIDTH :: 800
SCREEN_HEIGHT :: 600
GAME_TITLE :: "Hello Odin and Raylib"

// ----------------------------------------------------------------------------
main :: proc() {
	rl.InitWindow(SCREEN_WIDTH, SCREEN_HEIGHT, GAME_TITLE)
	rl.SetTargetFPS(30)

	for !rl.WindowShouldClose() {
		rl.BeginDrawing()
		rl.EndDrawing()
	}
}
