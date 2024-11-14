package main

import (
	"fmt"
	"github.com/abakunov/mazes/internal/application"
	"github.com/abakunov/mazes/internal/domain"
	"github.com/abakunov/mazes/internal/infrastructure"
	"math/rand"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	// Настройка graceful shutdown
	exitChan := make(chan os.Signal, 1)
	signal.Notify(exitChan, os.Interrupt, syscall.SIGTERM)

	// Запускаем goroutine для обработки сигнала завершения
	go func() {
		<-exitChan
		fmt.Println("\nПрограмма завершена пользователем.")
		os.Exit(0)
	}()

	rand.Seed(time.Now().UnixNano())

	// Получаем входные данные от пользователя через отдельные функции
	width := infrastructure.GetWidth()
	height := infrastructure.GetHeight()

	// Инициализация лабиринта
	maze := domain.NewMaze(width, height)

	// Выбор алгоритма генерации
	algorithmChoice := infrastructure.GetAlgorithmChoice()

	// Выбор способа ввода точек входа и выхода
	entryExitChoice := infrastructure.GetEntryExitChoice()
	entryPoint, exitPoint := infrastructure.GetEntryExitPoints(entryExitChoice, width, height)

	// Генерация лабиринта
	switch algorithmChoice {
	case 1:
		DFSGenerator := &application.DFSGenerator{}
		DFSGenerator.Generate(maze, entryPoint, exitPoint)
	case 2:
		kruskalGenerator := &application.KruskalGenerator{}
		kruskalGenerator.Generate(maze, entryPoint, exitPoint)
	}

	// Выбор алгоритма поиска пути
	pathSolverChoice := infrastructure.GetPathSolverChoice()

	// Поиск пути
	var path []domain.Point
	switch pathSolverChoice {
	case 1:
		bfsSolver := &application.BFSSolver{}
		path = bfsSolver.FindPath(maze, entryPoint, exitPoint)
	case 2:
		aStarSolver := &application.AStarSolver{}
		path = aStarSolver.FindPath(maze, entryPoint, exitPoint)
	}

	// Отрисовка
	renderer := &infrastructure.ConsoleRenderer{}

	fmt.Println("Сгенерированный лабиринт:")
	renderer.RenderMaze(maze)

	fmt.Println("\nЛабиринт с найденным путём:")
	renderer.RenderMazeWithPath(maze, path)
}
