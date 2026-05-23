package main

type Player struct {
	gold      int32
	inventory []Commodity
}

func (p Player) amountGold() int32 {
	return p.gold
}
