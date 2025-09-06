package engine

import (
	"fmt"
	"os"
	"time"
)

const (
	TIME_BUFFER = 32
)

type Game struct {
	isRunning bool
	renderer  *Renderer
	States    *States
	World     *World
}

func NewGame(world *World) *Game {
	return &Game{
		World:    world,
		renderer: NewRenderer(),
		States:   NewStates(),
	}
}

func (g *Game) Start() {
	g.isRunning = true
	g.loop()
}

func (g *Game) loop() {
	for g.isRunning {
		g.update()
		g.render()
	}
}
func (g *Game) render() {
	g.renderer.render(func(r *Renderer) {
		g.World.render(g)
		g.States.render(r)
		printMemStats(r)
	})

	time.Sleep(time.Millisecond * TIME_BUFFER)
}

func (g *Game) update() {
	g.States.update()
	g.World.update()
}

func (g *Game) Log(s string) {
	g.renderer.buf.WriteString(s)
	g.renderer.buf.WriteString("\r\n")
}

func (g *Game) Logf(f string, s ...any) {
	fmt.Fprintf(g.renderer.buf, f, s)
	g.renderer.buf.WriteString("\r\n")
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
	s.readInput()
}

func (s *States) readInput() {
	go func() {
		b := make([]byte, 1)
		os.Stdin.Read(b)
		s.Input = b[0]
	}()
}

func (s *States) render(r *Renderer) {
	if s.DeltaTime <= 0 {
		return
	}
	fmt.Fprintf(r.buf, "FPS %.2f\n", 1000.0/float64(s.DeltaTime)/1000.0)
}
