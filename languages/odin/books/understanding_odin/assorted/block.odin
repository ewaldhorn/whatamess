package just_the_basics

import "core:fmt"

main :: proc() {
	drawBlock(10, 5)
	drawBlock(10, 10)
	drawBlock(40, 10)
}

// ----------------------------------------------------------------------------
drawBlock :: proc(width: int, height: int) {
	fmt.printfln("\nBlock %dx%d", width, height)
	for y in 0 ..< height {
		for x in 0 ..< width {
			fmt.print("*")
		}
		fmt.println()
	}
}
