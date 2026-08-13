package main

import "core:fmt"

// ------------------------------------------------------------------------------------------------
User :: struct {
	id:   int,
	name: string
}

// ------------------------------------------------------------------------------------------------
main :: proc() {
	// fixed array, stored directly in memory
	fixed_scores: [3]int = {90, 85, 95}

	// slice - view into memory region
	score_view: []int = fixed_scores[0:2]

	// dynamic array, resizable, sits on the heap
	users := make([dynamic]User, context.allocator)
	defer delete(users)
}
