package application

import (
	"github.com/abakunov/mazes/internal/domain"
	"testing"
)

func TestBFSSolver_FindPath_SimplePath(t *testing.T) {
	// Create a simple 3x3 maze with a straight path
	maze := domain.NewMaze(3, 3)
	maze.Grid[0][0].Wall = false
	maze.Grid[0][1].Wall = false
	maze.Grid[0][2].Wall = false
	maze.Grid[1][2].Wall = false
	maze.Grid[2][2].Wall = false

	solver := &BFSSolver{}
	entry := domain.Point{X: 0, Y: 0}
	exit := domain.Point{X: 2, Y: 2}

	path := solver.FindPath(maze, entry, exit)

	if path == nil {
		t.Fatal("Expected a path to be found, but got nil")
	}

	// Check that the path starts at entry and ends at exit
	if path[0] != entry {
		t.Errorf("Expected path to start at %v, got %v", entry, path[0])
	}
	if path[len(path)-1] != exit {
		t.Errorf("Expected path to end at %v, got %v", exit, path[len(path)-1])
	}
}

func TestBFSSolver_FindPath_NoPath(t *testing.T) {
	// Create a 3x3 maze where the path is blocked
	maze := domain.NewMaze(3, 3)
	for y := 0; y < 3; y++ {
		for x := 0; x < 3; x++ {
			maze.Grid[y][x].Wall = true
		}
	}
	maze.Grid[0][0].Wall = false
	maze.Grid[2][2].Wall = false

	solver := &BFSSolver{}
	entry := domain.Point{X: 0, Y: 0}
	exit := domain.Point{X: 2, Y: 2}

	path := solver.FindPath(maze, entry, exit)
	if path != nil {
		t.Error("Expected no path, but found one")
	}
}

func TestBFSSolver_FindEntrance(t *testing.T) {
	// Create a 3x3 maze with an open cell on the left edge
	maze := domain.NewMaze(3, 3)
	maze.Grid[0][0].Wall = true
	maze.Grid[1][0].Wall = false
	maze.Grid[2][0].Wall = true

	solver := &BFSSolver{}
	expectedEntry := domain.Point{X: 0, Y: 1}
	entry := solver.findEntrance(maze)

	if entry != expectedEntry {
		t.Errorf("Expected entrance at %v, got %v", expectedEntry, entry)
	}
}

func TestBFSSolver_FindExit(t *testing.T) {
	// Create a 3x3 maze with an open cell on the right edge
	maze := domain.NewMaze(3, 3)
	maze.Grid[0][2].Wall = true
	maze.Grid[1][2].Wall = false
	maze.Grid[2][2].Wall = true

	solver := &BFSSolver{}
	expectedExit := domain.Point{X: 2, Y: 1}
	exit := solver.findExit(maze)

	if exit != expectedExit {
		t.Errorf("Expected exit at %v, got %v", expectedExit, exit)
	}
}
