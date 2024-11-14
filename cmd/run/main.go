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
	// Setting up graceful shutdown
	exitChan := make(chan os.Signal, 1)
	signal.Notify(exitChan, os.Interrupt, syscall.SIGTERM)

	// Launch a goroutine to handle the shutdown signal
	go func() {
		<-exitChan
		fmt.Println("\nProgram terminated by the user.")
		os.Exit(0)
	}()

	rand.Seed(time.Now().UnixNano())

	// Get input data from the user through separate functions
	width := infrastructure.GetWidth()
	height := infrastructure.GetHeight()

	// Maze initialization
	maze := domain.NewMaze(width, height)

	// Select generation algorithm
	algorithmChoice := infrastructure.GetAlgorithmChoice()

	// Define the maze generator corresponding to the Generator interface
	var generator domain.Generator
	switch algorithmChoice {
	case 1:
		generator = &application.DFSGenerator{}
	case 2:
		generator = &application.KruskalGenerator{}
	default:
		fmt.Println("Error: invalid generation algorithm choice.")
		os.Exit(1)
	}

	// Select the method for entering start and exit points
	entryExitChoice := infrastructure.GetEntryExitChoice()
	entryPoint, exitPoint := infrastructure.GetEntryExitPoints(entryExitChoice, width, height)

	// Maze generation
	generator.Generate(maze, entryPoint, exitPoint)

	// Select pathfinding algorithm
	pathSolverChoice := infrastructure.GetPathSolverChoice()

	// Define the pathfinding algorithm corresponding to the Solver interface
	var solver domain.Solver
	switch pathSolverChoice {
	case 1:
		solver = &application.BFSSolver{}
	case 2:
		solver = &application.AStarSolver{}
	default:
		fmt.Println("Error: invalid pathfinding algorithm choice.")
		os.Exit(1)
	}

	// Pathfinding
	path := solver.FindPath(maze, entryPoint, exitPoint)

	// Rendering
	renderer := &infrastructure.ConsoleRenderer{}

	fmt.Println("Generated maze:")
	renderer.RenderMaze(maze)

	fmt.Println("\nMaze with found path:")
	renderer.RenderMazeWithPath(maze, path)
}
