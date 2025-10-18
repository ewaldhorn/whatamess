package main

import (
	"flag"
	"fmt"
	"mazes/algorithm"
	"mazes/globals"
	"mazes/helper"
	"mazes/imager"
	"mazes/maze"
	"os"
	"time"
)

// ------------------------------------------------------------------------------------------------
const (
	DFS      = iota // depth first search
	BFS             // breadth first search
	GBFS            // greedy best first search
	ASTAR           // a-start
	DIJKSTRA        // dijkstra
)

// ------------------------------------------------------------------------------------------------
func init() {
	_ = os.Mkdir(globals.TemporaryDirectory, os.ModePerm)
	helper.ClearTempDirectory()
}

// ------------------------------------------------------------------------------------------------
func main() {
	var m maze.Maze
	var maze, searchType string

	flag.StringVar(&maze, "file", "../mazes/maze.txt", "maze file")
	flag.StringVar(&searchType, "search", "dfs", "search type")
	flag.BoolVar(&m.Debug, "debug", false, "write debugging info")
	flag.BoolVar(&m.Animate, "animate", false, "produce animation")
	flag.Parse()

	err := m.Load(maze)
	if err != nil {
		fmt.Printf("error loading maze: %v", err)
		os.Exit(1)
	}

	start := time.Now()

	switch searchType {
	case "dfs":
		m.SearchType = DFS
		solveDFS(&m)
	default:
		fmt.Println("Invalid search type specified.")
		os.Exit(1)
	}

	if len(m.Solution.Actions) > 0 {
		fmt.Println("Solution:")
		fmt.Println("Solution is", len(m.Solution.Cells), "steps.")
		fmt.Println("Time to solve:", time.Since(start))
		fmt.Println()
		m.Print()
		imager.RenderMazeToDisk(&m, "image.png")
		fmt.Println()
	} else {
		fmt.Println("No solution found!")
	}

	fmt.Println("Explored", len(m.Explored), "nodes.")

	if m.Animate {
		fmt.Println("Building animation...")
		imager.RenderAnimatedMazeToDisk()
		fmt.Println("Done!")
	}
}

// ------------------------------------------------------------------------------------------------
func solveDFS(m *maze.Maze) {
	var s algorithm.DepthFirstSearch
	s.Game = m
	fmt.Println("Goal is", s.Game.End)
	s.Solve()
}
