package application

import (
	"container/heap"
	"github.com/abakunov/mazes/internal/domain"
	"math"
)

// Node представляет узел в A*
type Node struct {
	Point    domain.Point
	Cost     int
	Priority int
	Index    int
	Parent   *Node
}

// PriorityQueue реализует очередь с приоритетом для алгоритма A*
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
	// Инициализация очереди с приоритетом
	pq := &PriorityQueue{}
	heap.Init(pq)

	// Создание начального узла
	startNode := &Node{
		Point:    entry,
		Cost:     0,
		Priority: s.heuristic(entry, exit),
		Parent:   nil,
	}
	heap.Push(pq, startNode)

	// Храним посещённые узлы
	visited := make(map[domain.Point]bool)
	visited[entry] = true

	// Направления для движения: вверх, вправо, вниз, влево
	directions := []domain.Point{
		{X: 0, Y: -1}, // Вверх
		{X: 1, Y: 0},  // Вправо
		{X: 0, Y: 1},  // Вниз
		{X: -1, Y: 0}, // Влево
	}

	// A* поиск пути
	for pq.Len() > 0 {
		// Извлекаем узел с наименьшим приоритетом
		currentNode := heap.Pop(pq).(*Node)
		currentPoint := currentNode.Point

		// Если достигли точки выхода, восстанавливаем путь
		if currentPoint == exit {
			path := []domain.Point{}
			for node := currentNode; node != nil; node = node.Parent {
				path = append([]domain.Point{node.Point}, path...)
			}
			return path
		}

		// Проходим по всем соседям
		for _, dir := range directions {
			neighborPoint := domain.Point{X: currentPoint.X + dir.X, Y: currentPoint.Y + dir.Y}

			// Проверяем, что соседняя точка находится в пределах лабиринта и является проходом (не стеной)
			if neighborPoint.X >= 0 && neighborPoint.X < maze.Width && neighborPoint.Y >= 0 && neighborPoint.Y < maze.Height &&
				!maze.Grid[neighborPoint.Y][neighborPoint.X].Wall && !visited[neighborPoint] {

				// Вычисляем стоимость перехода и приоритет
				newCost := currentNode.Cost + 1
				priority := newCost + s.heuristic(neighborPoint, exit)

				// Создаём узел для соседа
				neighborNode := &Node{
					Point:    neighborPoint,
					Cost:     newCost,
					Priority: priority,
					Parent:   currentNode,
				}

				// Добавляем соседа в очередь с приоритетом и отмечаем его как посещённого
				heap.Push(pq, neighborNode)
				visited[neighborPoint] = true
			}
		}
	}

	// Путь не найден
	return nil
}

// heuristic вычисляет эвристику Манхэттена (расстояние от текущей точки до точки выхода)
func (s *AStarSolver) heuristic(a, b domain.Point) int {
	return int(math.Abs(float64(a.X-b.X)) + math.Abs(float64(a.Y-b.Y)))
}
