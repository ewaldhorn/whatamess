package main

import "core:fmt"
import rl "vendor:raylib"

// ----------------------------------------------------------------------------
drawBasicShapes :: proc(circleAt: rl.Vector2, sides: i32) {
	rl.ClearBackground(GAME_BACKGROUND_COLOUR)
	rl.DrawCircleV(circleAt, 50.0, rl.GREEN)
	rl.DrawRectangleLinesEx({0, 0, SCREEN_WIDTH, SCREEN_HEIGHT}, 10, rl.RED)
	rl.DrawPoly({120, 120}, sides, 80, 0, rl.YELLOW)
	rl.DrawLineBezier({20, 500}, {700, 580}, 4, rl.BROWN)
}

// ----------------------------------------------------------------------------
drawShadowedTextOnImage :: proc(img: ^rl.Image, txt: cstring, textWidth: i32) {
	rl.ImageDrawText(img, txt, (img.width / 2 - textWidth / 2) + 1, img.height - 9, 10, rl.BLACK)
	rl.ImageDrawText(img, txt, img.width / 2 - textWidth / 2, img.height - 10, 10, rl.WHITE)
}

// ----------------------------------------------------------------------------
main :: proc() {
	rl.InitWindow(SCREEN_WIDTH, SCREEN_HEIGHT, GAME_TITLE)
	rl.SetTargetFPS(30)

	nofussLogo := rl.LoadImage("./images/favicon.png")
	text: cstring = "NoFuss Logo"
	textWidth := rl.MeasureText(text, 10)

	drawShadowedTextOnImage(&nofussLogo, text, textWidth)

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
