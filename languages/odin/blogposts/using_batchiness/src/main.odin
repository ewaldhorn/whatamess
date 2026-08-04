package main

import "../vendor/batchiness/batch"
import "base:runtime" // adjust relative path to match your layout

ctx: batch.Handle
cmd: batch.Buffer

@(export)
batchiness_main :: proc "c" () {
	context = runtime.default_context()
	canvas := batch.canvas_create(batch.get_element_by_id("app"), 800, 600)
	ctx = batch.canvas_get_context(canvas)
	batch.start_animation_loop(0) // callback id 0 = animation frame
}

@(export)
batchiness_invoke_callback :: proc "c" (id: u32) {
	context = runtime.default_context()
	batch.reset(&cmd)
	// --- your drawing code here ---
	batch.set_fill(&cmd, "#ff6347")
	batch.fill_rect(&cmd, 50, 50, 200, 100)
	// --------------------------------
	batch.flush(ctx, &cmd)
}
