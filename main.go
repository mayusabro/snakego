package main

import (
	"github.com/mayusabro/snakego/gameplay"
	"golang.org/x/term"
)

func main() {
	t, _ := term.MakeRaw(0)
	defer term.Restore(0, t)
	gm := &gameplay.GameManager{}
	gm.Init()
	gm.StartGame()
}
