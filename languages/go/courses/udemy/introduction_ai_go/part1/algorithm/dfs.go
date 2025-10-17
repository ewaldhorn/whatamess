package algorithm

import (
	"fmt"
	"math/rand"
	"mazes/helper"
	"mazes/maze"
	"mazes/node"
	"mazes/point"
	"mazes/solution"
	"slices"
)

// ------------------------------------------------------------------------------------------------
type DepthFirstSearch struct {
	Frontier []*node.Node
	Game     *maze.Maze
}

// ------------------------------------------------------------------------------------------------
func (d *DepthFirstSearch) GetFrontier() []*node.Node {
	return d.Frontier
}

// ------------------------------------------------------------------------------------------------
func (d *DepthFirstSearch) Add(n *node.Node) {
	d.Frontier = append(d.Frontier, n)
}

// ------------------------------------------------------------------------------------------------
func (d *DepthFirstSearch) ContainsState(n *node.Node) bool {
	for _, x := range d.Frontier {
		if x.State == n.State {
			return true
		}
	}
	return false
}

// ------------------------------------------------------------------------------------------------
func (d *DepthFirstSearch) IsEmpty() bool {
	return len(d.Frontier) == 0
}

// ------------------------------------------------------------------------------------------------
func (d *DepthFirstSearch) Remove() (*node.Node, error) {
	if len(d.Frontier) > 0 {
		if d.Game.Debug {
			fmt.Println("Frontier before remove:")
			for _, x := range d.Frontier {
				fmt.Println("Node:", x.State)
			}
		}

		node := d.Frontier[len(d.Frontier)-1]
		d.Frontier = d.Frontier[:len(d.Frontier)-1]
		return node, nil
	}

	return nil, fmt.Errorf("Frontier is empty, can't remove anything!")
}

// ------------------------------------------------------------------------------------------------
func (d *DepthFirstSearch) Neighbours(n *node.Node) []*node.Node {
	y := n.State.Y
	x := n.State.X

	candidates := []*node.Node{
		{State: point.Point{Y: y - 1, X: x}, Parent: n, Action: "up"},
		{State: point.Point{Y: y + 1, X: x}, Parent: n, Action: "down"},
		{State: point.Point{Y: y, X: x - 1}, Parent: n, Action: "left"},
		{State: point.Point{Y: y, X: x + 1}, Parent: n, Action: "right"},
	}

	var neighbours []*node.Node

	for _, c := range candidates {
		if 0 <= c.State.Y && c.State.Y < d.Game.Height {
			if 0 <= c.State.X && c.State.X < d.Game.Width {
				if !d.Game.Walls[c.State.Y][c.State.X].IsSolid {
					neighbours = append(neighbours, c)
				}
			}
		}
	}

	// randomise order
	for i := range neighbours {
		j := rand.Intn(i + 1)
		neighbours[i], neighbours[j] = neighbours[j], neighbours[i]
	}

	return neighbours
}

// ------------------------------------------------------------------------------------------------
func (d *DepthFirstSearch) Solve() {
	fmt.Println("Starting to solve maze using Depth First Search...")

	d.Game.NumExplored = 0

	start := node.Node{State: d.Game.Start, Parent: nil, Action: ""}

	d.Add(&start)
	d.Game.CurrentNode = &start

	for {
		if d.IsEmpty() {
			return
		}

		currentNode, err := d.Remove()
		if err != nil {
			fmt.Println("Error solving maze:", err)
			return
		}

		if d.Game.Debug {
			fmt.Println("Removed", currentNode.State)
			fmt.Println("-------")
			fmt.Println()
		}

		d.Game.CurrentNode = currentNode
		d.Game.NumExplored += 1

		// check for solution found
		if d.Game.End == currentNode.State {
			var actions []string
			var cells []point.Point
			for {
				if currentNode.Parent != nil {
					actions = append(actions, currentNode.Action)
					cells = append(cells, currentNode.State)
					currentNode = currentNode.Parent
				} else {
					break
				}
			}

			slices.Reverse(actions)
			slices.Reverse(cells)

			d.Game.Solution = solution.Solution{Actions: actions, Cells: cells}
			d.Game.Explored = append(d.Game.Explored, currentNode.State)
			break
		}

		d.Game.Explored = append(d.Game.Explored, currentNode.State)

		for _, n := range d.Neighbours(currentNode) {
			if !d.ContainsState(n) {
				if !helper.ContainsPoint(n.State, d.Game.Explored) {
					d.Add(&node.Node{State: n.State, Parent: currentNode, Action: n.Action})
				}
			}
		}
	}
}
