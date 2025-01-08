package main

import (
	"errors"
	"testing"
)

// ----------------------------------------------------------------------------
func Test_basicBinarySearch(t *testing.T) {
	tests := []struct {
		description    string
		itemList       []int
		searchFor      int
		expectedResult int
	}{
		{
			description:    "simple binary search - success",
			itemList:       []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
			searchFor:      6,
			expectedResult: 5,
		},
		{
			description:    "simple binary search - not found",
			itemList:       []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
			searchFor:      0,
			expectedResult: -1,
		},
	}

	for _, test := range tests {
		result := basicBinarySearch(test.itemList, test.searchFor)
		if result != test.expectedResult {
			t.Errorf("%s: failed with %d, expected %d", test.description, result, test.expectedResult)
		}
	}
}

// ----------------------------------------------------------------------------
func Test_improvedBinarySearch(t *testing.T) {
	tests := []struct {
		description         string
		itemList            []int
		searchFor           int
		expectedResultValue int
		expectedError       error
	}{
		{
			description:         "simple binary search - success",
			itemList:            []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
			searchFor:           6,
			expectedResultValue: 5,
			expectedError:       nil,
		},
		{
			description:         "simple binary search - not found",
			itemList:            []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
			searchFor:           0,
			expectedResultValue: -1,
			expectedError:       errors.New("target not found in array"),
		},
		{
			description:         "simple binary search - empty array",
			itemList:            []int{},
			searchFor:           0,
			expectedResultValue: -1,
			expectedError:       errors.New("array is empty"),
		},
		{
			description:         "simple binary search - not sorted",
			itemList:            []int{1, 2, 3, 4, 6, 5, 7, 8, 9, 10},
			searchFor:           0,
			expectedResultValue: -1,
			expectedError:       errors.New("array is not sorted"),
		},
	}

	for _, test := range tests {
		resultValue, resultError := binarySearch(test.itemList, test.searchFor)
		if resultValue != test.expectedResultValue {
			t.Errorf("%s: failed with %d, expected %d", test.description, resultValue, test.expectedResultValue)
		}

		if resultError != nil && resultError.Error() != test.expectedError.Error() {
			t.Errorf("%s: failed with '%v', expected '%v'", test.description, resultError, test.expectedError)
		}
	}
}

// ----------------------------------------------------------------------------
// AI Generated tests, manually tweaked error message comparison to fix that.
func Test_improvedBinarySearchForCoverage(t *testing.T) {
	tests := []struct {
		name        string
		sortedArray []int
		target      int
		wantIndex   int
		wantError   error
	}{
		{
			name:        "found in middle",
			sortedArray: []int{1, 2, 3, 4, 5},
			target:      3,
			wantIndex:   2,
			wantError:   nil,
		},
		{
			name:        "found at start",
			sortedArray: []int{1, 2, 3, 4, 5},
			target:      1,
			wantIndex:   0,
			wantError:   nil,
		},
		{
			name:        "found at end",
			sortedArray: []int{1, 2, 3, 4, 5},
			target:      5,
			wantIndex:   4,
			wantError:   nil,
		},
		{
			name:        "not found",
			sortedArray: []int{1, 2, 3, 4, 5},
			target:      6,
			wantIndex:   -1,
			wantError:   errors.New("target not found in array"),
		},
		{
			name:        "empty array",
			sortedArray: []int{},
			target:      1,
			wantIndex:   -1,
			wantError:   errors.New("array is empty"),
		},
		{
			name:        "unsorted array",
			sortedArray: []int{5, 2, 8, 1, 9},
			target:      1,
			wantIndex:   -1,
			wantError:   errors.New("array is not sorted"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotIndex, gotError := binarySearch(tt.sortedArray, tt.target)
			if gotIndex != tt.wantIndex {
				t.Errorf("binarySearch() gotIndex = %v, want %v", gotIndex, tt.wantIndex)
			}
			if gotError != nil && tt.wantError != nil && gotError.Error() != tt.wantError.Error() {
				t.Errorf("binarySearch() gotError = %v, want %v", gotError, tt.wantError)
			}
		})
	}
}
