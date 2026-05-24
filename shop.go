package main

import (
	"math/rand"
)

type Commodity interface {
	buy(p *Player, s *Shop) bool
	sell(p *Player, s *Shop) bool
	getCost() int32
	getId() int32
}

type Shop struct {
	inventory []Commodity
	capacity  int32
}

func generateShop() Shop {
	var numItems int32 = ShopCapacity
	s := Shop{
		inventory: []Commodity{},
		capacity:  numItems,
	}
	s.addItemsToShop(ShopStartingItems)
	return s
}

func (s *Shop) addItemsToShop(numItems int32) {
	// TODO: add a capacity check
	for i := range numItems {
		_ = i
		comm := rand.Int31n(2)
		var toAdd Commodity
		if comm == 0 {
			toAdd = NewRandomHelmet(1)
		} else if comm == 1 {
			toAdd = NewRandomModifier(0)
		}

		s.inventory = append(s.inventory, toAdd)
	}
}

func (s *Shop) removeItemsFromShop(inp Commodity) {
	DeleteFunc(&s.inventory, func(c Commodity) bool {
		return (c.getId() == inp.getId())
	})
}

func (s *Shop) updateShop() {
}
