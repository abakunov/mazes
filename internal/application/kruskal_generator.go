package application

import (
	"math/rand"

	"github.com/abakunov/mazes/internal/domain"
)

type KruskalGenerator struct {
	parent map[int]int
	rank   map[int]int
}

// initUnionFind initializes the Union-Find structure for each cell in the maze.
func (g *KruskalGenerator) initUnionFind(maze *domain.Maze) {
	g.parent = make(map[int]int)
	g.rank = make(map[int]int)
	for y := 0; y < maze.Height; y++ {
		for x := 0; x < maze.Width; x++ {
			index := y*maze.Width + x
			g.parent[index] = index
			g.rank[index] = 0
		}
	}
}

// find performs path compression for Union-Find.
func (g *KruskalGenerator) find(x int) int {
	if g.parent[x] != x {
		g.parent[x] = g.find(g.parent[x])
	}
	return g.parent[x]
}

// union merges two sets by rank.
func (g *KruskalGenerator) union(x, y int) {
	rootX := g.find(x)
	rootY := g.find(y)

	if rootX != rootY {
		if g.rank[rootX] > g.rank[rootY] {
			g.parent[rootY] = rootX
		} else if g.rank[rootX] < g.rank[rootY] {
			g.parent[rootX] = rootY
		} else {
			g.parent[rootY] = rootX
			g.rank[rootX]++
		}
	}
}

// Generate creates a maze using Kruskal's algorithm with connectivity checking.
func (g *KruskalGenerator) Generate(maze *domain.Maze, entry, exit domain.Point) {
	for {
		// Initialize all cells as walls
		for y := 0; y < maze.Height; y++ {
			for x := 0; x < maze.Width; x++ {
				maze.Grid[y][x].Wall = true
			}
		}

		// Initialize Union-Find for each cell
		g.initUnionFind(maze)

		// Create a list of all possible walls between cells
		var walls []struct {
			x1, y1, x2, y2 int
		}

		for y := 1; y < maze.Height-1; y += 2 {
			for x := 1; x < maze.Width-1; x += 2 {
				if x+2 < maze.Width {
					walls = append(walls, struct{ x1, y1, x2, y2 int }{x, y, x + 2, y})
				}
				if y+2 < maze.Height {
					walls = append(walls, struct{ x1, y1, x2, y2 int }{x, y, x, y + 2})
				}
			}
		}

		// Shuffle the walls
		rand.Shuffle(len(walls), func(i, j int) { walls[i], walls[j] = walls[j], walls[i] })

		// Main Kruskal's algorithm
		for _, wall := range walls {
			// Find indices for Union-Find
			cell1 := wall.y1*maze.Width + wall.x1
			cell2 := wall.y2*maze.Width + wall.x2

			// If the cells are not yet connected, remove the wall and unite them
			if g.find(cell1) != g.find(cell2) {
				g.union(cell1, cell2)
				maze.Grid[wall.y1][wall.x1].Wall = false
				maze.Grid[wall.y2][wall.x2].Wall = false
				maze.Grid[(wall.y1+wall.y2)/2][(wall.x1+wall.x2)/2].Wall = false // Remove wall between cells
			}
		}

		// Set entry and exit points as passages
		maze.Grid[entry.Y][entry.X].Wall = false
		maze.Grid[exit.Y][exit.X].Wall = false

		// Check if there is a path from entry to exit
		if g.isPathAvailable(maze, entry, exit) {
			break // Exit the loop if a path is found
		}
	}
}

// isPathAvailable checks if there is a path from the entry point to the exit point.
func (g *KruskalGenerator) isPathAvailable(maze *domain.Maze, entry, exit domain.Point) bool {
	solver := &BFSSolver{}
	path := solver.FindPath(maze, entry, exit)
	return path != nil
}
