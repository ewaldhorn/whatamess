package game

import rl "vendor:raylib"
// NEW LINE:
import la "core:math/linalg"

main :: proc() {
	rl.InitWindow(1280, 720, "My Odin + Raylib game")
	player := rl.LoadTexture("player.png")
	// NEW LINE:
	player_pos: [2]f32

	for !rl.WindowShouldClose() {
		// NEW BLOCK START
		input: [2]f32

		if rl.IsKeyDown(.UP) {
			input.y -= 10
		}
		if rl.IsKeyDown(.DOWN) {
			input.y += 10
		}
		if rl.IsKeyDown(.LEFT) {
			input.x -= 10
		}
		if rl.IsKeyDown(.RIGHT) {
			input.x += 10
		}

		player_pos += la.normalize0(input) * rl.GetFrameTime()
		// NEW BLOCK END

		rl.BeginDrawing()
		rl.ClearBackground({160, 200, 255, 255})
		// MODIFIED LINE (use player_pos):
		rl.DrawTextureV(player, player_pos, rl.WHITE)
		rl.EndDrawing()
	}

	rl.CloseWindow()
}

