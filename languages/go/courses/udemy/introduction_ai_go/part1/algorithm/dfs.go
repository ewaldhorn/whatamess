package algorithm

import (
	"fmt"
	"mazes/maze"
	"mazes/node"
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
