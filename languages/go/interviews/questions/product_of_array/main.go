package main

// Given an array, calculate the product of the numbers contained in the array,
// except for the current number.
//
// For example:
// Given [1, 2, 3, 4], output [24, 12, 8, 6]

func productOfArrayExceptForSelf(numbers []int) []int {
	output := make([]int, len(numbers))
	for outer := 0; outer < len(numbers); outer++ {
		tmp := 1

		for inner := 0; inner < len(numbers); inner++ {
			if inner != outer {
				tmp *= numbers[inner]
			}
		}

		output[outer] = tmp
	}

	return output
}
