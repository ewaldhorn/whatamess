package main

import "core:fmt"


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

}
