package domain

// MazeGenerator описывает интерфейс для генерации лабиринта.
type MazeGenerator interface {
	Generate(maze *Maze)
}

// PathSolver описывает интерфейс для нахождения пути в лабиринте.
type PathSolver interface {
	Solve(maze *Maze, start, end Point) []Point
}
