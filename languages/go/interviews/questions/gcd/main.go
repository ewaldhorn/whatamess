package main

/*
Calculate the GCD of two numbers using recursion.
*/
func calculateGCD(left int, right int) int {
	if right != 0 {
		return calculateGCD(right, left%right)
	} else {
		return left
	}
}
