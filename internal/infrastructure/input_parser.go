package infrastructure

import (
	"bufio"
	"fmt"
	"github.com/abakunov/mazes/internal/domain"
	"math/rand"
	"os"
	"strconv"
	"strings"
)

// Функции для запроса различных параметров

func GetWidth() int {
	return getOddIntInput("Ширина (только нечетные значения): ", 3, "Введите корректное нечетное число (минимум 3) для ширины.")
}

func GetHeight() int {
	return getOddIntInput("Высота (только нечетные значения): ", 3, "Введите корректное нечетное число (минимум 3) для высоты.")
}

func GetAlgorithmChoice() int {
	return getIntInput("Выберите алгоритм генерации лабиринта (1 - DFS, 2 - Kruskal): ", 1, "Ошибка: выберите 1 (DFS) или 2 (Kruskal).", 2)
}

func GetEntryExitChoice() int {
	return getIntInput("Ввести точки входа и выхода вручную или случайным образом (1 - вручную, 2 - случайным образом): ", 1, "Ошибка: выберите 1 (вручную) или 2 (случайным образом).", 2)
}

func GetPathSolverChoice() int {
	return getIntInput("Выберите алгоритм поиска пути (1 - BFS, 2 - A*): ", 1, "Ошибка: выберите 1 (BFS) или 2 (A*).", 2)
}

// GetEntryExitPoints получает точки входа и выхода либо вручную, либо случайным образом
func GetEntryExitPoints(choice, width, height int) (domain.Point, domain.Point) {
	var startX, startY, endX, endY int

	if choice == 1 {
		startX = getIntInput(fmt.Sprintf("Введите начальную точку x (от 0 до %d): ", width-1), 0, fmt.Sprintf("Ошибка: введите корректную координату x для начальной точки (от 0 до %d).", width-1), width-1)
		startY = getIntInput(fmt.Sprintf("Введите начальную точку y (от 0 до %d): ", height-1), 0, fmt.Sprintf("Ошибка: введите корректную координату y для начальной точки (от 0 до %d).", height-1), height-1)
		endX = getIntInput(fmt.Sprintf("Введите конечную точку x (от 0 до %d): ", width-1), 0, fmt.Sprintf("Ошибка: введите корректную координату x для конечной точки (от 0 до %d).", width-1), width-1)
		endY = getIntInput(fmt.Sprintf("Введите конечную точку y (от 0 до %d): ", height-1), 0, fmt.Sprintf("Ошибка: введите корректную координату y для конечной точки (от 0 до %d).", height-1), height-1)

		if startX == endX && startY == endY {
			fmt.Println("Ошибка: начальная и конечная точки не могут совпадать.")
			os.Exit(1)
		}
	} else {
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

	return domain.Point{X: startX, Y: startY}, domain.Point{X: endX, Y: endY}
}

// getOddIntInput запрашивает нечетное целое число у пользователя и проверяет корректность ввода
func getOddIntInput(prompt string, min int, errorMessage string) int {
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print(prompt)
		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("\nПрограмма завершена пользователем.")
			os.Exit(0)
		}

		input = strings.TrimSpace(input)
		value, convErr := strconv.Atoi(input)
		if convErr == nil && value >= min && value%2 != 0 {
			return value
		}

		fmt.Println(errorMessage + " Принимаются только нечетные числа.")
	}
}

// getIntInput запрашивает целое число у пользователя и проверяет корректность ввода
func getIntInput(prompt string, min int, errorMessage string, max ...int) int {
	reader := bufio.NewReader(os.Stdin)
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
