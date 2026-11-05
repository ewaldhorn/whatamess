package wall

import (
	"testing"

	"mazes/point"
)

func TestWall_NewWall(t *testing.T) {
	tests := []struct {
		name     string
		state    point.Point
		isSolid  bool
		expected Wall
	}{
		{
			name:     "solid wall at origin",
			state:    point.Point{X: 0, Y: 0},
			isSolid:  true,
			expected: Wall{State: point.Point{X: 0, Y: 0}, IsSolid: true},
		},
		{
			name:     "non-solid wall at positive coordinates",
			state:    point.Point{X: 5, Y: 10},
			isSolid:  false,
			expected: Wall{State: point.Point{X: 5, Y: 10}, IsSolid: false},
		},
		{
			name:     "solid wall at negative coordinates",
			state:    point.Point{X: -3, Y: -7},
			isSolid:  true,
			expected: Wall{State: point.Point{X: -3, Y: -7}, IsSolid: true},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			wall := Wall{
				State:   tt.state,
				IsSolid: tt.isSolid,
			}

			if wall.State != tt.expected.State {
				t.Errorf("Expected state %v, got %v", tt.expected.State, wall.State)
			}

			if wall.IsSolid != tt.expected.IsSolid {
				t.Errorf("Expected IsSolid %v, got %v", tt.expected.IsSolid, wall.IsSolid)
			}
		})
	}
}

func TestWall_Equality(t *testing.T) {
	wall1 := Wall{State: point.Point{X: 1, Y: 2}, IsSolid: true}
	wall2 := Wall{State: point.Point{X: 1, Y: 2}, IsSolid: true}
	wall3 := Wall{State: point.Point{X: 1, Y: 2}, IsSolid: false} // different solidity
	wall4 := Wall{State: point.Point{X: 3, Y: 4}, IsSolid: true}  // different position

	if wall1 != wall2 {
		t.Error("Walls with same state and solidity should be equal")
	}

	if wall1 == wall3 {
		t.Error("Walls with different solidity should not be equal")
	}

	if wall1 == wall4 {
		t.Error("Walls with different positions should not be equal")
	}
}

func TestWall_ZeroValue(t *testing.T) {
	var zeroWall Wall

	if zeroWall.State != (point.Point{}) {
		t.Error("Zero value Wall should have zero value State")
	}

	if zeroWall.IsSolid != false {
		t.Error("Zero value Wall should have IsSolid false")
	}
}

func TestWall_Modification(t *testing.T) {
	wall := Wall{State: point.Point{X: 1, Y: 1}, IsSolid: true}

	// Test state modification
	newState := point.Point{X: 5, Y: 5}
	wall.State = newState
	if wall.State != newState {
		t.Errorf("Expected state %v after modification, got %v", newState, wall.State)
	}

	// Test IsSolid modification
	wall.IsSolid = false
	if wall.IsSolid != false {
		t.Error("Expected IsSolid false after modification")
	}
}
