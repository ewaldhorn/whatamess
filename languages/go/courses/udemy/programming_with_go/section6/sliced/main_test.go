package main

import "testing"

// --------------------------------------------------------------------- Test_sliceIterationWithFor
func Test_sliceIterationWithFor(t *testing.T) {
	mySlice := []string{"any", "string", "will", "do"}

	for i := 0; i < len(mySlice); i++ {
		if mySlice[i] == "" {
			t.Fatalf("unexpected empty string at %d\n", i)
		}
	}
}

// ------------------------------------------------------------------- Test_sliceIterationWithRange
func Test_sliceIterationWithRange(t *testing.T) {
	mySlice := []int{1, 2, 3, 4, 5, 6}

	for pos, el := range mySlice {
		if el == 0 {
			t.Fatalf("unexpected empty string at %d\n", pos)
		}
	}
}

// --------------------------------------------------------------------------- Test_removeFromSlice
func Test_removeFromSlice(t *testing.T) {
	mySlice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	removeIndex := 4

	mySlice = append(mySlice[:removeIndex], mySlice[removeIndex+1:]...)

	expectedSlice := []int{1, 2, 3, 4, 6, 7, 8, 9, 10}

	for i := range len(mySlice) {
		if mySlice[i] != expectedSlice[i] {
			t.Fatalf("unexpected error %d != %d at pos %d\n", mySlice[i], expectedSlice[i], i)
		}
	}

}
