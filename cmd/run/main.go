package main

import (
	"fmt"

	"github.com/abakunov/mazes/internal/application"
	"github.com/abakunov/mazes/internal/domain"
	"github.com/abakunov/mazes/internal/infrastructure"
)

func main() {
	width, height := 21, 21

	// Инициализация лабиринта
	maze := domain.NewMaze(width, height)

	// Установка стен по умолчанию
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			if x%2 == 1 && y%2 == 1 {
				maze.Grid[y][x].Wall = false
			} else {
				maze.Grid[y][x].Wall = true
			}
		}
	}

	// Добавляем вход и выход
	maze.Grid[1][0].Wall = false
	maze.Grid[height-2][width-1].Wall = false

	// Генерация лабиринта
	generator := &application.DFSGenerator{}
	generator.Generate(maze, domain.Point{X: 1, Y: 1})

	// Поиск пути
	solver := &application.BFSSolver{}
	start := domain.Point{X: 1, Y: 0}
	end := domain.Point{X: width - 2, Y: height - 1}
	path := solver.FindPath(maze, start, end)

	// Отрисовка
	renderer := &infrastructure.ConsoleRenderer{}

	fmt.Println("Сгенерированный лабиринт:")
	renderer.RenderMaze(maze)

	fmt.Println("\nЛабиринт с найденным путём:")
	renderer.RenderMazeWithPath(maze, path)
}
