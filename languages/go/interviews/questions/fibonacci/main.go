package main

import "math"

/*
	Various ways to calculate the Fibonacci sequence.
*/

// ----------------------------------------------------------------------------------- By Recursion
func fibonacci_recursive(num int) int {
	if num == 0 {
		return num
	}

	if num < 3 {
		return 1
	}

	return fibonacci_recursive(num-1) + fibonacci_recursive(num-2)
}

// ----------------------------------------------------------------------------------- By Iteration
func fibonacci_iterative(num int) int {
	nums := []int{0, 1}

	for i := 2; i <= num; i++ {
		next := nums[i-1] + nums[i-2]
		nums = append(nums, next)
	}

	return nums[num]
}

// -------------------------------------------------------------------------------- By Golden Ratio
// This does get inaccurate at some stage due to floating-point math errors.
func fibonacci_golden_ration(num int) int {
	numerator := math.Pow(1.618034, float64(num)) - math.Pow(-0.618034, float64(num))
	denominator := math.Sqrt(5)

	return int(numerator / denominator)
}
