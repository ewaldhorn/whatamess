package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

// ------------------------------------------------------------------------------------------------
type Maze struct {
	Width, Height int
	Start, End    Point
	Walls         [][]Wall
}

// ------------------------------------------------------------------------------------------------
func (m *Maze) Load(filename string) error {
	f, err := os.Open(filename)
	if err != nil {
		fmt.Printf("error opening `%s`: %v\n", filename, err)
		return err
	}
	defer f.Close()

	var lines []string

	reader := bufio.NewReader(f)
	for {
		line, err := reader.ReadString('\n')
		if err == io.EOF {
			break
		} else if err != nil {
			return fmt.Errorf("cannot read from file `%s`: %s", filename, err)
		}

		lines = append(lines, line)
	}

	return m.parseRawMaze(lines)
}

// ------------------------------------------------------------------------------------------------
func (m *Maze) parseRawMaze(lines []string) error {
	foundStart, foundEnd := false, false
	for _, line := range lines {
		if strings.Contains(line, "A") {
			foundStart = true
		}

		if strings.Contains(line, "B") {
			foundEnd = true
		}
	}

	if !foundStart {
		return fmt.Errorf("unable to find Start")
	}

	if !foundEnd {
		return fmt.Errorf("unable to find End")
	}

	// set maze width and height
	m.Height = len(lines)
	m.Width = len(lines[0])

	var rows [][]Wall

	for i, row := range lines {
		var cols []Wall
		for j, col := range row {
			curLetter := fmt.Sprintf("%c", col)
			var wall Wall
			switch curLetter {
			case "A":
				m.Start = Point{Row: i, Col: j}
				wall.State.Row = i
				wall.State.Col = j
				wall.IsSolid = false
			case "B":
				m.End = Point{Row: i, Col: j}
				wall.State.Row = i
				wall.State.Col = j
				wall.IsSolid = false
			case " ":
				wall.State.Row = i
				wall.State.Col = j
				wall.IsSolid = false
			case "#":
				wall.State.Row = i
				wall.State.Col = j
				wall.IsSolid = true
			default:
				continue
			}
			cols = append(cols, wall)
		}
		rows = append(rows, cols)
	}

	m.Walls = rows
	return nil
}
