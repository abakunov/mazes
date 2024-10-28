package application

import (
	"math/rand"

	"github.com/abakunov/mazes/internal/domain"
)

// RECURSIVE BACKTRACKING

type DFSGenerator struct{}

func shuffleDirections() []int {
	directions := []int{0, 1, 2, 3}

	rand.Shuffle(4, func(i, j int) { directions[i], directions[j] = directions[j], directions[i] })

	return directions
}

func (g *DFSGenerator) Generate(maze *domain.Maze, start domain.Point) {
	x, y := start.X, start.Y
	maze.Grid[y][x].Visited = true    // Отмечаем текущую клетку как посещённую
	directions := shuffleDirections() // Перемешиваем направления

	for _, dir := range directions {
		dx, dy := 0, 0

		switch dir {
		case 0: // Вверх
			dy = -2
		case 1: // Вправо
			dx = 2
		case 2: // Вниз
			dy = 2
		case 3: // Влево
			dx = -2
		}

		nx, ny := x+dx, y+dy
		// Проверяем границы и посещённость клетки
		if ny >= 0 && ny < maze.Height && nx >= 0 && nx < maze.Width && !maze.Grid[ny][nx].Visited {
			maze.Grid[ny][nx].Visited = true             // Отмечаем новую клетку как посещённую
			maze.Grid[y+dy/2][x+dx/2].Wall = false       // Убираем стену между клетками
			g.Generate(maze, domain.Point{X: nx, Y: ny}) // Рекурсивно генерируем лабиринт
		}
	}
}
