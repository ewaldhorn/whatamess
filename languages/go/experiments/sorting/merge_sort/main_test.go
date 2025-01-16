package main

import "testing"

// ----------------------------------------------------------------------------
func Test_MergeSort(t *testing.T) {
	items := []int{9, 8, 7, 6, 5, 4, 3, 2, 1}
	result := mergeSort(items)
	for i := 0; i < len(result)-1; i++ {
		if result[i] > result[i+1] {
			t.Errorf("MergeSort failed with %v", result)
		}
	}
}

// ----------------------------------------------------------------------------
func Test_MergeSort_AlreadySorted(t *testing.T) {
	items := []int{2, 4, 6, 8, 10, 12}
	result := mergeSort(items)

	for i := 0; i < len(result); i++ {
		if result[i] != items[i] {
			t.Errorf("MergeSort_AlreadySorted failed with %v", result)
		}
	}
}

// ----------------------------------------------------------------------------
func Test_MergeSort_Empty(t *testing.T) {
	items := []int{}
	result := mergeSort(items)

	if len(result) != len(items) {
		t.Errorf("MergeSort_Empty failed with %d, expected 0.", len(result))
	}
}

// ----------------------------------------------------------------------------
func Test_MergeSort_Duplicates(t *testing.T) {
	items := []int{9, 8, 7, 6, 7, 5, 4, 3, 2, 9, 9, 9, 1}
	result := mergeSort(items)

	if len(result) != len(items) {
		t.Errorf("MergeSort_Duplicates failed with %d, expected %d", len(result), len(items))
		t.Errorf("Result: %v  vs Items: %v", result, items)

	}

	for i := 0; i < len(result)-1; i++ {
		if result[i] > result[i+1] {
			t.Errorf("MergeSort_Duplicates failed with %v", result)
		}
	}
}
