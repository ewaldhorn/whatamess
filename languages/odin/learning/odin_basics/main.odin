package main

import "core:fmt"

// ----------------------------------------------------------------------------
productOf :: proc(left, right: int) -> int {
	return left * right
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
	blanklines(2)

	sample := "sample"
	fmt.printf("Looking at \"sample\":\n")
	printLettersInString(sample)
	blanklines(1)

	fmt.printf("Looking at letter values (ASCII) now:\n")
	printValuesOfLettersInString(sample)

	sample = "0123456789"
	fmt.printf("Looking at letter values (HEX, ASCII) now:\n")
	printHexAndBinaryValuesOfLettersInString(sample)

	blanklines(2)
	fmt.println("Using labels in loops:")
	loopLabels()

	blanklines(2)
}
