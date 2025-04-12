package main

import "core:fmt"
import "core:math/rand"

main :: proc() {
	// Generate a random number between 1 and 100
	random_number := rand.int31_max(100) + 1

	// Print the random number
	fmt.println(random_number)
}
