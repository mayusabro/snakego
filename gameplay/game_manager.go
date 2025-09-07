package gameplay

import (
	"os"

	"github.com/mayusabro/snakego/engine"
	"github.com/mayusabro/snakego/gameplay/entities"
)

type GameManager struct {
	channel     chan bool
	initialized bool
	game        *engine.Game
	player      *entities.Player
}

func (gm *GameManager) Init() {
	if gm.initialized {
		return
	}
	gm.initialized = true
	gm.ReadInput()
}

func (gm *GameManager) StartGame() int {
	if !gm.initialized {
		println("Game not initialized")
		return -1
	}
	level := engine.NewLevel(engine.Size{Width: 40, Height: 20})
	level.Init()
	world := engine.NewWorld(level)
	gm.player = entities.NewPlayer()
	world.Spawn(gm.player, engine.Position{X: 10, Y: 10})
	gm.game = engine.NewGame(world)
	for range 10 {
		gm.player.AddTail(gm.game)
	}
	return gm.game.Start()
}

func (gm *GameManager) ReadInput() {
	go func() {
		for {
			b := make([]byte, 1)
			os.Stdin.Read(b)
			gm.game.States.Input = b[0]
		}
	}()
}
