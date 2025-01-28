package main

import "testing"

// ----------------------------------------------------------------------------
func Test_collatz(t *testing.T) {
	tests := []struct {
		name     string
		starter  int64
		expected bool
	}{
		{"simple", 1, true},
		{"odd", 11, true},
		{"even", 200, true},
		{"big odd", 12383838733, true},
		{"big even", 12844838210, true},
		{"really big odd", 23123921831838219, true},
		{"really big even", 9219329131030210, true},
	}

	for _, test := range tests {
		result := collatz(test.starter)

		if result != test.expected {
			t.Errorf("%s: failed with %t, expected %t", test.name, result, test.expected)
		}
	}
}
