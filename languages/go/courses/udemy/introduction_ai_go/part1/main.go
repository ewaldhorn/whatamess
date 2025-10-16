package main

import (
	"flag"
	"fmt"
	"mazes/maze"
	"os"
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
func main() {
	var m maze.Maze
	var maze, searchType string

	flag.StringVar(&maze, "file", "../mazes/maze.txt", "maze file")
	flag.StringVar(&searchType, "search", "dfs", "search type")
	flag.Parse()

	err := m.Load(maze)
	if err != nil {
		fmt.Printf("error loading maze: %v", err)
		os.Exit(1)
	}

	fmt.Printf("The maze is %dx%d big.\n", m.Height, m.Width)
}
