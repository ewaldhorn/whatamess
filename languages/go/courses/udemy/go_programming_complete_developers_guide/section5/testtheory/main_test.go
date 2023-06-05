package main

import "testing"

func TestIsValidCalculation(t *testing.T) {
	expected := 10
	received := sum(1, 2, 3, 4)

	if expected != received {
		t.Errorf("The sum of 1,2,3 and 4 should be 10!, not %d", received)
	}
}

func TestMultipleCalculations(t *testing.T) {
	tests := []struct {
		values []int
		want   int
	}{
		{[]int{1, 2, 3, 4}, 10},
		{[]int{1, 1, 1, 1}, 4},
		{[]int{1, 2, 1, 2}, 6},
		{[]int{1, 2, 2, 2}, 7},
	}

	for _, test := range tests {
		received := sum(test.values...)

		if received != test.want {
			t.Errorf("The sum of %v should be %d!, not %d", test.values, test.want, received)
		}
	}
}
