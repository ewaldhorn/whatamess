package main

import "testing"

func Test_isInRect(t *testing.T) {
	tests := []struct {
		name      string
		rectangle Rect
		xPos      int
		yPos      int
		expected  bool
	}{
		{name: "2x2a", rectangle: Rect{topLeftX: 1, topLeftY: 1, width: 1, height: 1}, xPos: 1, yPos: 1, expected: true},
		{name: "2x2b", rectangle: Rect{topLeftX: 1, topLeftY: 1, width: 1, height: 1}, xPos: 2, yPos: 1, expected: true},
		{name: "2x2c", rectangle: Rect{topLeftX: 1, topLeftY: 1, width: 1, height: 1}, xPos: 2, yPos: 2, expected: true},
		{name: "2x2d", rectangle: Rect{topLeftX: 1, topLeftY: 1, width: 1, height: 1}, xPos: 3, yPos: 1, expected: false},
	}

	for _, test := range tests {
		result := test.rectangle.isInRect(test.xPos, test.yPos)

		if result != test.expected {
			t.Fatalf("[%s] expected %v, got %v", test.name, test.expected, result)
		}
	}
}
