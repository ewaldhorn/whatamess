package main

import "testing"

// ----------------------------------------------------------------------------
func Test_CalculateLargeNumber(t *testing.T) {
	const expected int64 = 15432092901235
	result := CalculateLargeNumber()

	if result != expected {
		t.Errorf("%s: failed with %d, expected %d", "Test_CalculateLargeNumber", result, expected)
	}
}
