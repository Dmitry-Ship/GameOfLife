package game

import (
	"image/color"

	"GitHub/GameOfLife/gameOfLife"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
)

type World interface {
	NextGeneration()
	GetCells() [][]gameOfLife.Cell
}

type Game struct {
	world           World
	width           int
	height          int
	cellSize        int
	cellSizeFloat64 float64
}

func NewGame(world World, width int, height int, cellSize int) *Game {
	return &Game{
		world:           world,
		width:           width,
		height:          height,
		cellSize:        cellSize,
		cellSizeFloat64: float64(cellSize),
	}
}

func (g *Game) Update(screen *ebiten.Image) error {
	g.world.NextGeneration()

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	cells := g.world.GetCells()

	for y := range cells {
		yFloat64 := float64(y) * g.cellSizeFloat64
		for x, cell := range cells[y] {
			xFloat64 := float64(x) * g.cellSizeFloat64
			if cell.IsAlive {
				ebitenutil.DrawRect(screen, xFloat64, yFloat64, g.cellSizeFloat64, g.cellSizeFloat64, color.White)
			}
		}
	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return g.height, g.width
}
