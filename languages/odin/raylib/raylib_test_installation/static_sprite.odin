package game

import rl "vendor:raylib"

main :: proc() {
	rl.InitWindow(1280, 720, "My Odin + Raylib game")
	// NEW LINE:
	player := rl.LoadTexture("player.png")

	for !rl.WindowShouldClose() {
		rl.BeginDrawing()
		rl.ClearBackground({160, 200, 255, 255})
		// NEW LINE:
		rl.DrawTextureV(player, {60, 40}, rl.WHITE)
		rl.EndDrawing()
	}

	rl.CloseWindow()
}

