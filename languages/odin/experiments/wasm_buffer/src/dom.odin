package main

// ------------------------------------------------------------------------------------------------
foreign import dom_env "dom_env"

// ------------------------------------------------------------------------------------------------
@(default_calling_convention = "contextless")
foreign dom_env {
	dom_now                 :: proc() -> f64 ---
}

// ------------------------------------------------------------------------------------------------
// now returns milliseconds since page load (performance.now()), for wall-clock-accurate timing
// (e.g. a sequencer stepping at a fixed BPM regardless of frame rate).
now :: proc "contextless" () -> f64 {
	return dom_now()
}
