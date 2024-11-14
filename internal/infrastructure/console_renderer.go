package infrastructure

import (
	"fmt"

	"github.com/fatih/color"

	"github.com/abakunov/mazes/internal/domain"
)

type ConsoleRenderer struct{}

func (r *ConsoleRenderer) RenderMaze(maze *domain.Maze) {
	wallColor := color.New(color.FgRed).SprintFunc()
	pathColor := color.New(color.FgWhite).SprintFunc()

	for y := 0; y < maze.Height; y++ {
		for x := 0; x < maze.Width; x++ {
			if maze.Grid[y][x].Wall {
				fmt.Print(wallColor("██"))
			} else {
				fmt.Print(pathColor("  "))
			}
		}

		fmt.Println()
	}
}

func (r *ConsoleRenderer) RenderMazeWithPath(maze *domain.Maze, path []domain.Point) {
	wallColor := color.New(color.FgRed).SprintFunc()
	pathColor := color.New(color.FgWhite).SprintFunc()
	solutionColor := color.New(color.BgGreen).SprintFunc()

	pathSet := make(map[domain.Point]bool)
	for _, p := range path {
		pathSet[p] = true
	}

	for y := 0; y < maze.Height; y++ {
		for x := 0; x < maze.Width; x++ {
			p := domain.Point{X: x, Y: y}

			switch {
			case maze.Grid[y][x].Wall:
				fmt.Print(wallColor("██"))
			case pathSet[p]:
				fmt.Print(solutionColor("  "))
			default:
				fmt.Print(pathColor("  "))
			}
		}

		fmt.Println()
	}
}
