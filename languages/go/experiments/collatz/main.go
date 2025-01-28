package main

// ----------------------------------------------------------------------------
func collatz(num int64) bool {
	if num == 1 {
		return true
	}

	if num%2 == 0 {
		return collatz(num / 2)
	} else {
		return collatz(3*num + 1)
	}
}
