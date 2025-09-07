package engine

import (
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
	gmChan    chan int
}

func NewGame(world *World) *Game {
	return &Game{
		World:    world,
		renderer: NewRenderer(),
		States:   NewStates(),
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

func (g *Game) loop() {
	for g.isRunning {
		g.update()
		g.render()
	}
}
func (g *Game) render() {
	g.renderer.render(func(r *Renderer) {
		r.renderGame(g)
		printMemStats(r)
	})

	time.Sleep(time.Millisecond * TIME_BUFFER)
}

func (g *Game) update() {
	g.States.update()
	g.World.update(g)
}

func (g *Game) Log(s string) {
	g.renderer.buf.WriteString(s)
	g.renderer.buf.WriteString("\r\n")
}

func (g *Game) Logf(f string, s ...any) {
	fmt.Fprintf(g.renderer.buf, f, s...)
	g.renderer.buf.WriteString("\r\n")
}

func (g *Game) GameOver() {
	if !g.World.gameOver {
		g.World.gameOver = true
		g.renderer.addMessageLine("Game Over")
		g.gmChan <- 1
	}
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
