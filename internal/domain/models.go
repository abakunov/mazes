package domain

type Point struct {
	X int
	Y int
}

type Cell struct {
	Visited bool
	Wall    bool
}

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
