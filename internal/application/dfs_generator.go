package application

import (
	"github.com/abakunov/mazes/internal/domain"
	"math/rand"
	"time"
)

type DFSGenerator struct{}

// Generate создает лабиринт с использованием алгоритма DFS (поиск в глубину).
func (p *DFSGenerator) Generate(maze *domain.Maze, entryPoint, exitPoint domain.Point) {
	rand.Seed(time.Now().UnixNano())

	// Инициализация всех клеток как стены
	for y := 0; y < maze.Height; y++ {
		for x := 0; x < maze.Width; x++ {
			maze.Grid[y][x] = domain.Cell{Wall: true, Visited: false}
		}
	}

	// Установка внешних границ как стен, кроме точек входа и выхода
	p.setOuterWalls(maze, entryPoint, exitPoint)

	// Начинаем генерацию с точки входа, исключая внешние границы
	stack := []domain.Point{entryPoint}
	maze.Grid[entryPoint.Y][entryPoint.X].Visited = true
	maze.Grid[entryPoint.Y][entryPoint.X].Wall = false

	// Поиск в глубину
	for len(stack) > 0 {
		// Берём последнюю точку из стека
		current := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		// Получаем список непосещённых соседей
		neighbors := p.getUnvisitedNeighbors(maze, current)
		if len(neighbors) > 0 {
			// Добавляем текущую точку обратно в стек
			stack = append(stack, current)

			// Выбираем случайного непосещённого соседа
			next := neighbors[rand.Intn(len(neighbors))]

			// Убираем стену между текущей клеткой и выбранным соседом
			p.removeWallBetween(maze, current, next)

			// Помечаем соседа как посещённого и добавляем его в стек
			maze.Grid[next.Y][next.X].Visited = true
			maze.Grid[next.Y][next.X].Wall = false
			stack = append(stack, next)
		}
	}

	// Соединяем точку выхода с лабиринтом, если она изолирована
	p.connectExitPoint(maze, exitPoint)
}

// setOuterWalls устанавливает внешние границы как стены, оставляя проходы в точках входа и выхода
func (p *DFSGenerator) setOuterWalls(maze *domain.Maze, entryPoint, exitPoint domain.Point) {
	for x := 0; x < maze.Width; x++ {
		maze.Grid[0][x].Wall = true
		maze.Grid[maze.Height-1][x].Wall = true
	}
	for y := 0; y < maze.Height; y++ {
		maze.Grid[y][0].Wall = true
		maze.Grid[y][maze.Width-1].Wall = true
	}

	// Оставляем проходы для входа и выхода
	maze.Grid[entryPoint.Y][entryPoint.X].Wall = false
	maze.Grid[exitPoint.Y][exitPoint.X].Wall = false
}

// connectExitPoint соединяет точку выхода с лабиринтом, если она изолирована
func (p *DFSGenerator) connectExitPoint(maze *domain.Maze, exitPoint domain.Point) {
	// Находим соседей точки выхода
	neighbors := p.getUnvisitedNeighbors(maze, exitPoint)
	if len(neighbors) > 0 {
		// Выбираем случайного соседа
		next := neighbors[rand.Intn(len(neighbors))]

		// Убираем стену между точкой выхода и выбранным соседом
		p.removeWallBetween(maze, exitPoint, next)

		// Помечаем соседа как посещённого
		maze.Grid[next.Y][next.X].Visited = true
		maze.Grid[next.Y][next.X].Wall = false
	}
}

// getUnvisitedNeighbors возвращает список непосещённых соседей для указанной клетки
func (p *DFSGenerator) getUnvisitedNeighbors(maze *domain.Maze, cell domain.Point) []domain.Point {
	var neighbors []domain.Point
	directions := []struct{ x, y int }{{2, 0}, {-2, 0}, {0, 2}, {0, -2}}

	for _, d := range directions {
		nx, ny := cell.X+d.x, cell.Y+d.y
		if nx > 0 && nx < maze.Width-1 && ny > 0 && ny < maze.Height-1 && !maze.Grid[ny][nx].Visited {
			neighbors = append(neighbors, domain.Point{X: nx, Y: ny})
		}
	}
	return neighbors
}

// removeWallBetween убирает стену между двумя соседними клетками
func (p *DFSGenerator) removeWallBetween(maze *domain.Maze, a, b domain.Point) {
	wallX := (a.X + b.X) / 2
	wallY := (a.Y + b.Y) / 2
	maze.Grid[wallY][wallX].Wall = false
}
