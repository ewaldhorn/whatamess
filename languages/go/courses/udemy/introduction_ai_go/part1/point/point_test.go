package point

import "testing"

// ------------------------------------------------------------------------------------------------
func TestPoint(t *testing.T) {
	p := Point{Row: 1, Col: 2}

	if p.Row != 1 {
		t.Errorf("Expected Row to be 1, but got %d", p.Row)
	}

	if p.Col != 2 {
		t.Errorf("Expected Col to be 2, but got %d", p.Col)
	}

	p.Row = 3
	p.Col = 4

	if p.Row != 3 {
		t.Errorf("Expected Row to be 3, but got %d", p.Row)
	}

	if p.Col != 4 {
		t.Errorf("Expected Col to be 4, but got %d", p.Col)
	}
}
