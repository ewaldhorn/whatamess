package main

import "testing"

func Test_BasicSum(t *testing.T) {
	const val1 = 20
	const val2 = 30
	const expected = 50
	result := goSum(val1, val2)

	if expected != result {
		t.Fatalf("Expected %d, got %d.\n", expected, result)
	}
}
