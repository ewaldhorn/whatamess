package main

import "core:fmt"

// ----------------------------------------------------------------------------
productOf :: proc(left, right: int) -> int {
	return left * right
}

// ----------------------------------------------------------------------------
Directions :: enum {
	North,
	South,
	East,
	West,
}

enums :: proc() {
	current_direction := Directions.West

	switch current_direction {
	case Directions.North:
		fmt.println("Heading North")
	case Directions.South:
		fmt.println("Heading South")
	case Directions.East:
		fmt.println("Heading East")
	case Directions.West:
		fmt.println("Heading West")
	}

	#partial switch current_direction {
	case Directions.North:
		fmt.println("Heading North")
	case Directions.South:
		fmt.println("Heading South")
	case:
		fmt.println("Going somewhere, but I don't care about that direction")
	}
}

// ----------------------------------------------------------------------------
switches :: proc(grade: int) {
	switch grade {
	case 0 ..< 50:
		fmt.println("FAIL")
	case 50 ..< 80:
		fmt.println("PASS")
	case:
		fmt.println("DISTINCTION")
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
	enums()

	blanklines(2)
	fmt.println("Doing some grades...")

	fmt.print("A grade of 40: ")
	switches(40)
	fmt.print("A grade of 50: ")
	switches(50)
	fmt.print("A grade of 70: ")
	switches(70)
	fmt.print("A grade of 92: ")
	switches(92)

	blanklines(2)
}
