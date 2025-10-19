package algorithm

import (
	"fmt"
	"log"
	"math/rand"
	"mazes/globals"
	"mazes/helper"
	"mazes/imager"
	"mazes/maze"
	"mazes/node"
	"mazes/point"
	"mazes/solution"
	"slices"
)

// ------------------------------------------------------------------------------------------------
type BreadthFirstSearch struct {
	Frontier []*node.Node
	Game     *maze.Maze
}

// ------------------------------------------------------------------------------------------------
func (b *BreadthFirstSearch) GetFrontier() []*node.Node {
	return b.Frontier
}

// ------------------------------------------------------------------------------------------------
func (b *BreadthFirstSearch) AddNode(i *node.Node) {
	b.Frontier = append(b.Frontier, i)
}

// ------------------------------------------------------------------------------------------------
func (b *BreadthFirstSearch) ContainsState(i *node.Node) bool {
	for _, x := range b.Frontier {
		if x.State == i.State {
			return true
		}
	}
	return false
}

// ------------------------------------------------------------------------------------------------
func (b *BreadthFirstSearch) IsEmpty() bool {
	return len(b.Frontier) == 0
}

// ------------------------------------------------------------------------------------------------
func (b *BreadthFirstSearch) Remove() (*node.Node, error) {
	if len(b.Frontier) > 0 {
		if b.Game.Debug {
			fmt.Println("Frontier before remove:")
			for _, f := range b.Frontier {
				fmt.Println("Node:", f.State)
			}
		}

		node := b.Frontier[0]
		b.Frontier = b.Frontier[1:]
		return node, nil
	}

	return nil, fmt.Errorf("frontier is empty")
}

// ------------------------------------------------------------------------------------------------
func (b *BreadthFirstSearch) Neighbours(n *node.Node) []*node.Node {
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
		if 0 <= c.State.Y && c.State.Y < b.Game.Height {
			if 0 <= c.State.X && c.State.X < b.Game.Width {
				if !b.Game.Walls[c.State.Y][c.State.X].IsSolid {
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
func (b *BreadthFirstSearch) Solve() {
	fmt.Println("Starting to solve maze using Breadth First Search...")

	b.Game.NumExplored = 0
	start := node.Node{
		State:  b.Game.Start,
		Parent: nil,
		Action: "",
	}

	b.AddNode(&start)
	b.Game.CurrentNode = &start

	for {
		if b.IsEmpty() {
			return
		}

		currentNode, err := b.Remove()
		if err != nil {
			log.Println(err)
			return
		}

		if b.Game.Debug {
			fmt.Println("Removed", currentNode.State)
			fmt.Println("-------")
			fmt.Println()
		}

		b.Game.CurrentNode = currentNode
		b.Game.NumExplored += 1

		if b.Game.End == currentNode.State {
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

			b.Game.Solution = solution.Solution{
				Actions: actions,
				Cells:   cells,
			}
			b.Game.Explored = append(b.Game.Explored, currentNode.State)
			break
		}

		b.Game.Explored = append(b.Game.Explored, currentNode.State)

		if b.Game.Animate {
			imager.RenderMazeToDisk(b.Game, fmt.Sprintf("%s%06d.png", globals.TemporaryDirectory, b.Game.NumExplored))
		}

		for _, x := range b.Neighbours(currentNode) {
			if !b.ContainsState(x) {
				if !helper.ContainsPoint(x.State, b.Game.Explored) {
					b.AddNode(&node.Node{
						State:  x.State,
						Parent: currentNode,
						Action: x.Action,
					})
				}
			}
		}
	}
}
