package queues

import "mazes/node"

// ------------------------------------------------------------------------------------------------
type PriorityQueueDijkstra []*node.Node

// ------------------------------------------------------------------------------------------------
func (pq PriorityQueueDijkstra) Len() int {
	return len(pq)
}

// ------------------------------------------------------------------------------------------------
func (pq PriorityQueueDijkstra) Less(i, j int) bool {
	return pq[i].CostToGoal < pq[j].CostToGoal
}

// ------------------------------------------------------------------------------------------------
func (pq PriorityQueueDijkstra) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].Index = i
	pq[j].Index = j
}

// ------------------------------------------------------------------------------------------------
func (pq *PriorityQueueDijkstra) Push(x any) {
	n := len(*pq)
	item := x.(*node.Node)
	item.Index = n
	*pq = append(*pq, item)
}

// ------------------------------------------------------------------------------------------------
func (pq *PriorityQueueDijkstra) Pop() any {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil
	item.Index = -1
	*pq = old[0 : n-1]

	return item
}
