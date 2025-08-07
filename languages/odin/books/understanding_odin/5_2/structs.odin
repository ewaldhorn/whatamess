package structs

import "core:fmt"

// ----------------------------------------------------------------------------
Entity :: struct {
	id:       int,
	position: [2]f32,
}

// ----------------------------------------------------------------------------
Player :: struct {
	using entity: Entity,
	can_jump:     bool,
}

// ----------------------------------------------------------------------------
print_position :: proc(e: Entity) {
	fmt.println(e.position) // [5, 2]
}

// ----------------------------------------------------------------------------
main :: proc() {
	p := Player {
		id       = 7,
		position = {5, 2},
		can_jump = true,
	}

	fmt.println(p.position) // [5, 2]
	print_position(p) // [5, 2]
}
