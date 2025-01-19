package main

import (
	"flag"
	"log"

	"GitHub/GameOfLife/game"
	"GitHub/GameOfLife/gameOfLife"
)

func main() {
	width := flag.Int("width", 160, "world width")
	height := flag.Int("height", 40, "world height")
	density := flag.Int("density", 5, "density max: 10")

	flag.Parse()

	if *density > 10 {
		*density = 10
	}

	world := gameOfLife.NewWorld(*width, *height, *density)
	g := game.NewGame(world, *width, *height)

	if err := g.Run(); err != nil {
		log.Fatal(err)
	}
}
