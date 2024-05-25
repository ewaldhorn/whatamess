package main

// returning multiple values from a function
func minMaxOf(values []int) (int, int) {
	min := values[0]
	max := values[0]

	for _, val := range values {
		if val < min {
			min = val
		}

		if val > max {
			max = val
		}
	}

	return min, max
}
