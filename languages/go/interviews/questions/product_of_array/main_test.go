package main

import "testing"

func Test_productOfArray(t *testing.T) {
	tests := []struct {
		descr    string
		input    []int
		expected []int
	}{
		{descr: "Basic", input: []int{1, 2, 3, 4}, expected: []int{24, 12, 8, 6}},
		{descr: "WithNegs", input: []int{-1, 1, 0, -3, 3}, expected: []int{0, 0, 9, 0, 0}},
	}

	for _, test := range tests {
		result := productOfArrayExceptForSelf(test.input)

		for i := 0; i < len(test.input); i++ {
			if result[i] != test.expected[i] {
				t.Fatalf("expected %v, got %v\n", test.expected, result)
			}
		}
	}
}
