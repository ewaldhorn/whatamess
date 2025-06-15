package main

import "core:fmt"

part_one :: proc() {
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
	blanklines(2)

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
	fmt.printf("The total of 6 and 8 is %d.\n", namedReturnValue(6, 8))

	blanklines(2)
	left, right := 17, 5
	d, m := divMod(left, right)
	fmt.printf("DivMod for %d and %d is %d and %d.\n", left, right, d, m)
}
