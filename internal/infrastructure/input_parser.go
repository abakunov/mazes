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

// Functions for requesting various parameters

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

// GetEntryExitPoints gets the entry and exit points either manually or randomly
// GetEntryExitPoints gets the entry and exit points either manually or randomly
func GetEntryExitPoints(choice, width, height int) (domain.Point, domain.Point) {
	var startX, startY, endX, endY int

	// Function to check if a point is in the corners
	isCorner := func(x, y int) bool {
		return (x == 0 && y == 0) || (x == width-1 && y == 0) ||
			(x == 0 && y == height-1) || (x == width-1 && y == height-1)
	}

	if choice == 1 {
		for {
			startX = getIntInput(fmt.Sprintf("Введите начальную точку x (от 0 до %d): ", width-1), 0, fmt.Sprintf("Ошибка: введите корректную координату x для начальной точки (от 0 до %d).", width-1), width-1)
			startY = getIntInput(fmt.Sprintf("Введите начальную точку y (от 0 до %d): ", height-1), 0, fmt.Sprintf("Ошибка: введите корректную координату y для начальной точки (от 0 до %d).", height-1), height-1)

			// Check if the start point is on the boundary and not in a corner
			if (startX == 0 || startX == width-1 || startY == 0 || startY == height-1) && !isCorner(startX, startY) {
				break
			}
			fmt.Println("Ошибка: начальная точка должна находиться на границе лабиринта, но не в углу. Попробуйте снова.")
		}

		for {
			endX = getIntInput(fmt.Sprintf("Введите конечную точку x (от 0 до %d): ", width-1), 0, fmt.Sprintf("Ошибка: введите корректную координату x для конечной точки (от 0 до %d).", width-1), width-1)
			endY = getIntInput(fmt.Sprintf("Введите конечную точку y (от 0 до %d): ", height-1), 0, fmt.Sprintf("Ошибка: введите корректную координату y для конечной точки (от 0 до %d).", height-1), height-1)

			// Check if the end point is on the boundary, not the same as the start point, and not in a corner
			if (endX == 0 || endX == width-1 || endY == 0 || endY == height-1) && !isCorner(endX, endY) && (endX != startX || endY != startY) {
				break
			}
			if endX == startX && endY == startY {
				fmt.Println("Ошибка: начальная и конечная точки не могут совпадать.")
			} else {
				fmt.Println("Ошибка: конечная точка должна находиться на границе лабиринта, но не в углу. Попробуйте снова.")
			}
		}
	} else {
		for {
			// Generate entry point
			switch rand.Intn(4) {
			case 0: // Top border
				startX, startY = rand.Intn(width-2)+1, 0
			case 1: // Bottom border
				startX, startY = rand.Intn(width-2)+1, height-1
			case 2: // Left border
				startX, startY = 0, rand.Intn(height-2)+1
			case 3: // Right border
				startX, startY = width-1, rand.Intn(height-2)+1
			}

			// Generate exit point
			switch rand.Intn(4) {
			case 0: // Top border
				endX, endY = rand.Intn(width-2)+1, 0
			case 1: // Bottom border
				endX, endY = rand.Intn(width-2)+1, height-1
			case 2: // Left border
				endX, endY = 0, rand.Intn(height-2)+1
			case 3: // Right border
				endX, endY = width-1, rand.Intn(height-2)+1
			}

			// Ensure that start and end points are not the same and are not in corners
			if (startX != endX || startY != endY) && !isCorner(startX, startY) && !isCorner(endX, endY) {
				break
			}
		}
	}

	return domain.Point{X: startX, Y: startY}, domain.Point{X: endX, Y: endY}
}

// getOddIntInput requests an odd integer from the user and checks input validity
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

// getIntInput requests an integer from the user and checks input validity
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
