package main

import (
	"log"
	"math/rand"
	"time"

	"GitHub/GameOfLife/game"
	"GitHub/GameOfLife/gameOfLife"

	"github.com/hajimehoshi/ebiten"
)

func main() {
	width := 1800
	height := 900
	cellSize := 2
	rand.Seed(time.Now().UnixNano())

	world := gameOfLife.NewWorld(width/cellSize, height/cellSize)
	g := game.NewGame(world, height, width, cellSize)
	ebiten.SetWindowSize(width, height)
	ebiten.SetWindowTitle("Game of life")

	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}
}
