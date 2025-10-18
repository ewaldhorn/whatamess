package helper

import (
	"mazes/point"
	"testing"
)

// ------------------------------------------------------------------------------------------------
func TestContainsPoint(t *testing.T) {
	// Test case 1: Point exists in the slice
	p1 := point.Point{X: 1, Y: 2}
	p2 := point.Point{X: 3, Y: 4}
	p3 := point.Point{X: 5, Y: 6}
	haystack := []point.Point{p1, p2, p3}
	needle := point.Point{X: 3, Y: 4}

	if !ContainsPoint(needle, haystack) {
		t.Error("Expected ContainsPoint to return true when point exists in slice")
	}

	// Test case 2: Point does not exist in the slice
	needle2 := point.Point{X: 7, Y: 8}
	if ContainsPoint(needle2, haystack) {
		t.Error("Expected ContainsPoint to return false when point does not exist in slice")
	}

	// Test case 3: Empty slice
	emptyHaystack := []point.Point{}
	if ContainsPoint(needle, emptyHaystack) {
		t.Error("Expected ContainsPoint to return false when haystack is empty")
	}

	// Test case 4: Exact same point reference
	if !ContainsPoint(p1, haystack) {
		t.Error("Expected ContainsPoint to return true when checking exact same point reference")
	}
}