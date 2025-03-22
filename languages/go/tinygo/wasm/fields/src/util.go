package main

// ----------------------------------------------------------------------------
func notZero(left, right int) int {
	if left != 0 {
		return left
	}

	if right != 0 {
		return right
	}

	return 1
}
