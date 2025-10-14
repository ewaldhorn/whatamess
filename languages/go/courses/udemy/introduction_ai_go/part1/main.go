package main

import (
	"fmt"
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
type Point struct {
	Row, Col int
}

// ------------------------------------------------------------------------------------------------
type Wall struct {
	State   Point
	IsSolid bool
}

// ------------------------------------------------------------------------------------------------
func main() {
	fmt.Println("Hello, World!")
}
