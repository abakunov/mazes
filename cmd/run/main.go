package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"os/signal"
	"strconv"
	"strings"
	"syscall"
	"time"
	"github.com/abakunov/mazes/internal/domain"
	"github.com/abakunov/mazes/internal/application"
	"github.com/abakunov/mazes/internal/infrastructure"
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
	reader := bufio.NewReader(os.Stdin)
	
	width := getIntInput("Ширина: ", reader, 3, "Введите корректное целое число (минимум 3) для ширины.")
	height := getIntInput("Высота: ", reader, 3, "Введите корректное целое число (минимум 3) для высоты.")
	
	// Инициализация лабиринта
	maze := domain.NewMaze(width, height)

	// Выбор алгоритма генерации
	algorithmChoice := getIntInput("Выберите алгоритм генерации лабиринта (1 - DFS, 2 - Kruskal): ", reader, 1, "Ошибка: выберите 1 (DFS) или 2 (Kruskal).", 2)

	// Выбор способа ввода точек входа и выхода
	entryExitChoice := getIntInput("Ввести точки входа и выхода вручную или случайным образом (1 - вручную, 2 - случайным образом): ", reader, 1, "Ошибка: выберите 1 (вручную) или 2 (случайным образом).", 2)

	var startX, startY, endX, endY int

	switch entryExitChoice {
	case 1:
		startX = getIntInput("Введите начальную точку x: ", reader, 0, fmt.Sprintf("Ошибка: введите корректную координату x для начальной точки (от 0 до %d).", width-1), width-1)
		startY = getIntInput("Введите начальную точку y: ", reader, 0, fmt.Sprintf("Ошибка: введите корректную координату y для начальной точки (от 0 до %d).", height-1), height-1)
		endX = getIntInput("Введите конечную точку x: ", reader, 0, fmt.Sprintf("Ошибка: введите корректную координату x для конечной точки (от 0 до %d).", width-1), width-1)
		endY = getIntInput("Введите конечную точку y: ", reader, 0, fmt.Sprintf("Ошибка: введите корректную координату y для конечной точки (от 0 до %d).", height-1), height-1)

		if startX == endX && startY == endY {
			fmt.Println("Ошибка: начальная и конечная точки не могут совпадать.")
			return
		}
	case 2:
		// Генерация случайных точек на границах, но не в углах и не равных друг другу
		for {
			// Генерация точки входа
			switch rand.Intn(4) {
			case 0: // Верхняя граница
				startX, startY = rand.Intn(width-2)+1, 0
			case 1: // Нижняя граница
				startX, startY = rand.Intn(width-2)+1, height-1
			case 2: // Левая граница
				startX, startY = 0, rand.Intn(height-2)+1
			case 3: // Правая граница
				startX, startY = width-1, rand.Intn(height-2)+1
			}

			// Генерация точки выхода
			switch rand.Intn(4) {
			case 0: // Верхняя граница
				endX, endY = rand.Intn(width-2)+1, 0
			case 1: // Нижняя граница
				endX, endY = rand.Intn(width-2)+1, height-1
			case 2: // Левая граница
				endX, endY = 0, rand.Intn(height-2)+1
			case 3: // Правая граница
				endX, endY = width-1, rand.Intn(height-2)+1
			}

			// Проверка, что точки не равны
			if startX != endX || startY != endY {
				break
			}
		}
	}

	entryPoint := domain.Point{X: startX, Y: startY}
	exitPoint := domain.Point{X: endX, Y: endY}

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
	pathSolverChoice := getIntInput("Выберите алгоритм поиска пути (1 - BFS, 2 - A*): ", reader, 1, "Ошибка: выберите 1 (BFS) или 2 (A*).", 2)

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

// getIntInput запрашивает целое число у пользователя и проверяет корректность ввода
func getIntInput(prompt string, reader *bufio.Reader, min int, errorMessage string, max ...int) int {
	for {
		fmt.Print(prompt)
		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("\nПрограмма завершена пользователем.")
			os.Exit(0)
		}

		input = strings.TrimSpace(input)
		value, convErr := strconv.Atoi(input)
		if convErr == nil && value >= min && (len(max) == 0 || value <= max[0]) {
			return value
		}

		fmt.Println(errorMessage)
	}
}
