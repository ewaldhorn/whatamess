package point

import "testing"

// ------------------------------------------------------------------------------------------------
func TestPoint(t *testing.T) {
	p := Point{X: 2, Y: 1}

	if p.Y != 1 {
		t.Errorf("Expected Y to be 1, but got %d", p.Y)
	}

	if p.X != 2 {
		t.Errorf("Expected X to be 2, but got %d", p.X)
	}

	p.Y = 3
	p.X = 4

	if p.Y != 3 {
		t.Errorf("Expected Y to be 3, but got %d", p.Y)
	}

	if p.X != 4 {
		t.Errorf("Expected X to be 4, but got %d", p.X)
	}
}
