package main

type Commodity interface {
	buy(p *Player, s *Shop) bool
	sell(p *Player, s *Shop) bool
	getCost() int32
	getId() int32
}

type Shop struct {
	inventory []Commodity
}
