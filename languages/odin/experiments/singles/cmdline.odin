package main

import "core:fmt"
import "core:os"
import "core:strconv"

main :: proc() {
	sum := 0
	for a in os.args[1:] {
		n, ok := strconv.parse_int(a)

		if !ok {
			fmt.eprintfln("\tWarning: Could not parse '%s' as a number.", a)
		} else {
			sum += n
		}
	}

	fmt.printfln("The total is %d", sum)
}
