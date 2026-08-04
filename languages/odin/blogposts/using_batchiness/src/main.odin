package main

import "../vendor/batchiness/batch"
import "base:runtime"
import "core:math"

// ------------------------------------------------------------------------------------------------
// Canvas dimensions
// ------------------------------------------------------------------------------------------------
CANVAS_W :: 800
CANVAS_H :: 600

// ------------------------------------------------------------------------------------------------
// Callback IDs
// ------------------------------------------------------------------------------------------------
CB_FRAME :: 0
CB_CLICK :: 1

// ------------------------------------------------------------------------------------------------
// Ball colours — two palettes, one per ball, cycled on click
// ------------------------------------------------------------------------------------------------
BALL_COLOURS := [2][6]string{
	{"#e63946", "#f4a261", "#e9c46a", "#2a9d8f", "#457b9d", "#9b5de5"},
	{"#f72585", "#7209b7", "#3a0ca3", "#4361ee", "#4cc9f0", "#06d6a0"},
}

// ------------------------------------------------------------------------------------------------
// App state (package-level globals, zero-initialised)
// ------------------------------------------------------------------------------------------------
ctx:       batch.Handle
cmd:       batch.Buffer
canvas_el: batch.Handle

Ball :: struct {
	x, y:    f32,
	vx, vy:  f32,
	r:       f32,
	col_idx: int,
	palette: int,
}

balls: [2]Ball

// ------------------------------------------------------------------------------------------------
// Custom JS functions provided via extraImports in index.html.
// ------------------------------------------------------------------------------------------------
foreign import app_env "app_env"

@(default_calling_convention = "contextless")
foreign app_env {
	get_last_click_x :: proc() -> f32 ---
	get_last_click_y :: proc() -> f32 ---
}

// ------------------------------------------------------------------------------------------------
// dist2 — squared Euclidean distance
// ------------------------------------------------------------------------------------------------
dist2 :: proc "contextless" (ax, ay, bx, by: f32) -> f32 {
	dx := ax - bx
	dy := ay - by
	return dx * dx + dy * dy
}

// ------------------------------------------------------------------------------------------------
// draw_ball — records a filled circle + a radial-gradient highlight into cmd
// ------------------------------------------------------------------------------------------------
draw_ball :: proc "contextless" (b: ^Ball, grad_id: u16) {
	color := BALL_COLOURS[b.palette][b.col_idx]

	// Radial gradient: bright highlight at top-left → ball colour
	batch.radial_gradient(
		&cmd,
		grad_id,
		b.x - b.r * 0.35,
		b.y - b.r * 0.35,
		b.r * 0.1, // inner circle (highlight)
		b.x,
		b.y,
		b.r, // outer circle
	)
	batch.add_color_stop(&cmd, grad_id, 0, "#ffffff")
	batch.add_color_stop(&cmd, grad_id, 0.4, color)
	batch.add_color_stop(&cmd, grad_id, 1, "#000000")
	batch.use_gradient_fill(&cmd, grad_id)

	// Draw the circle
	batch.begin_path(&cmd)
	batch.arc(&cmd, b.x, b.y, b.r, 0, math.TAU)
	batch.fill(&cmd)

	// Subtle stroke
	batch.set_stroke(&cmd, "#ffffff44")
	batch.set_line_width(&cmd, 1.5)
	batch.stroke(&cmd)
}

// ------------------------------------------------------------------------------------------------
// batchiness_main — called once by JS after WASM instantiation
// ------------------------------------------------------------------------------------------------
@(export)
batchiness_main :: proc "c" () {
	context = runtime.default_context()

	canvas_el = batch.canvas_create(batch.get_element_by_id("app"), CANVAS_W, CANVAS_H)
	ctx = batch.canvas_get_context(canvas_el)

	// Initialise balls
	balls[0] = Ball {
		x       = 200,
		y       = 200,
		vx      = 1.2,
		vy      = 1.1,
		r       = 40,
		col_idx = 0,
		palette = 0,
	}
	balls[1] = Ball {
		x       = 580,
		y       = 380,
		vx      = -1.5,
		vy      = 1.4,
		r       = 50,
		col_idx = 0,
		palette = 1,
	}

	// Register callbacks (using pointerdown for unified, instant touch + click response)
	batch.start_animation_loop(CB_FRAME)
	batch.add_event_listener(canvas_el, "pointerdown", CB_CLICK)
}

// ------------------------------------------------------------------------------------------------
// update_ball — advance physics and bounce off walls
// ------------------------------------------------------------------------------------------------
update_ball :: proc "contextless" (b: ^Ball) {
	b.x += b.vx
	b.y += b.vy

	if b.x - b.r < 0 {
		b.x = b.r
		b.vx = abs(b.vx)
	} else if b.x + b.r > CANVAS_W {
		b.x = CANVAS_W - b.r
		b.vx = -abs(b.vx)
	}

	if b.y - b.r < 0 {
		b.y = b.r
		b.vy = abs(b.vy)
	} else if b.y + b.r > CANVAS_H {
		b.y = CANVAS_H - b.r
		b.vy = -abs(b.vy)
	}
}

// ------------------------------------------------------------------------------------------------
// batchiness_invoke_callback — dispatched for animation frames and DOM events
// ------------------------------------------------------------------------------------------------
@(export)
batchiness_invoke_callback :: proc "c" (id: u32) {
	context = runtime.default_context()

	switch id {

	// ---- Animation frame -----------------------------------------------------------------------
	case CB_FRAME:
		// Advance physics
		update_ball(&balls[0])
		update_ball(&balls[1])

		batch.reset(&cmd)

		// Dark background
		batch.set_fill(&cmd, "#0d1117")
		batch.fill_rect(&cmd, 0, 0, CANVAS_W, CANVAS_H)

		// Grid lines for depth
		batch.set_stroke(&cmd, "#ffffff0d")
		batch.set_line_width(&cmd, 1)
		for gx: f32 = 0; gx <= CANVAS_W; gx += 80 {
			batch.begin_path(&cmd)
			batch.move_to(&cmd, gx, 0)
			batch.line_to(&cmd, gx, CANVAS_H)
			batch.stroke(&cmd)
		}
		for gy: f32 = 0; gy <= CANVAS_H; gy += 80 {
			batch.begin_path(&cmd)
			batch.move_to(&cmd, 0, gy)
			batch.line_to(&cmd, CANVAS_W, gy)
			batch.stroke(&cmd)
		}

		// Soft glow shadow per ball
		batch.set_shadow(&cmd, BALL_COLOURS[balls[0].palette][balls[0].col_idx], 24)
		draw_ball(&balls[0], 1)
		batch.set_shadow(&cmd, BALL_COLOURS[balls[1].palette][balls[1].col_idx], 24)
		draw_ball(&balls[1], 2)
		batch.clear_shadow(&cmd)

		// HUD label
		batch.set_fill(&cmd, "#ffffff66")
		batch.set_font(&cmd, "14px monospace")
		batch.set_text_align(&cmd, "left")
		batch.fill_text(&cmd, "tap / click a ball to change its colour", 12, CANVAS_H - 12)

		batch.flush(ctx, &cmd)

	// ---- Click / touch -------------------------------------------------------------------------
	case CB_CLICK:
		mx := get_last_click_x()
		my := get_last_click_y()

		for i in 0 ..< 2 {
			b := &balls[i]
			// Give a slightly generous hit radius (+12px) for easy tapping on moving targets
			hit_r := b.r + 12
			if dist2(mx, my, b.x, b.y) <= hit_r * hit_r {
				b.col_idx = (b.col_idx + 1) % len(BALL_COLOURS[b.palette])
			}
		}
	}
}
