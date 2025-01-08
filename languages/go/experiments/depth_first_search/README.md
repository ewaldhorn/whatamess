# DFS - Depth First Search

## Definition

Depth-First Search (DFS) is a fundamental algorithm in computer science and graph theory.
It's a traversal approach used to search and explore nodes or vertices in a graph or tree data structure.

### How DFS Works
The algorithm starts at a chosen node (called the root node) and explores as far as possible
along each branch before backtracking. It uses a stack data structure to keep track of nodes to visit next.

Here's a step-by-step breakdown:
- Choose a starting node: Select a node in the graph as the starting point.
- Mark the node as visited: Keep track of the node as visited to avoid revisiting it.
- Explore neighbors: Visit all unvisited neighbors of the current node.
- Backtrack: If all neighbors are visited, backtrack to the previous node.
- Repeat: Continue exploring nodes until all reachable nodes are visited.


### Types of DFS
There are three main types of DFS:
- Pre-order DFS: Visit the current node before exploring its neighbors.
- In-order DFS: Visit the current node between exploring its neighbors.
- Post-order DFS: Visit the current node after exploring its neighbors.

### Example Use Cases
DFS has numerous applications:
- Finding connected components: Identify connected subgraphs within a larger graph.
- Topological sorting: Order nodes in a directed acyclic graph (DAG) such that edges point from earlier nodes to later nodes.
- Finding strongly connected components: Identify subgraphs where every node is reachable from every other node.
- Solving mazes or puzzles: Explore possible paths to find a solution.

### Advantages and Disadvantages
__Advantages__:
- Simple to implement: DFS is a straightforward algorithm to understand and code.
- Efficient for sparse graphs: DFS performs well on graphs with few edges.

__Disadvantages__:
- Can get stuck in infinite loops: If the graph contains cycles, DFS may revisit nodes indefinitely.
- Not suitable for very large graphs: DFS can be slow for massive graphs due to the recursive function calls.
- In summary, Depth-First Search is a powerful algorithm for exploring graphs and trees. While it has its limitations, DFS is an essential technique in computer science and graph theory.
