package main

import "testing"

func Test_calculateGCD(t *testing.T) {
	tests := []struct {
		descr    string
		left     int
		right    int
		expected int
	}{
		{descr: "10 10 gives 10", left: 10, right: 10, expected: 10},
		{descr: "60 36 gives 12", left: 60, right: 36, expected: 12},
		{descr: "36 60 gives 12", left: 36, right: 60, expected: 12},
		{descr: "5 15 gives 5", left: 5, right: 15, expected: 5},
		{descr: "15 5 gives 5", left: 15, right: 5, expected: 5},
	}

	for _, test := range tests {
		result := calculateGCD(test.left, test.right)

		if result != test.expected {
			t.Fatalf("'%s': expected %d, got %d\n", test.descr, test.expected, result)
		}
	}
}
