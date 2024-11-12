package application

import "github.com/abakunov/mazes/internal/domain"

type BFSSolver struct{}

func (s *BFSSolver) FindPath(maze *domain.Maze, entry, exit domain.Point) []domain.Point {
	// Инициализация очереди для BFS
	queue := []domain.Point{entry}
	visited := make(map[domain.Point]bool)
	parent := make(map[domain.Point]domain.Point)

	// Помечаем точку входа как посещённую
	visited[entry] = true

	// Направления для движения: вверх, вправо, вниз, влево
	directions := []domain.Point{
		{X: 0, Y: -1}, // Вверх
		{X: 1, Y: 0},  // Вправо
		{X: 0, Y: 1},  // Вниз
		{X: -1, Y: 0}, // Влево
	}

	// BFS-поиск пути
	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]

		// Если достигли точки выхода, восстанавливаем путь
		if current == exit {
			path := []domain.Point{}
			for p := current; p != entry; p = parent[p] {
				path = append([]domain.Point{p}, path...)
			}
			return append([]domain.Point{entry}, path...)
		}

		// Проходим по всем соседям
		for _, dir := range directions {
			neighbor := domain.Point{X: current.X + dir.X, Y: current.Y + dir.Y}

			// Проверяем, что соседняя точка находится в пределах лабиринта и является проходом (не стеной)
			if neighbor.X >= 0 && neighbor.X < maze.Width && neighbor.Y >= 0 && neighbor.Y < maze.Height &&
				!maze.Grid[neighbor.Y][neighbor.X].Wall && !visited[neighbor] {
				queue = append(queue, neighbor)
				visited[neighbor] = true
				parent[neighbor] = current
			}
		}
	}

	// Путь не найден
	return nil
}

// findEntrance находит первую открытую клетку на левой границе лабиринта.
func (s *BFSSolver) findEntrance(maze *domain.Maze) domain.Point {
	for y := 0; y < maze.Height; y++ {
		if !maze.Grid[y][0].Wall {
			return domain.Point{X: 0, Y: y}
		}
	}
	return domain.Point{X: 0, Y: 1} // Если не найдена, возвращаем дефолтное значение
}

// findExit находит первую открытую клетку на правой границе лабиринта.
func (s *BFSSolver) findExit(maze *domain.Maze) domain.Point {
	for y := 0; y < maze.Height; y++ {
		if !maze.Grid[y][maze.Width-1].Wall {
			return domain.Point{X: maze.Width - 1, Y: y}
		}
	}
	return domain.Point{X: maze.Width - 1, Y: maze.Height - 2} // Если не найдена, возвращаем дефолтное значение
}
