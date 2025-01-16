package main

// ----------------------------------------------------------------------------
func mergeSort(items []int) []int {
	// pointless to sort less than 2 items
	if len(items) < 2 {
		return items
	}

	midPoint := len(items) / 2

	left := mergeSort(items[:midPoint])  // 0..midPoint-1
	right := mergeSort(items[midPoint:]) // midPoint..end

	return mergeLists(left, right)
}

// ----------------------------------------------------------------------------
func mergeLists(left, right []int) (result []int) {

	leftPos, rightPos := 0, 0

	// go through each list, comparing items one by one
	for leftPos < len(left) && rightPos < len(right) {
		if left[leftPos] < right[rightPos] {
			result = append(result, left[leftPos])
			leftPos++
		} else {
			result = append(result, right[rightPos])
			rightPos++
		}
	}

	result = append(result, left[leftPos:]...)
	result = append(result, right[rightPos:]...)

	return
}
