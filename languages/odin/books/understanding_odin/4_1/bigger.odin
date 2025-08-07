package main

import "core:fmt"

is_bigger_than :: proc(number: int, compare_to: int) -> bool {
	return number > compare_to
}

main :: proc() {
	number := 7
	fmt.println(number) // prints 7
	number = 12
	fmt.println(number) // prints 12
	fmt.println("7 > 12:", is_bigger_than(7, 12))
	fmt.println("12 > 7:", is_bigger_than(12, 7))
}
