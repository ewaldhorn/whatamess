package main

/*
Calculate the GCD of two numbers using various algorithms.
*/

// -------------------------------------------------------------------------------------- Recursion
func calculateGCD_Recursion(left int, right int) int {
	if right != 0 {
		return calculateGCD_Recursion(right, left%right)
	} else {
		return left
	}
}

// -------------------------------------------------------------------------------------- Euclidian
func calculateGCD_Euclidean(left, right int) int {
	for left != right {
		if left > right {
			left -= right
		} else {
			right -= left
		}
	}

	return left
}

// -------------------------------------------------------------------------------------- Remainder
func calculateGCD_Remainder(left, right int) int {
	for right != 0 {
		left, right = right, left%right
	}

	return left
}
