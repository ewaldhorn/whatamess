package maze

import (
	"mazes/point"
	"mazes/wall"
	"os"
	"path/filepath"
	"testing"
)

// ------------------------------------------------------------------------------------------------
func TestMazeInitialization(t *testing.T) {
	start := point.Point{Y: 0, X: 0}
	end := point.Point{Y: 2, X: 2}
	walls := [][]wall.Wall{
		{
			{State: point.Point{Y: 0, X: 0}, IsSolid: false},
			{State: point.Point{Y: 0, X: 1}, IsSolid: true},
			{State: point.Point{Y: 0, X: 2}, IsSolid: false},
		},
		{
			{State: point.Point{Y: 1, X: 0}, IsSolid: false},
			{State: point.Point{Y: 1, X: 1}, IsSolid: true},
			{State: point.Point{Y: 1, X: 2}, IsSolid: false},
		},
		{
			{State: point.Point{Y: 2, X: 0}, IsSolid: false},
			{State: point.Point{Y: 2, X: 1}, IsSolid: false},
			{State: point.Point{Y: 2, X: 2}, IsSolid: false},
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
		t.Errorf("Expected Walls to have 3 Ys, got %d", len(maze.Walls))
	}

	if len(maze.Walls[0]) != 3 {
		t.Errorf("Expected Walls to have 3 Xumns, got %d", len(maze.Walls[0]))
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

// ------------------------------------------------------------------------------------------------
func TestParseRawMaze(t *testing.T) {
	t.Run("valid maze", func(t *testing.T) {
		lines := []string{
			"A# ",
			" #B",
		}
		var m Maze
		err := m.parseRawMaze(lines)
		if err != nil {
			t.Fatalf("parseRawMaze failed: %v", err)
		}
		if m.Height != 2 {
			t.Errorf("Expected Height 2, got %d", m.Height)
		}
		if m.Width != 3 {
			t.Errorf("Expected Width 3, got %d", m.Width)
		}
		if len(m.Walls) != 2 {
			t.Errorf("Expected 2 wall rows, got %d", len(m.Walls))
		}
		if len(m.Walls[0]) != 3 {
			t.Errorf("Expected 3 walls in row 0, got %d", len(m.Walls[0]))
		}
		if (m.Start != point.Point{Y: 0, X: 0}) {
			t.Errorf("Incorrect start point: got %+v", m.Start)
		}
		if (m.End != point.Point{Y: 1, X: 2}) {
			t.Errorf("Incorrect end point: got %+v", m.End)
		}
	})

	t.Run("maze with inconsistent line length", func(t *testing.T) {
		lines := []string{
			"A#",
			" #B ",
		}
		var m Maze
		err := m.parseRawMaze(lines)
		if err != nil {
			t.Fatalf("parseRawMaze failed: %v", err)
		}
		if m.Height != 2 {
			t.Errorf("Expected Height 2, got %d", m.Height)
		}
		if m.Width != 4 {
			t.Errorf("Expected Width 4, got %d", m.Width)
		}
		if len(m.Walls[0]) != 2 {
			t.Errorf("Expected 2 walls in row 0, got %d", len(m.Walls[0]))
		}
		if len(m.Walls[1]) != 4 {
			t.Errorf("Expected 4 walls in row 1, got %d", len(m.Walls[1]))
		}
	})

	t.Run("maze with no start", func(t *testing.T) {
		lines := []string{" #", " B"}
		var m Maze
		err := m.parseRawMaze(lines)
		if err == nil {
			t.Error("Expected error for no start, got nil")
		}
	})

	t.Run("maze with no end", func(t *testing.T) {
		lines := []string{"A#", "  "}
		var m Maze
		err := m.parseRawMaze(lines)
		if err == nil {
			t.Error("Expected error for no end, got nil")
		}
	})
}

// ------------------------------------------------------------------------------------------------
func TestLoad(t *testing.T) {
	t.Run("valid maze file", func(t *testing.T) {
		content := "A#\n# B"
		dir := t.TempDir()
		tmpfn := filepath.Join(dir, "maze.txt")
		if err := os.WriteFile(tmpfn, []byte(content), 0666); err != nil {
			t.Fatal(err)
		}

		var m Maze
		err := m.Load(tmpfn)
		if err != nil {
			t.Fatalf("Load failed: %v", err)
		}

		if m.Height != 2 {
			t.Errorf("Expected Height 2, got %d", m.Height)
		}
		// First line is "A#\n", length 3. Width is based on this.
		if m.Width != 3 {
			t.Errorf("Expected Width 3, got %d", m.Width)
		}
		// The parser skips the newline, so row 0 has 2 walls.
		if len(m.Walls[0]) != 2 {
			t.Errorf("Expected 2 walls in row 0, got %d", len(m.Walls[0]))
		}
		// The second line has 3 characters.
		if len(m.Walls[1]) != 3 {
			t.Errorf("Expected 3 walls in row 1, got %d", len(m.Walls[1]))
		}
		if (m.Start != point.Point{Y: 0, X: 0}) {
			t.Errorf("Incorrect start point: got %+v", m.Start)
		}
		if (m.End != point.Point{Y: 1, X: 2}) {
			t.Errorf("Incorrect end point: got %+v", m.End)
		}
	})

	t.Run("file not found", func(t *testing.T) {
		var m Maze
		err := m.Load("nonexistent-file.txt")
		if err == nil {
			t.Error("Expected an error for a non-existent file, but got nil")
		}
	})

	t.Run("empty file", func(t *testing.T) {
		dir := t.TempDir()
		tmpfn := filepath.Join(dir, "empty.txt")
		if err := os.WriteFile(tmpfn, []byte(""), 0666); err != nil {
			t.Fatal(err)
		}

		var m Maze
		err := m.Load(tmpfn)
		if err == nil {
			t.Error("Expected an error for an empty file, but got nil")
		}
	})
}
