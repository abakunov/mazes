package infrastructure

import (
	"bufio"
	"crypto/rand"
	"fmt"
	"math/big"
	"os"
	"strconv"
	"strings"

	"github.com/abakunov/mazes/internal/domain"
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
	return getIntInput("Ввести точки входа и выхода вручную или случайным образом (1 - вручную, 2 - случайным образом): ", 1,
		"Ошибка: выберите 1 (вручную) или 2 (случайным образом).", 2)
}

func GetPathSolverChoice() int {
	return getIntInput("Выберите алгоритм поиска пути (1 - BFS, 2 - A*): ", 1, "Ошибка: выберите 1 (BFS) или 2 (A*).", 2)
}

// GetEntryExitPoints gets the entry and exit points either manually or randomly.
func GetEntryExitPoints(choice, width, height int) (entryPoint, exitPoint domain.Point) {
	if choice == 1 {
		entryPoint = getValidBoundaryPoint("начальную", width, height)
		exitPoint = getValidBoundaryPoint("конечную", width, height)
		// Ensure that start and end points are not the same
		for entryPoint == exitPoint {
			fmt.Println("Ошибка: начальная и конечная точки не могут совпадать.")

			exitPoint = getValidBoundaryPoint("конечную", width, height)
		}

		return entryPoint, exitPoint
	}

	return generateRandomEntryExit(width, height)
}

// getValidBoundaryPoint prompts the user to enter a boundary point that is not in a corner.
func getValidBoundaryPoint(pointName string, width, height int) domain.Point {
	for {
		x := getIntInput(fmt.Sprintf("Введите %s точку x (от 0 до %d): ", pointName, width-1), 0,
			fmt.Sprintf("Ошибка: введите корректную координату x для %s точки (от 0 до %d).", pointName, width-1), width-1)
		y := getIntInput(fmt.Sprintf("Введите %s точку y (от 0 до %d): ", pointName, height-1), 0,
			fmt.Sprintf("Ошибка: введите корректную координату y для %s точки (от 0 до %d).", pointName, height-1), height-1)

		if isOnBoundary(x, y, width, height) && !isCorner(x, y, width, height) {
			return domain.Point{X: x, Y: y}
		}

		fmt.Println("Ошибка: точка должна находиться на границе лабиринта, но не в углу. Попробуйте снова.")
	}
}

// generateRandomEntryExit generates random entry and exit points on the boundary, excluding corners.
func generateRandomEntryExit(width, height int) (startPoint, endPoint domain.Point) {
	for {
		startPoint = randomBoundaryPoint(width, height)
		endPoint = randomBoundaryPoint(width, height)

		if startPoint != endPoint {
			break
		}
	}

	return startPoint, endPoint
}

// randomBoundaryPoint generates a random point on the boundary, excluding corners, using crypto/rand.
func randomBoundaryPoint(width, height int) domain.Point {
	for {
		edge, _ := cryptoRandInt(4)
		switch edge {
		case 0:
			x, _ := cryptoRandInt(width - 2)
			return domain.Point{X: x + 1, Y: 0}
		case 1:
			x, _ := cryptoRandInt(width - 2)
			return domain.Point{X: x + 1, Y: height - 1}
		case 2:
			y, _ := cryptoRandInt(height - 2)
			return domain.Point{X: 0, Y: y + 1}
		case 3:
			y, _ := cryptoRandInt(height - 2)
			return domain.Point{X: width - 1, Y: y + 1}
		}
	}
}

// cryptoRandInt generates a cryptographically secure random integer in the range [0, max).
func cryptoRandInt(maxI int) (int, error) {
	nBig, err := rand.Int(rand.Reader, big.NewInt(int64(maxI)))
	if err != nil {
		return 0, err
	}

	return int(nBig.Int64()), nil
}

// isOnBoundary checks if the point is on the boundary of the maze.
func isOnBoundary(x, y, width, height int) bool {
	return x == 0 || x == width-1 || y == 0 || y == height-1
}

// isCorner checks if a point is in one of the corners.
func isCorner(x, y, width, height int) bool {
	return (x == 0 && y == 0) || (x == width-1 && y == 0) || (x == 0 && y == height-1) || (x == width-1 && y == height-1)
}

// getOddIntInput requests an odd integer from the user and checks input validity.
func getOddIntInput(prompt string, minimum int, errorMessage string) int {
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

		if convErr == nil && value >= minimum && value%2 != 0 {
			return value
		}

		fmt.Println(errorMessage + " Принимаются только нечетные числа.")
	}
}

// getIntInput requests an integer from the user and checks input validity.
func getIntInput(prompt string, minimum int, errorMessage string, maximum ...int) int {
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

		if convErr == nil && value >= minimum && (len(maximum) == 0 || value <= maximum[0]) {
			return value
		}

		fmt.Println(errorMessage)
	}
}
