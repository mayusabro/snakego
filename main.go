package main

import (
	"github.com/mayusabro/snakego/engine"
	"github.com/mayusabro/snakego/gameplay/entities"
	"golang.org/x/term"
)

func main() {
	t, _ := term.MakeRaw(0)
	defer term.Restore(0, t)
	level := engine.NewLevel(engine.Size{Width: 40, Height: 20})
	level.Init()
	world := engine.NewWorld(level)
	player := entities.NewPlayer()

	world.Spawn(player, engine.Position{X: 10, Y: 10})
	game := engine.NewGame(world)
	for range 100 {
		player.AddTail(game)
	}
	game.Start()

}
