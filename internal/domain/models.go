package domain

type Point struct {
	X int
	Y int
}

// Cell представляет собой клетку лабиринта (проход или стена).
type Cell struct {
	Visited bool
	Wall    bool
}

// Maze представляет собой сам лабиринт.
type Maze struct {
	Width  int
	Height int
	Grid   [][]Cell
}

func NewMaze(width, height int) *Maze {
	cells := make([][]Cell, height)
	for i := range cells {
		cells[i] = make([]Cell, width)
	}

	return &Maze{
		Width:  width,
		Height: height,
		Grid:   cells,
	}
}
