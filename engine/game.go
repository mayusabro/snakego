package engine

import (
	"bytes"
	"fmt"
	"time"
)

const (
	TIME_BUFFER = 16
)

type Game struct {
	isRunning bool
	renderer  *Renderer
	States    *States
	World     *World
	logger    func(string, ...any) *bytes.Buffer
	gmChan    chan int
}

func NewGame(world *World) *Game {
	return &Game{
		World:    world,
		renderer: NewRenderer(),
		States:   NewStates(),
		logger:   logger(),
	}
}

func (g *Game) Start() int {
	g.isRunning = true
	g.gmChan = make(chan int)
	go func() {
		g.loop()
	}()
	return <-g.gmChan

}

func (g *Game) AddScore(score int) {
	g.World.ScoreListener <- score
}

func (g *Game) loop() {
	for g.isRunning {
		g.update()
		g.render()
		if g.World.gameOver {
			g.isRunning = false
			g.gmChan <- g.World.Score
		}
	}
}
func (g *Game) render() {
	g.renderer.render(func(r *Renderer) {
		r.renderGame(g)
		g.log()
		printMemStats(r)
	})

	time.Sleep(time.Millisecond * TIME_BUFFER)
}

func (g *Game) update() {
	g.States.update()
	g.World.update(g)
}

func (g *Game) GameOver() {
	if !g.World.gameOver {
		g.World.gameOver = true
		g.renderer.addMessageLine("Game Over")
		g.renderer.addMessageLine("Score:%d", g.World.Score)
	}
}

//============== LOGGER

func (g *Game) log() {
	buffer := g.logger("")
	g.renderer.buf.WriteString("\r\n" + buffer.String())
	buffer.Reset()
}

func logger() func(f string, s ...any) *bytes.Buffer {
	buf := bytes.NewBufferString("")
	return func(f string, s ...any) *bytes.Buffer {
		if len(f) == 0 {
			return buf
		}
		fmt.Fprintf(buf, f, s...)
		fmt.Fprint(buf, "\r\n")
		return buf
	}
}

func (g *Game) Logf(f string, s ...any) {
	g.logger(f, s...)
}

//==============

type States struct {
	tick      int
	DeltaTime float64
	Input     byte
	start     time.Time
}

func NewStates() *States {
	return &States{
		start: time.Now(),
	}
}

func (s *States) update() {
	s.DeltaTime = float64(time.Since(s.start).Milliseconds()) / 1000.0
	s.start = time.Now()
}

func (s *States) render(r *Renderer) {
	if s.DeltaTime <= 0 {
		return
	}
	fmt.Fprintf(r.buf, "FPS %.2f\n", 1.0/s.DeltaTime)
}
