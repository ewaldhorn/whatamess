package main

import (
	"errors"
)

// ----------------------------------------------------------------------------
// Assumes a sorted values array, otherwise it just won't work
func basicBinarySearch(sortedIntArray []int, itemToFind int) int {
	low, high := 0, len(sortedIntArray)-1

	for low <= high {
		mid := (low + high) / 2

		if sortedIntArray[mid] < itemToFind {
			low = mid + 1
		} else if sortedIntArray[mid] > itemToFind {
			high = mid - 1
		} else {
			return mid
		}
	}

	return -1
}

// ----------------------------------------------------------------------------
// Improved binary search with some better error handling
func binarySearch(sortedArray []int, target int) (int, error) {
	if len(sortedArray) == 0 {
		return -1, errors.New("array is empty")
	}

	// Check if the array is sorted
	for i := 1; i < len(sortedArray); i++ {
		if sortedArray[i-1] > sortedArray[i] {
			return -1, errors.New("array is not sorted")
		}
	}

	low := 0
	high := len(sortedArray) - 1

	for low <= high {
		mid := (low + high) / 2

		if sortedArray[mid] < target {
			low = mid + 1
		} else if sortedArray[mid] > target {
			high = mid - 1
		} else {
			return mid, nil
		}
	}

	return -1, errors.New("target not found in array")
}
