package main

import "core:fmt"
import rl "vendor:raylib"
// ----------------------------------------------------------------------------
SCREEN_WIDTH :: 800
SCREEN_HEIGHT :: 600
GAME_TITLE :: "Hello Odin and Raylib"

// ----------------------------------------------------------------------------
drawBasicShapes :: proc(circleAt: rl.Vector2, sides: i32) {
	rl.ClearBackground(rl.BLACK)
	rl.DrawCircleV(circleAt, 50.0, rl.GREEN)
	rl.DrawRectangleLinesEx({0, 0, SCREEN_WIDTH, SCREEN_HEIGHT}, 10, rl.RED)
	rl.DrawPoly({120, 120}, sides, 80, 0, rl.YELLOW)
	rl.DrawLineBezier({20, 500}, {700, 580}, 4, rl.BROWN)
}

// ----------------------------------------------------------------------------


// ----------------------------------------------------------------------------
main :: proc() {
	rl.InitWindow(SCREEN_WIDTH, SCREEN_HEIGHT, GAME_TITLE)
	rl.SetTargetFPS(30)

	nofussLogo := rl.LoadImage("./images/favicon.png")
	text: cstring = "NoFuss Logo"
	textWidth := rl.MeasureText(text, 10)

	rl.ImageDrawText(
		&nofussLogo,
		text,
		(nofussLogo.width / 2 - textWidth / 2) + 1,
		nofussLogo.height - 9,
		10,
		rl.BLACK,
	)

	rl.ImageDrawText(
		&nofussLogo,
		text,
		nofussLogo.width / 2 - textWidth / 2,
		nofussLogo.height - 10,
		10,
		rl.WHITE,
	)


	texture := rl.LoadTextureFromImage(nofussLogo)
	//also rl.LoadTexture("./images/favicon.png")

	circlePos: rl.Vector2 = {SCREEN_WIDTH - 200, SCREEN_HEIGHT / 2}
	sides: i32 = 3

	for !rl.WindowShouldClose() {
		rl.BeginDrawing()

		// if rl.IsMouseButtonPressed(.LEFT) {
		// circlePos = rl.GetMousePosition()
		// }

		if rl.IsMouseButtonDown(.LEFT) {
			circlePos = rl.GetMousePosition()
		}

		if rl.IsKeyDown(.UP) {
			if sides < 24 {
				sides += 1
			}
		}

		if rl.IsKeyDown(.DOWN) {
			if sides > 3 {
				sides -= 1
			}
		}

		drawBasicShapes(circlePos, sides)

		rl.DrawTexture(texture, 350, 300, rl.WHITE)

		rl.EndDrawing()
	}
}
