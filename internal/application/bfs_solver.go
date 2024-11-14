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
			path := []domain.Point{}
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

// findEntrance finds the first open cell on the left edge of the maze.
func (s *BFSSolver) findEntrance(maze *domain.Maze) domain.Point {
	for y := 0; y < maze.Height; y++ {
		if !maze.Grid[y][0].Wall {
			return domain.Point{X: 0, Y: y}
		}
	}
	return domain.Point{X: 0, Y: 1} // If not found, return a default value
}

// findExit finds the first open cell on the right edge of the maze.
func (s *BFSSolver) findExit(maze *domain.Maze) domain.Point {
	for y := 0; y < maze.Height; y++ {
		if !maze.Grid[y][maze.Width-1].Wall {
			return domain.Point{X: maze.Width - 1, Y: y}
		}
	}
	return domain.Point{X: maze.Width - 1, Y: maze.Height - 2} // If not found, return a default value
}
