package gameOfLife

import (
	"math/rand"
	"time"
)

type World struct {
	cells  [][]Cell
	width  int
	height int
}

func NewWorld(width, height, density int) *World {
	rand.Seed(time.Now().UnixNano())
	cells := make([][]Cell, height)

	getAliveStatus := func() bool {
		return rand.Intn(12-density) == 1
	}

	for y := range cells {
		cells[y] = make([]Cell, width)
		for x := range cells[y] {
			cells[y][x] = Cell{IsAlive: getAliveStatus()}
		}
	}

	return &World{cells, width, height}
}

func (w *World) GetCells() [][]Cell {
	return w.cells
}

func (w *World) GetCellNextState(c *Cell, y, x int) Cell {
	neighbors := w.getAliveNeighbors(y, x)
	// Any dead cell with exactly three live neighbours becomes a live cell, as if by reproduction.
	if !c.IsAlive && neighbors == 3 {
		return Cell{IsAlive: true}
	}

	// Any live cell with more than three live neighbours dies, as if by overpopulation.
	if c.IsAlive && neighbors > 3 {
		return Cell{IsAlive: false}
	}

	// Any live cell with fewer than two live neighbours dies, as if caused by under-population.
	if c.IsAlive && neighbors < 2 {
		return Cell{IsAlive: false}
	}

	// Any live cell with two or three live neighbours lives on to the next generation.
	if c.IsAlive && (neighbors == 2 || neighbors == 3) {
		return Cell{IsAlive: true}
	}

	return Cell{IsAlive: c.IsAlive}
}

func (w *World) NextGeneration() {
	newCells := make([][]Cell, w.height)
	for i := range newCells {
		newCells[i] = make([]Cell, w.width)
	}

	for y := range w.cells {
		for x, oldCell := range w.cells[y] {

			newCells[y][x] = w.GetCellNextState(&oldCell, y, x)
		}
	}

	w.cells = newCells
}

func (w *World) getAliveNeighbors(cellY, cellX int) int {
	aliveNeighbors := 0

	for y := -1; y <= 1; y++ {
		// out of bounds
		if cellY+y < 0 || cellY+y >= w.height {
			continue
		}
		for x := -1; x <= 1; x++ {
			// out of bounds
			if cellX+x < 0 || cellX+x >= w.width {
				continue
			}

			// self
			if x == 0 && y == 0 {
				continue
			}

			if w.cells[cellY+y][cellX+x].IsAlive {
				aliveNeighbors++
			}
		}
	}

	return aliveNeighbors
}
