package main

type Modifier struct {
	id   int32
	cost int32
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
