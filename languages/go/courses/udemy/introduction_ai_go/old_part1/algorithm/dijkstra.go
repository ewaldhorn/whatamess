package algorithm

import (
	"container/heap"
	"fmt"
	"log"
	"math/rand"
	"mazes/globals"
	"mazes/helper"
	"mazes/imager"
	"mazes/maze"
	"mazes/node"
	"mazes/point"
	"mazes/queues"
	"mazes/solution"
	"slices"
)

// ------------------------------------------------------------------------------------------------
type Dijkstra struct {
	Frontier queues.PriorityQueueDijkstra
	Game     *maze.Maze
}

// ------------------------------------------------------------------------------------------------
func (d *Dijkstra) GetFrontier() []*node.Node {
	return d.Frontier
}

// ------------------------------------------------------------------------------------------------
func (d *Dijkstra) AddNode(i *node.Node) {
	i.CostToGoal = i.ManhattanDistance(d.Game.Start)
	d.Frontier.Push(i)
	heap.Init(&d.Frontier)
}

// ------------------------------------------------------------------------------------------------
func (d *Dijkstra) ContainsState(i *node.Node) bool {
	for _, x := range d.Frontier {
		if x.State == i.State {
			return true
		}
	}
	return false
}

// ------------------------------------------------------------------------------------------------
func (d *Dijkstra) IsEmpty() bool {
	return len(d.Frontier) == 0
}

// ------------------------------------------------------------------------------------------------
func (d *Dijkstra) Remove() (*node.Node, error) {
	if len(d.Frontier) > 0 {
		if d.Game.Debug {
			fmt.Println("Frontier before remove:")
			for _, f := range d.Frontier {
				fmt.Println("Node:", f.State)
			}
		}

		return heap.Pop(&d.Frontier).(*node.Node), nil
	}

	return nil, fmt.Errorf("frontier is empty")
}

// ------------------------------------------------------------------------------------------------
func (d *Dijkstra) Neighbours(n *node.Node) []*node.Node {
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
func (d *Dijkstra) Solve() {
	fmt.Println("Starting to solve maze using Dijkstra Search...")

	d.Game.NumExplored = 0
	start := node.Node{
		State: d.Game.Start,
	}

	d.AddNode(&start)
	d.Game.CurrentNode = &start

	for {
		if d.IsEmpty() {
			return
		}

		currentNode, err := d.Remove()
		if err != nil {
			log.Println(err)
			return
		}

		if d.Game.Debug {
			fmt.Println("Removed", currentNode.State)
			fmt.Println("-------")
			fmt.Println()
		}

		d.Game.CurrentNode = currentNode
		d.Game.NumExplored += 1

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

			d.Game.Solution = solution.Solution{
				Actions: actions,
				Cells:   cells,
			}
			d.Game.Explored = append(d.Game.Explored, currentNode.State)
			break
		}

		d.Game.Explored = append(d.Game.Explored, currentNode.State)

		if d.Game.Animate {
			imager.RenderMazeToDisk(d.Game, fmt.Sprintf("%s%06d.png", globals.TemporaryDirectory, d.Game.NumExplored))
		}

		for _, x := range d.Neighbours(currentNode) {
			if !d.ContainsState(x) {
				if !helper.ContainsPoint(x.State, d.Game.Explored) {
					d.AddNode(&node.Node{
						State:  x.State,
						Parent: currentNode,
						Action: x.Action,
					})
				}
			}
		}
	}
}
