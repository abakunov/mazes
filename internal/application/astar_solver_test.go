package application_test

import (
	"container/heap"
	"testing"

	"github.com/abakunov/mazes/internal/application"
	"github.com/abakunov/mazes/internal/domain"
)

func TestAStarSolver_FindPath_SimplePath(t *testing.T) {
	// Create a simple 3x3 maze with a straight path
	maze := domain.NewMaze(3, 3)
	maze.Grid[0][0].Wall = false
	maze.Grid[0][1].Wall = false
	maze.Grid[0][2].Wall = false
	maze.Grid[1][2].Wall = false
	maze.Grid[2][2].Wall = false

	solver := &application.AStarSolver{}
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

func TestAStarSolver_FindPath_NoPath(t *testing.T) {
	// Create a 3x3 maze where the path is blocked
	maze := domain.NewMaze(3, 3)

	for y := 0; y < 3; y++ {
		for x := 0; x < 3; x++ {
			maze.Grid[y][x].Wall = true
		}
	}

	maze.Grid[0][0].Wall = false
	maze.Grid[2][2].Wall = false

	solver := &application.AStarSolver{}
	entry := domain.Point{X: 0, Y: 0}
	exit := domain.Point{X: 2, Y: 2}

	path := solver.FindPath(maze, entry, exit)
	if path != nil {
		t.Error("Expected no path, but found one")
	}
}

func TestPriorityQueue(t *testing.T) {
	// Test the functionality of the priority queue
	pq := &application.PriorityQueue{}
	heap.Init(pq)

	nodes := []*application.Node{
		{Point: domain.Point{X: 0, Y: 0}, Priority: 3},
		{Point: domain.Point{X: 1, Y: 1}, Priority: 1},
		{Point: domain.Point{X: 2, Y: 2}, Priority: 2},
	}

	// Add nodes to the queue
	for _, node := range nodes {
		heap.Push(pq, node)
	}

	// Check that nodes are extracted in order of priority
	expectedPriorities := []int{1, 2, 3}
	for _, expected := range expectedPriorities {
		node := heap.Pop(pq).(*application.Node)
		if node.Priority != expected {
			t.Errorf("Expected priority %d, got %d", expected, node.Priority)
		}
	}
}
