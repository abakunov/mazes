package application

import "github.com/abakunov/mazes/internal/domain"

type BFSSolver struct{}

func (s *BFSSolver) FindPath(maze *domain.Maze, entry, exit domain.Point) []domain.Point {
	// Initialize queue for BFS
	queue := []domain.Point{entry}
	visited := make(map[domain.Point]bool)
	parent := make(map[domain.Point]domain.Point)

	// Mark the entry point as visited
	visited[entry] = true

	// Movement directions: up, right, down, left
	directions := []domain.Point{
		{X: 0, Y: -1}, // Up
		{X: 1, Y: 0},  // Right
		{X: 0, Y: 1},  // Down
		{X: -1, Y: 0}, // Left
	}

	// BFS pathfinding
	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]

		// If the exit point is reached, reconstruct the path
		if current == exit {
			var path []domain.Point
			for p := current; p != entry; p = parent[p] {
				path = append([]domain.Point{p}, path...)
			}

			return append([]domain.Point{entry}, path...)
		}

		// Iterate over all neighbors
		for _, dir := range directions {
			neighbor := domain.Point{X: current.X + dir.X, Y: current.Y + dir.Y}

			// Check if the neighbor is within maze bounds and is a passage (not a wall)
			if neighbor.X >= 0 && neighbor.X < maze.Width && neighbor.Y >= 0 && neighbor.Y < maze.Height &&
				!maze.Grid[neighbor.Y][neighbor.X].Wall && !visited[neighbor] {
				queue = append(queue, neighbor)
				visited[neighbor] = true
				parent[neighbor] = current
			}
		}
	}

	// Path not found
	return nil
}
