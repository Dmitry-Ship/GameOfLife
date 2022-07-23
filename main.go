package main

import (
	"flag"
	"log"

	"GitHub/GameOfLife/game"
	"GitHub/GameOfLife/gameOfLife"

	"github.com/hajimehoshi/ebiten"
)

func main() {
	width := flag.Int("width", 1800, "world width")
	height := flag.Int("height", 900, "world height")
	cellSize := flag.Int("cellSize", 2, "sell size")
	density := flag.Int("density", 5, "density max: 10")

	flag.Parse()

	if *density > 10 {
		*density = 10
	}

	world := gameOfLife.NewWorld(*width/(*cellSize), *height/(*cellSize), *density)
	g := game.NewGame(world, *height, *width, *cellSize)
	ebiten.SetWindowSize(*width, *height)
	ebiten.SetWindowTitle("Game of life")

	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}
}
