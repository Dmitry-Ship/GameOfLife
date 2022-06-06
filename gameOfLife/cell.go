package gameOfLife

type Cell struct {
	IsAlive bool
}

func (c *Cell) nextState(neighbors int) {
	// Any dead cell with exactly three live neighbours becomes a live cell, as if by reproduction.
	if !c.IsAlive && neighbors == 3 {
		c.IsAlive = true
		return
	}

	// Any live cell with more than three live neighbours dies, as if by overpopulation.
	if c.IsAlive && neighbors > 3 {
		c.IsAlive = false
		return
	}

	// Any live cell with fewer than two live neighbours dies, as if caused by under-population.
	if c.IsAlive && neighbors < 2 {
		c.IsAlive = false
		return
	}

	// Any live cell with two or three live neighbours lives on to the next generation.
	if c.IsAlive && (neighbors == 2 || neighbors == 3) {
		c.IsAlive = true
		return
	}
}
