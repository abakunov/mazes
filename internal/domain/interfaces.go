package domain

// Generator describes the interface for maze generation.
type Generator interface {
	Generate(maze *Maze, entryPoint, exitPoint Point)
}

// Solver describes the interface for finding a path in the maze.
type Solver interface {
	FindPath(maze *Maze, entryPoint, exitPoint Point) []Point
}
