package arrays

import "core:fmt"

// ------------------------------------------------------------------------------------------------
Vector4 :: [4]f32

// ------------------------------------------------------------------------------------------------
main :: proc() {
	a := Vector4{1, 2, 3, 4}
	b := Vector4{10, 20, 30, 40}

	// traditional looping method
	sum_loop: Vector4
	for i in 0 ..< 4 {
		sum_loop[i] = a[i] + b[i]
	}

	fmt.println("Loop sum is:", sum_loop)
}
