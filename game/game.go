package game

import (
	"GitHub/GameOfLife/gameOfLife"
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"strings"
	"time"
)

type World interface {
	NextGeneration()
	GetCells() [][]gameOfLife.Cell
}

type Game struct {
	world     World
	width     int
	height    int
	frameRate time.Duration
}

func NewGame(world World, width int, height int) *Game {
	return &Game{
		world:     world,
		width:     width,
		height:    height,
		frameRate: time.Second / 10, // 10 FPS
	}
}

func (g *Game) Run() error {
	for {
		g.Draw()
		g.world.NextGeneration()
		time.Sleep(g.frameRate)
	}
}

func clearConsole() {
	switch runtime.GOOS {
	case "darwin", "linux":
		cmd := exec.Command("clear")
		cmd.Stdout = os.Stdout
		cmd.Run()
	case "windows":
		cmd := exec.Command("cmd", "/c", "cls")
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
}

func (g *Game) Draw() {
	clearConsole()

	var sb strings.Builder
	cells := g.world.GetCells()
	for y := range cells {
		for _, cell := range cells[y] {
			if cell.IsAlive {
				sb.WriteString("O")
			} else {
				sb.WriteString(" ")
			}
		}
		sb.WriteString("\n")
	}

	fmt.Print(sb.String())
}
