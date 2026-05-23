package main

type Helmet struct {
	id   int32
	cost int32
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
