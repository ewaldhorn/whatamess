package main

import "core:fmt"

// ----------------------------------------------------------------------------
productOf :: proc(left, right: int) -> int {
	return left * right
}

// ----------------------------------------------------------------------------
namedReturnValue :: proc(left, right: int) -> (total: int) {
	total = left + right
	return
}

// ----------------------------------------------------------------------------
divMod :: proc(left, right: int) -> (int, int) {
	return left / right, left % right
}
