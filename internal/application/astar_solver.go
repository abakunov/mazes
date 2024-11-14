package application

import (
	"container/heap"
	"math"

	"github.com/abakunov/mazes/internal/domain"
)

// Node represents a node in A*.
type Node struct {
	Point    domain.Point
	Cost     int
	Priority int
	Index    int
	Parent   *Node
}

// PriorityQueue implements a priority queue for the A* algorithm.
type PriorityQueue []*Node

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].Priority < pq[j].Priority
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].Index = i
	pq[j].Index = j
}

func (pq *PriorityQueue) Push(x interface{}) {
	n := x.(*Node)
	n.Index = len(*pq)
	*pq = append(*pq, n)
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	node := old[n-1]
	old[n-1] = nil  // avoid memory leak
	node.Index = -1 // for safety
	*pq = old[0 : n-1]

	return node
}

type AStarSolver struct{}

func (s *AStarSolver) FindPath(maze *domain.Maze, entry, exit domain.Point) []domain.Point {
	// Initialize priority queue
	pq := &PriorityQueue{}
	heap.Init(pq)

	// Create the start node
	startNode := &Node{
		Point:    entry,
		Cost:     0,
		Priority: s.heuristic(entry, exit),
		Parent:   nil,
	}
	heap.Push(pq, startNode)

	// Store visited nodes
	visited := make(map[domain.Point]bool)
	visited[entry] = true

	// Movement directions: up, right, down, left
	directions := []domain.Point{
		{X: 0, Y: -1}, // Up
		{X: 1, Y: 0},  // Right
		{X: 0, Y: 1},  // Down
		{X: -1, Y: 0}, // Left
	}

	// A* pathfinding
	for pq.Len() > 0 {
		// Extract the node with the lowest priority
		currentNode := heap.Pop(pq).(*Node)
		currentPoint := currentNode.Point

		// If exit point is reached, reconstruct the path
		if currentPoint == exit {
			path := []domain.Point{}
			for node := currentNode; node != nil; node = node.Parent {
				path = append([]domain.Point{node.Point}, path...)
			}

			return path
		}

		// Iterate over all neighbors
		for _, dir := range directions {
			neighborPoint := domain.Point{X: currentPoint.X + dir.X, Y: currentPoint.Y + dir.Y}

			// Check if the neighbor is within the maze bounds and is a passage (not a wall)
			if neighborPoint.X >= 0 && neighborPoint.X < maze.Width && neighborPoint.Y >= 0 && neighborPoint.Y < maze.Height &&
				!maze.Grid[neighborPoint.Y][neighborPoint.X].Wall && !visited[neighborPoint] {
				// Calculate movement cost and priority
				newCost := currentNode.Cost + 1
				priority := newCost + s.heuristic(neighborPoint, exit)

				// Create a node for the neighbor
				neighborNode := &Node{
					Point:    neighborPoint,
					Cost:     newCost,
					Priority: priority,
					Parent:   currentNode,
				}

				// Add neighbor to the priority queue and mark as visited
				heap.Push(pq, neighborNode)

				visited[neighborPoint] = true
			}
		}
	}

	// Path not found
	return nil
}

// heuristic calculates the Manhattan heuristic (distance from the current point to the exit point).
func (s *AStarSolver) heuristic(a, b domain.Point) int {
	return int(math.Abs(float64(a.X-b.X)) + math.Abs(float64(a.Y-b.Y)))
}
