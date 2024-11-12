package application

import (
	"math/rand"

	"github.com/abakunov/mazes/internal/domain"
)

type KruskalGenerator struct {
	parent map[int]int
	rank   map[int]int
}

// initUnionFind инициализирует структуру Union-Find для каждой ячейки лабиринта.
func (g *KruskalGenerator) initUnionFind(maze *domain.Maze) {
	g.parent = make(map[int]int)
	g.rank = make(map[int]int)
	for y := 0; y < maze.Height; y++ {
		for x := 0; x < maze.Width; x++ {
			index := y*maze.Width + x
			g.parent[index] = index
			g.rank[index] = 0
		}
	}
}

// find выполняет поиск сжатия пути для Union-Find.
func (g *KruskalGenerator) find(x int) int {
	if g.parent[x] != x {
		g.parent[x] = g.find(g.parent[x])
	}
	return g.parent[x]
}

// union объединяет два множества по рангу.
func (g *KruskalGenerator) union(x, y int) {
	rootX := g.find(x)
	rootY := g.find(y)

	if rootX != rootY {
		if g.rank[rootX] > g.rank[rootY] {
			g.parent[rootY] = rootX
		} else if g.rank[rootX] < g.rank[rootY] {
			g.parent[rootX] = rootY
		} else {
			g.parent[rootY] = rootX
			g.rank[rootX]++
		}
	}
}

// Generate создает лабиринт с использованием алгоритма Краскала с проверкой связности.
func (g *KruskalGenerator) Generate(maze *domain.Maze, entry, exit domain.Point) {
	for {
		// Инициализируем все клетки как стены
		for y := 0; y < maze.Height; y++ {
			for x := 0; x < maze.Width; x++ {
				maze.Grid[y][x].Wall = true
			}
		}

		// Инициализируем Union-Find для каждой ячейки
		g.initUnionFind(maze)

		// Создаем список всех возможных стен между ячейками
		var walls []struct {
			x1, y1, x2, y2 int
		}

		for y := 1; y < maze.Height-1; y += 2 {
			for x := 1; x < maze.Width-1; x += 2 {
				if x+2 < maze.Width {
					walls = append(walls, struct{ x1, y1, x2, y2 int }{x, y, x + 2, y})
				}
				if y+2 < maze.Height {
					walls = append(walls, struct{ x1, y1, x2, y2 int }{x, y, x, y + 2})
				}
			}
		}

		// Перемешиваем стены
		rand.Shuffle(len(walls), func(i, j int) { walls[i], walls[j] = walls[j], walls[i] })

		// Основной алгоритм Краскала
		for _, wall := range walls {
			// Находим индексы для Union-Find
			cell1 := wall.y1*maze.Width + wall.x1
			cell2 := wall.y2*maze.Width + wall.x2

			// Если ячейки ещё не соединены, убираем стену и объединяем их
			if g.find(cell1) != g.find(cell2) {
				g.union(cell1, cell2)
				maze.Grid[wall.y1][wall.x1].Wall = false
				maze.Grid[wall.y2][wall.x2].Wall = false
				maze.Grid[(wall.y1+wall.y2)/2][(wall.x1+wall.x2)/2].Wall = false // Убираем стену между ячейками
			}
		}

		// Устанавливаем точки входа и выхода как проходы
		maze.Grid[entry.Y][entry.X].Wall = false
		maze.Grid[exit.Y][exit.X].Wall = false

		// Проверяем, есть ли путь от входа до выхода
		if g.isPathAvailable(maze, entry, exit) {
			break // Если путь найден, выходим из цикла
		}
	}
}

// isPathAvailable проверяет, существует ли путь от точки entry до точки exit.
func (g *KruskalGenerator) isPathAvailable(maze *domain.Maze, entry, exit domain.Point) bool {
	solver := &BFSSolver{}
	path := solver.FindPath(maze, entry, exit)
	return path != nil
}
