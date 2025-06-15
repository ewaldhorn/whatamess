package main

import "core:fmt"

// ----------------------------------------------------------------------------
blanklines :: proc(count: int) {
	for i in 0 ..< count {
		fmt.println()
	}
}
