package gameplay

import (
	"os"
	"time"

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

	gm.listenScore()
	gm.createGame()

	go func() {
		time.Sleep(time.Second * 2)
		entities.SpawnRandomItem(gm.game)
	}()

	return gm.game.Start()
}

func (gm *GameManager) createGame() {
	level := engine.NewLevel(engine.Size{Width: 20, Height: 20})
	level.Init()
	world := engine.NewWorld(level)
	gm.player = entities.NewPlayer()
	world.Spawn(gm.player, engine.Position{X: level.Size.Width / 2, Y: level.Size.Height / 2})
	gm.game = engine.NewGame(world)

	for range 5 {
		gm.player.AddTail(gm.game)
	}

}

func (gm *GameManager) listenScore() {
	go func() {
		for {
			gm.addScore(<-gm.game.World.ScoreListener)
		}
	}()
}

func (gm *GameManager) addScore(score int) {
	_, _, inc := gm.game.World.AddScore(gm.game, score)
	for range inc {
		gm.player.AddTail(gm.game)
		gm.player.Speed += 1
	}

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
