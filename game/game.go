package game

import (
	"image/color"

	"GitHub/GameOfLife/gameOfLife"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
)

type World interface {
	NextGeneration()
	GetCells() [][]*gameOfLife.Cell
}

type Game struct {
	world    World
	width    int
	height   int
	cellSize int
}

func NewGame(world World, width int, height int, cellSize int) *Game {
	return &Game{
		world:    world,
		width:    width,
		height:   height,
		cellSize: cellSize}
}

func (g *Game) Update(screen *ebiten.Image) error {
	g.world.NextGeneration()

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	cells := g.world.GetCells()

	for y := range cells {
		for x := range cells[y] {
			if cells[y][x].IsAlive {
				ebitenutil.DrawRect(screen, float64(x*g.cellSize), float64(y*g.cellSize), float64(g.cellSize), float64(g.cellSize), color.White)
			}
		}
	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return g.height, g.width
}
