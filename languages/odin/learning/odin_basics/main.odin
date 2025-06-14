package main

import "core:fmt"

// ----------------------------------------------------------------------------
// ----------------------------------------------------------------------------
printOddNumbersOnly :: proc(range: int) {
	for i := 0; i <= range; i += 1 {
		if i % 2 != 0 {
			fmt.printf("%d, ", i)
		}
	}
}
// ----------------------------------------------------------------------------
productOf :: proc(left, right: int) -> int {
	return left * right
}

// ----------------------------------------------------------------------------
blanklines :: proc(count: int) {
	for i in 0 ..< count {
		fmt.println()
	}
}
// ----------------------------------------------------------------------------
main :: proc() {
	blanklines(2)

	fmt.println("Hello world, from Odin!")

	x := 5
	y := 3
	fmt.printf("The product of %d and %d is %d.\n", x, y, productOf(x, y))

	blanklines(1)
	fmt.print("Odd numbers under 20: ")
	printOddNumbersOnly(20)
	blanklines(1)

	blanklines(2)
}
