package main

import "core:fmt"

// ----------------------------------------------------------------------------
loopLabels :: proc() {
	outer: for i in 0 ..< 3 {
		for j in 30 ..< 35 {
			if j == 32 {
				continue outer
			}
			fmt.println(i, j)
		}
		fmt.println("last line in outer loop")
	}
}
// ----------------------------------------------------------------------------
printHexAndBinaryValuesOfLettersInString :: proc(src: string) {
	for i in 0 ..< len(src) {
		fmt.printf("%#x (%#b)\n", src[i], src[i])
	}
}

// ----------------------------------------------------------------------------
printValuesOfLettersInString :: proc(src: string) {
	for i in 0 ..< len(src) {
		fmt.println(src[i])
	}
}

// ----------------------------------------------------------------------------
printLettersInString :: proc(src: string) {
	for c, pos in src {
		fmt.printf("We have '%c' at position %d\n", c, pos)
	}
}

// ----------------------------------------------------------------------------
printOddNumbersOnly :: proc(range: int) {
	for i := 0; i <= range; i += 1 {
		if i % 2 != 0 {
			fmt.printf("%d, ", i)
		}
	}
}
