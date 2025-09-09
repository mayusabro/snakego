package entities

import (
	"github.com/mayusabro/snakego/dict"
	"github.com/mayusabro/snakego/engine"
	"math/rand"
)

func smallFood(g *engine.Game) IItem {
	return newItem(dict.SMALL_FOOD, func(any any) {
		if _, ok := any.(IPlayer); ok {
			g.AddScore(10)
		}
	})
}
func bigFood(g *engine.Game) IItem {
	return newItem(dict.BIG_FOOD, func(any any) {
		if _, ok := any.(IPlayer); ok {
			g.AddScore(30)
		}
	})
}

func speedFood(g *engine.Game) IItem {
	return newItem(dict.SPEED_FOOD, func(any any) {
		if _, ok := any.(IPlayer); ok {
			g.AddScore(10)
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

func (i *Item) spawn(g *engine.Game) {
	g.World.Spawn(i, engine.Position{}.Undefined())
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
