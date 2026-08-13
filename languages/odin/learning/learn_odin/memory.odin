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

	append(&users, User{id=1,name="Allsop"})
	append(&users, User{2,"Jones"})

	// temporary scratch buffer
	scratch_buffer := make([]u8, 512, context.temp_allocator)

	fmt.println("Fixed scores  :", fixed_scores)
	fmt.println("Slice view    :", score_view)
	fmt.println("Dynamic users :", users)
	fmt.printf("Scratch buffer : %v bytes\n", len(scratch_buffer))
}
