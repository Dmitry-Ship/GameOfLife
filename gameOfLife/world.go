package gameOfLife

import (
	"math/rand"
)

type Cell struct {
	IsAlive bool
}

func (c *Cell) nextState(neighbors int) {
	if !c.IsAlive && neighbors == 3 {
		c.IsAlive = true
		return
	}

	if c.IsAlive && neighbors > 3 {
		c.IsAlive = false
		return
	}

	if c.IsAlive && neighbors < 2 {
		c.IsAlive = false
		return
	}

	if c.IsAlive && (neighbors == 2 || neighbors == 3) {
		c.IsAlive = true
		return
	}
}

type World struct {
	cells  [][]Cell
	width  int
	height int
}

func NewWorld(width, height int) *World {
	cells := make([][]Cell, height)

	getAliveStatus := func(y, x int) bool {
		// if float32(y) > (float32(height)*0.75) || float32(y) < (float32(height)*0.25) || float32(x) > (float32(width)*0.75) || float32(x) < (float32(width)*0.25) {
		// 	return false
		// }

		return rand.Intn(20) == 1
	}

	for y := range cells {
		cells[y] = make([]Cell, width)
		for x := range cells[y] {
			cells[y][x] = Cell{IsAlive: getAliveStatus(y, x)}
		}
	}

	return &World{cells, width, height}
}

func (w *World) GetCells() [][]Cell {
	return w.cells
}

func (w *World) NextGeneration() {
	newWorld := NewWorld(w.width, w.height)
	newCells := newWorld.GetCells()

	for y := range w.cells {
		for x := range w.cells[y] {
			oldCell := w.cells[y][x]
			oldNeighbors := w.getAliveNeighbors(y, x)

			newCell := Cell{IsAlive: oldCell.IsAlive}
			newCell.nextState(oldNeighbors)

			newCells[y][x] = newCell
		}
	}

	w.cells = newCells

}

func (w *World) getAliveNeighbors(cellY, cellX int) int {
	neighbors := 0

	for y := -1; y <= 1; y++ {
		for x := -1; x <= 1; x++ {
			// self
			if x == 0 && y == 0 {
				continue
			}

			// out of bounds
			if cellY+y < 0 || cellY+y >= w.height || cellX+x < 0 || cellX+x >= w.width {
				continue
			}

			if w.cells[cellY+y][cellX+x].IsAlive {
				neighbors++
			}
		}
	}

	return neighbors
}