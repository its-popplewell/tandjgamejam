package main

import (
	"math/rand"
)

type Modifier struct {
	id    int32
	cost  int32
	size  int32
	dmg   int32
	block int32
}

func NewModifier(id, cost, size, dmg, block int32) Modifier {
	return Modifier{
		id:    id,
		cost:  cost,
		size:  size,
		dmg:   dmg,
		block: block,
	}
}

func NewRandomModifier(id int32) Modifier {
	return Modifier{
		id:    id,
		cost:  rand.Int31n(ModifierCostRange) + ModifierMinCost, // 1–20
		size:  rand.Int31n(ModifierSizeRange) + ModifierMinSize, // 1–5
		dmg:   rand.Int31n(ModifierDamageRange),                 // 0–9
		block: rand.Int31n(ModifierBlockRange),                  // 0–9
	}
}

func (m Modifier) getId() int32 {
	return m.id
}

func (m Modifier) getCost() int32 {
	return m.cost
}

func (m Modifier) buy(p *Player, s *Shop) bool {
	if m.getCost() > p.amountGold() {
		return false
	}
	DeleteFunc(&s.inventory, func(c Commodity) bool {
		return (c.getId() == m.getId())
	})
	p.inventory = append(p.inventory, m)
	p.gold -= m.getCost()
	return true
}

func (m Modifier) sell(p *Player, s *Shop) bool {
	DeleteFunc(&p.inventory, func(c Commodity) bool {
		return c.getCost() == m.getCost()
	})
	s.inventory = append(s.inventory, m)
	p.gold += m.getCost()
	return true
}

func (m Modifier) getAttackModifier() [2]int32 {
	return [2]int32{m.dmg, m.block}
}
