package main

import "core:fmt"

// ----------------------------------------------------------------------------
modifyNumber :: proc(number: ^int) {
	number^ += 2
}

// ----------------------------------------------------------------------------
pointers :: proc() {
	starter := 0
	fmt.printf("Starter is %d right now, before calling modifyNumber.\n", starter)
	modifyNumber(&starter)
	fmt.printf("Starter is %d now after calling modifyNumber. It was changed in-place.\n", starter)
}

// ----------------------------------------------------------------------------
main :: proc() {
	blanklines(2)
	defer blanklines(2)

	defer {
		blanklines(1)
		fmt.println("Tada!")
		fmt.println("And that's all folks!")
	}

	part_one()

	blanklines(2)
	structs()
}
