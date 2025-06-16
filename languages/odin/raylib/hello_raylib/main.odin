package main

import "core:fmt"
import rl "vendor:raylib"
// ----------------------------------------------------------------------------
SCREEN_WIDTH :: 800
SCREEN_HEIGHT :: 600
GAME_TITLE :: "Hello Odin and Raylib"

// ----------------------------------------------------------------------------
drawBasicShapes :: proc(circleAt: rl.Vector2) {
	rl.DrawCircleV(circleAt, 50.0, rl.GREEN)
	rl.DrawRectangleLinesEx({0, 0, SCREEN_WIDTH, SCREEN_HEIGHT}, 10, rl.RED)
	rl.DrawPoly({90, 90}, 5, 50, 0, rl.YELLOW)
	rl.DrawLineBezier({20, 500}, {700, 580}, 4, rl.BROWN)
}

// ----------------------------------------------------------------------------
main :: proc() {
	rl.InitWindow(SCREEN_WIDTH, SCREEN_HEIGHT, GAME_TITLE)
	rl.SetTargetFPS(30)

	circlePos: rl.Vector2 = {SCREEN_WIDTH / 2, SCREEN_HEIGHT / 2}

	for !rl.WindowShouldClose() {
		rl.BeginDrawing()

		if rl.IsMouseButtonPressed(.LEFT) {
			circlePos = rl.GetMousePosition()
		}

		drawBasicShapes(circlePos)

		rl.EndDrawing()
	}
}
