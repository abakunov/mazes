package application

import (
	"github.com/abakunov/mazes/internal/domain"
	"math/rand"
	"time"
)

type DFSGenerator struct{}

// Generate creates a maze using the DFS (Depth-First Search) algorithm.
func (p *DFSGenerator) Generate(maze *domain.Maze, entryPoint, exitPoint domain.Point) {
	rand.Seed(time.Now().UnixNano())

	// Initialize all cells as walls
	for y := 0; y < maze.Height; y++ {
		for x := 0; x < maze.Width; x++ {
			maze.Grid[y][x] = domain.Cell{Wall: true, Visited: false}
		}
	}

	// Set outer boundaries as walls, except for entry and exit points
	p.setOuterWalls(maze, entryPoint, exitPoint)

	// Start generation from the entry point, excluding outer boundaries
	stack := []domain.Point{entryPoint}
	maze.Grid[entryPoint.Y][entryPoint.X].Visited = true
	maze.Grid[entryPoint.Y][entryPoint.X].Wall = false

	// Depth-First Search
	for len(stack) > 0 {
		// Take the last point from the stack
		current := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		// Get a list of unvisited neighbors
		neighbors := p.getUnvisitedNeighbors(maze, current)
		if len(neighbors) > 0 {
			// Add the current point back to the stack
			stack = append(stack, current)

			// Choose a random unvisited neighbor
			next := neighbors[rand.Intn(len(neighbors))]

			// Remove the wall between the current cell and the chosen neighbor
			p.removeWallBetween(maze, current, next)

			// Mark the neighbor as visited and add it to the stack
			maze.Grid[next.Y][next.X].Visited = true
			maze.Grid[next.Y][next.X].Wall = false
			stack = append(stack, next)
		}
	}

	// Connect the exit point to the maze if it is isolated
	p.connectExitPoint(maze, exitPoint)
}

// setOuterWalls sets the outer boundaries as walls, leaving passages at the entry and exit points
func (p *DFSGenerator) setOuterWalls(maze *domain.Maze, entryPoint, exitPoint domain.Point) {
	for x := 0; x < maze.Width; x++ {
		maze.Grid[0][x].Wall = true
		maze.Grid[maze.Height-1][x].Wall = true
	}
	for y := 0; y < maze.Height; y++ {
		maze.Grid[y][0].Wall = true
		maze.Grid[y][maze.Width-1].Wall = true
	}

	// Leave passages for entry and exit
	maze.Grid[entryPoint.Y][entryPoint.X].Wall = false
	maze.Grid[exitPoint.Y][exitPoint.X].Wall = false
}

// connectExitPoint connects the exit point to the maze if it is isolated
func (p *DFSGenerator) connectExitPoint(maze *domain.Maze, exitPoint domain.Point) {
	// Find neighbors of the exit point
	neighbors := p.getUnvisitedNeighbors(maze, exitPoint)
	if len(neighbors) > 0 {
		// Choose a random neighbor
		next := neighbors[rand.Intn(len(neighbors))]

		// Remove the wall between the exit point and the chosen neighbor
		p.removeWallBetween(maze, exitPoint, next)

		// Mark the neighbor as visited
		maze.Grid[next.Y][next.X].Visited = true
		maze.Grid[next.Y][next.X].Wall = false
	}
}

// getUnvisitedNeighbors returns a list of unvisited neighbors for the specified cell
func (p *DFSGenerator) getUnvisitedNeighbors(maze *domain.Maze, cell domain.Point) []domain.Point {
	var neighbors []domain.Point
	directions := []struct{ x, y int }{{2, 0}, {-2, 0}, {0, 2}, {0, -2}}

	for _, d := range directions {
		nx, ny := cell.X+d.x, cell.Y+d.y
		if nx > 0 && nx < maze.Width-1 && ny > 0 && ny < maze.Height-1 && !maze.Grid[ny][nx].Visited {
			neighbors = append(neighbors, domain.Point{X: nx, Y: ny})
		}
	}
	return neighbors
}

// removeWallBetween removes the wall between two neighboring cells
func (p *DFSGenerator) removeWallBetween(maze *domain.Maze, a, b domain.Point) {
	wallX := (a.X + b.X) / 2
	wallY := (a.Y + b.Y) / 2
	maze.Grid[wallY][wallX].Wall = false
}
