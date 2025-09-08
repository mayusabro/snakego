package entities

import (
	"math/rand"
	"time"

	"github.com/mayusabro/snakego/dict"
	"github.com/mayusabro/snakego/engine"
)

func smallFood(g *engine.Game) IItem {
	return newItem(dict.SMALL_FOOD, func(any any) {
		if ip, ok := any.(IPlayer); ok {
			speedMultiplier := 1.1
			p := ip.GetPlayer()
			addSpeed := int(float64(p.Speed) * speedMultiplier)
			p.Speed += addSpeed
			g.AddScore(10)
			go func() {
				time.Sleep(time.Second * 2)
				p.Speed -= addSpeed
			}()
		}
	})
}
func bigFood(g *engine.Game) IItem {
	return newItem(dict.BIG_FOOD, func(any any) {
		if ip, ok := any.(IPlayer); ok {
			speedMultiplier := 1.2
			p := ip.GetPlayer()
			addSpeed := int(float64(p.Speed) * speedMultiplier)
			p.Speed += addSpeed
			g.AddScore(30)
			go func() {
				time.Sleep(time.Second * 2)
				p.Speed -= addSpeed
			}()
		}
	})
}

func speedFood(g *engine.Game) IItem {
	return newItem(dict.SPEED_FOOD, func(any any) {
		if ip, ok := any.(IPlayer); ok {
			speedMultiplier := 1.5
			p := ip.GetPlayer()
			addSpeed := int(float64(p.Speed) * speedMultiplier)
			p.Speed += addSpeed
			g.AddScore(10)
			go func() {
				time.Sleep(time.Second * 10)
				p.Speed -= addSpeed
			}()
		}
	})
}

//===============================

type IItem interface {
	engine.IEntity
	spawn(g *engine.Game)
	StartEffect(any)
}

type Item struct {
	engine.Entity
	effect func(affected any)
}

func newItem(id int, effect func(affected any)) IItem {
	return &Item{
		Entity: engine.Entity{
			Id: id,
		},
		effect: effect,
	}
}

func (i *Item) Update(g *engine.Game) {
}

func (i *Item) spawn(g *engine.Game) {
	l := g.World.GetCurrentLevel()
	lb := l.Bytes
	pos := getRandomPosition(l)
	for lb[pos.X][pos.Y] != dict.SPACE {
		pos = getRandomPosition(l)
	}
	g.World.Spawn(i, pos)
}

func getRandomPosition(l *engine.Level) engine.Position {
	return engine.Position{X: rand.Intn(2 + l.Size.Width - 4), Y: rand.Intn(2 + l.Size.Height - 4)}
}

func (i *Item) StartEffect(affected any) {
	i.effect(affected)
}

func SpawnRandomItem(g *engine.Game) {
	i := dict.Foods[rand.Intn(len(dict.Foods))]
	switch i {
	case dict.SMALL_FOOD:
		smallFood(g).spawn(g)
	case dict.BIG_FOOD:
		bigFood(g).spawn(g)
	case dict.SPEED_FOOD:
		speedFood(g).spawn(g)
	}
}

//==================
