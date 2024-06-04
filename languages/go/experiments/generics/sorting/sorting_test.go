package main

import (
	"slices"
	"testing"
)

// ------------------------------------------------------------------------------ Test_sorting_int8
func Test_sorting_int8(t *testing.T) {
	tests := []struct {
		name     string
		input    []int8
		expected []int8
	}{
		{name: "bytes", input: []int8{1, 3, 4, 2, 5}, expected: []int8{1, 2, 3, 4, 5}},
		{name: "bytes - already sorted", input: []int8{1, 2, 3, 4, 5}, expected: []int8{1, 2, 3, 4, 5}},
	}

	for number, test := range tests {
		result := EasySort(test.input)

		if !slices.Equal(test.expected, result) {
			t.Fatalf("test %d failed with result %v vs expected %v\n", number, result, test.expected)
		}
	}
}

// ----------------------------------------------------------------------------- Test_sorting_int16
func Test_sorting_int16(t *testing.T) {
	tests := []struct {
		name     string
		input    []int16
		expected []int16
	}{
		{name: "bytes", input: []int16{1, 3, 4, 2, 5}, expected: []int16{1, 2, 3, 4, 5}},
		{name: "bytes - already sorted", input: []int16{1, 2, 3, 4, 5}, expected: []int16{1, 2, 3, 4, 5}},
	}

	for number, test := range tests {
		result := EasySort(test.input)

		if !slices.Equal(test.expected, result) {
			t.Fatalf("test %d failed with result %v vs expected %v\n", number, result, test.expected)
		}
	}
}
