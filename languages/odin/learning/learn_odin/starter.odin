package main

import "core:fmt"

main :: proc() {
	age := 20 //variable declared, type infered
	target_distance: f64 = 42.195 // explicit type declaration
	MAX_VELOCITY :: 300 //compile-time constant

	age_as_float := f64(age) // explicit cast, Odin does not do implicit casting

	fmt.printf(
		"Age: %v, Distance %v, Constant %v, Cast:%v\n",
		age,
		target_distance,
		MAX_VELOCITY,
		age_as_float,
	)
}
