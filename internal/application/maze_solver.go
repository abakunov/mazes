package application

import "github.com/abakunov/mazes/internal/domain"

type BFSSolver struct{}

func (s *BFSSolver) FindPath(maze *domain.Maze, start, end domain.Point) []domain.Point {
	// Очередь для BFS
	queue := []domain.Point{start}
	visited := make(map[domain.Point]bool)        // Мапа посещённых точек
	parent := make(map[domain.Point]domain.Point) // Мапа родителей для восстановления пути

	visited[start] = true // Помечаем вход как посещённый

	// Направления: вниз, вправо, вверх, влево
	directions := []domain.Point{{X: 0, Y: 1}, {X: 1, Y: 0}, {X: 0, Y: -1}, {X: -1, Y: 0}}

	// BFS
	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]

		// Если точка на границе и это не вход, считаем её выходом
		if current == end || s.isExit(current, maze) {
			// Восстанавливаем путь
			var path []domain.Point
			for p := current; p != start; p = parent[p] {
				path = append([]domain.Point{p}, path...)
			}

			path = append([]domain.Point{start}, path...) // Добавляем вход в путь

			return path
		}

		// Проходим по всем соседям
		for _, dir := range directions {
			nx, ny := current.X+dir.X, current.Y+dir.Y
			next := domain.Point{X: nx, Y: ny}

			// Проверяем, что следующая точка в пределах лабиринта и не стена
			if nx >= 0 && nx < maze.Width && ny >= 0 && ny < maze.Height && !maze.Grid[ny][nx].Wall && !visited[next] {
				queue = append(queue, next)
				visited[next] = true
				parent[next] = current
			}
		}
	}

	return nil // Путь не найден
}

// isExit проверяет, находится ли точка на границе лабиринта и является ли она выходом.
func (s *BFSSolver) isExit(p domain.Point, maze *domain.Maze) bool {
	return p.X == maze.Width-1 && p.Y == maze.Height-2
}
