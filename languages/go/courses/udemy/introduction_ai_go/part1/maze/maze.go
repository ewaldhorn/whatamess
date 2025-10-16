package maze

import (
	"bufio"
	"fmt"
	"mazes/node"
	"mazes/point"
	"mazes/solution"
	"mazes/wall"
	"os"
	"strings"
)

// ------------------------------------------------------------------------------------------------
type Maze struct {
	Width, Height      int
	Start, End         point.Point
	Walls              [][]wall.Wall
	CurrentNode        *node.Node
	Solution           solution.Solution
	Explored           []point.Point
	Steps, NumExplored int
	Debug              bool
	SearchType         int
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
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		return fmt.Errorf("cannot read from file `%s`: %s", filename, err)
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

	m.Height = len(lines)
	m.Width = 0
	if m.Height > 0 {
		maxWidth := 0
		for _, l := range lines {
			if len(l) > maxWidth {
				maxWidth = len(l)
			}
		}
		m.Width = maxWidth
	}

	var rows [][]wall.Wall

	for i, row := range lines {
		var cols []wall.Wall
		for j, col := range row {
			curLetter := fmt.Sprintf("%c", col)
			var wall wall.Wall
			switch curLetter {
			case "A":
				m.Start = point.Point{Y: i, X: j}
				wall.State.Y = i
				wall.State.X = j
				wall.IsSolid = false
			case "B":
				m.End = point.Point{Y: i, X: j}
				wall.State.Y = i
				wall.State.X = j
				wall.IsSolid = false
			case " ":
				wall.State.Y = i
				wall.State.X = j
				wall.IsSolid = false
			case "#":
				wall.State.Y = i
				wall.State.X = j
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

// ------------------------------------------------------------------------------------------------
func (m *Maze) Print() {
	for y, row := range m.Walls {
		for x, col := range row {
			if col.IsSolid {
				fmt.Println("█")
			} else if
		}
	}
}
