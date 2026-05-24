package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

type Player struct {
	gold      int32
	inventory []Commodity
	dino      Dino
}

func (p Player) amountGold() int32 {
	return p.gold
}

// func newPlayer() Player {
// 	return Player{
// 		gold: 0,
// 		dino: newRandomDino()
// 	}
// }

func defaultPlayer() Player {
	return Player{
		gold: PlayerStartGold,
		dino: Dino{
			Health:    PlayerStartHealth,
			MaxHealth: PlayerStartMaxHealth,
			Damage:    PlayerStartDamage,
			Block:     PlayerStartBlock, // by default dino has no block so they take full damage
			Color:     rl.Yellow,
		},
	}
}
