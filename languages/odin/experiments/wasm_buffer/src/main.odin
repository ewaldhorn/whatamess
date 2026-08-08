package main
import "base:runtime"

// ------------------------------------------------------------------------------------------------
// Imports from JavaScript
//
foreign import app "app_env"

@(default_calling_convention = "contextless")
foreign app {
	app_init      :: proc(w, h: u32) ---
	app_update :: proc(a: f32) ---
}
// ------------------------------------------------------------------------------------------------
@(export)
bootup :: proc "c" () {
	context = runtime.default_context()
	app_init(SCREEN_WIDTH, SCREEN_HEIGHT)
}
