package main

import "core:fmt"

// ----------------------------------------------------------------------------
// ----------------------------------------------------------------------------
// ----------------------------------------------------------------------------
productOf :: proc(left, right: int) -> int {
	return left * right
}

// ----------------------------------------------------------------------------
blanklines :: proc(count: int) {
	fmt.println()
}
// ----------------------------------------------------------------------------
main :: proc() {
	blanklines(2)

	fmt.println("Hello world, from Odin!")

	x := 5
	y := 3
	fmt.printf("The product of %d and %d is %d.\n", x, y, productOf(x, y))

	blanklines(2)
}
