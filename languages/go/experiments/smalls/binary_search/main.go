package main

/* Exploring binary searching, based on an email by Jon Calhoun. */

import "fmt"

var numbers = []int{1, 4, 8, 12, 14, 16, 17, 21, 43, 56, 99}

// Good old fashioned brute-force search
func bruteForce(want int) int {
	for index, value := range numbers {
		if value == want {
			return index
		}
	}
	return -1
}

// Returns -1 if the number is not in the list, otherwise it returns
// the index of the value in the sorted list
func searchWithSplit(want int, numbers []int) int {
	if len(numbers) == 0 {
		return -1
	}
	wantIndex := 0
	for len(numbers) > 1 {
		// Split the list in half by finding the middle index
		middle := len(numbers) / 2
		lowerHalf := numbers[0:middle]
		upperHalf := numbers[middle:]

		// Then we ask if our number is in the lower or upper half
		if want < numbers[middle] {
			numbers = lowerHalf
		} else {
			numbers = upperHalf
			// Add the length of the lower half to the index since we are throwing
			// those numbers away, because we know they are all lower than
			// the value we are searching for.
			wantIndex += len(lowerHalf)
		}
	}
	if want != numbers[0] {
		return -1
	}
	return wantIndex
}

// Doesn't split into different lists
func searchV2(want int, numbers []int) int {
	lo, hi := 0, len(numbers)
	// The length of the range we are searching in must have more than one value,
	// otherwise we can stop searching and just look at that value.
	for (hi - lo) > 1 {
		mid := (lo + hi) / 2
		if want < numbers[mid] {
			// Use the lower half
			hi = mid
		} else {
			// Use the upper half
			lo = mid
		}
	}
	// We are done searching, and our value should be at index lo if it is in
	// the list.
	if lo >= len(numbers) || want != numbers[lo] {
		return -1
	}
	return lo
}

func main() {
	lookFor := 43

	fmt.Println()

	fmt.Printf("With bruteForce, the index of 43 is:       %d\n", bruteForce(lookFor))
	fmt.Printf("With searchWithSplit the index of 43 is:   %d\n", searchWithSplit(lookFor, numbers))
	fmt.Printf("With searchV2 the index of 43 is:          %d\n", searchV2(lookFor, numbers))

	fmt.Println()
}
