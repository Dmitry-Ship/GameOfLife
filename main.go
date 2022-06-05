package main

import (
	"log"

	"GitHub/GameOfLife/game"
	"GitHub/GameOfLife/gameOfLife"

	"github.com/hajimehoshi/ebiten"
)

func main() {
	world := gameOfLife.NewWorld(900, 450)
	// world := gameOfLife.NewWorld(180, 90)

	g := game.NewGame(world, 900, 1800)
	ebiten.SetWindowSize(1800, 900)
	ebiten.SetWindowTitle("Game of life")

	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}
}
