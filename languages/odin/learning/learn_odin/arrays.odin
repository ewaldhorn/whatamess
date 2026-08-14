package main

import "core:fmt"

// ------------------------------------------------------------------------------------------------
// Type alias: Vector4 is a shorthand for an array of 4 f32 (32-bit float) values.
// The `::` keyword defines a type synonym, so `Vector4` and `[4]f32` are interchangeable.
// ------------------------------------------------------------------------------------------------
Vector4 :: [4]f32

// ------------------------------------------------------------------------------------------------
main :: proc() {
	// Element-wise initialization: creates a Vector4 with values 1, 2, 3, 4.
	a := Vector4{1, 2, 3, 4}
	// Element-wise initialization: creates a Vector4 with values 10, 20, 30, 40.
	b := Vector4{10, 20, 30, 40}

	// Traditional element-by-element loop:
	// Declares `sum_loop` as an uninitialized Vector4, then fills each element
	// by iterating indices 0..3 (half-open range `0 ..< 4`).
	sum_loop: Vector4
	for i in 0 ..< 4 {
		sum_loop[i] = a[i] + b[i]
	}

	fmt.println("Loop sum is :", sum_loop)

	// Array programming: `a + b` adds corresponding elements of both arrays
	// in a single expression, producing a new array with the element-wise sums.
	sum := a + b
	fmt.println("Array sum is:", sum)
}
