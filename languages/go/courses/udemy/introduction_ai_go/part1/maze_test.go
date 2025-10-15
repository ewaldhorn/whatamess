package main

import (
	"mazes/point"
	"testing"
)

// ------------------------------------------------------------------------------------------------
func TestMazeInitialization(t *testing.T) {
	start := point.Point{Row: 0, Col: 0}
	end := point.Point{Row: 2, Col: 2}
	walls := [][]Wall{
		{
			{State: point.Point{Row: 0, Col: 0}, IsSolid: false},
			{State: point.Point{Row: 0, Col: 1}, IsSolid: true},
			{State: point.Point{Row: 0, Col: 2}, IsSolid: false},
		},
		{
			{State: point.Point{Row: 1, Col: 0}, IsSolid: false},
			{State: point.Point{Row: 1, Col: 1}, IsSolid: true},
			{State: point.Point{Row: 1, Col: 2}, IsSolid: false},
		},
		{
			{State: point.Point{Row: 2, Col: 0}, IsSolid: false},
			{State: point.Point{Row: 2, Col: 1}, IsSolid: false},
			{State: point.Point{Row: 2, Col: 2}, IsSolid: false},
		},
	}

	maze := Maze{
		Width:  3,
		Height: 3,
		Start:  start,
		End:    end,
		Walls:  walls,
	}

	if maze.Width != 3 {
		t.Errorf("Expected Width to be 3, got %d", maze.Width)
	}

	if maze.Height != 3 {
		t.Errorf("Expected Height to be 3, got %d", maze.Height)
	}

	if maze.Start != start {
		t.Errorf("Expected Start to be %+v, got %+v", start, maze.Start)
	}

	if maze.End != end {
		t.Errorf("Expected End to be %+v, got %+v", end, maze.End)
	}

	if len(maze.Walls) != 3 {
		t.Errorf("Expected Walls to have 3 rows, got %d", len(maze.Walls))
	}

	if len(maze.Walls[0]) != 3 {
		t.Errorf("Expected Walls to have 3 columns, got %d", len(maze.Walls[0]))
	}

	if maze.Walls[0][1].IsSolid != true {
		t.Errorf("Expected wall at (0,1) to be solid, but it was not")
	}

	if maze.Walls[1][1].IsSolid != true {
		t.Errorf("Expected wall at (1,1) to be solid, but it was not")
	}

	if maze.Walls[2][1].IsSolid != false {
		t.Errorf("Expected wall at (2,1) to not be solid, but it was")
	}
}
