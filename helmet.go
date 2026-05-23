package main

import (
	"math/rand"
)

type Helmet struct {
	id         int32
	cost       int32
	numHits    int32
	baseBlock  int32
	baseDamage int32
	helmetType string
	modifiers  HelmetInventory
}

func NewHelmet(id, cost, numHits, baseBlock, baseDamage, capacity int32) Helmet {
	return Helmet{
		id:         id,
		cost:       cost,
		numHits:    numHits,
		baseBlock:  baseBlock,
		baseDamage: baseDamage,
		helmetType: "default",
		modifiers:  NewHelmetInventory(capacity),
	}
}

func NewRandomHelmet(id int32) Helmet {
	capacity := rand.Int31n(16) + 15 // 15–30 (fits ~5–6 modifiers of avg size 3)
	return Helmet{
		id:         id,
		cost:       rand.Int31n(41) + 10, // 10–50
		numHits:    rand.Int31n(31) + 20, // 20–50
		baseBlock:  rand.Int31n(5) + 1,   // 1–5
		baseDamage: rand.Int31n(4),       // 0–3
		helmetType: "default",
		modifiers:  NewHelmetInventory(capacity),
	}
}

func (h Helmet) getId() int32 {
	return h.id
}

func (h Helmet) getCost() int32 {
	return h.cost
}

func (h Helmet) buy(p *Player, s *Shop) bool {
	if h.getCost() > p.amountGold() {
		return false
	}
	DeleteFunc(&s.inventory, func(c Commodity) bool {
		return (c.getId() == h.getId())
	})
	p.inventory = append(p.inventory, h)
	p.gold -= h.getCost()
	return true
}

func (h Helmet) sell(p *Player, s *Shop) bool {
	DeleteFunc(&p.inventory, func(c Commodity) bool {
		return c.getCost() == h.getCost()
	})
	s.inventory = append(s.inventory, h)
	p.gold += h.getCost()
	return true
}

func (h Helmet) getAttackModifier() [2]int32 {
	temp := h.modifiers.getAttackModifier()
	outp := [2]int32{h.baseDamage, h.baseBlock}

	outp[0] += temp[0]
	outp[1] += temp[1]

	return outp
}
