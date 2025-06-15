package main

import "core:fmt"

// ----------------------------------------------------------------------------
Directions :: enum {
	North,
	South,
	East,
	West,
}

// ----------------------------------------------------------------------------
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
