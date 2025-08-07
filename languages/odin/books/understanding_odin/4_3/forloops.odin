package main

import "core:fmt"

main :: proc() {
	age := 5

	for {
		if age > 10 {
			break
		}

		fmt.println("Age:", age, "- welcome!")
		age += 1
	}

	fmt.println("Too old for this play section!")

	box(5)
}

stars :: proc(count: int) {
	for i in 0 ..< count {
		fmt.print("*")
	}
}

box :: proc(size: int) {
	for x in 0 ..< size {
		stars(size)
		fmt.println("")
	}
}
