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
	world  World
	width  int
	height int
}

func NewGame(world World, width int, height int) *Game {
	return &Game{world, width, height}
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
				ebitenutil.DrawRect(screen, float64(x*2), float64(y*2), 2, 2, color.White)
			}
		}
	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return g.height, g.width
}
