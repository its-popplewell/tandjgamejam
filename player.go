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
		gold: 100,
		dino: Dino{
			X: 0,
			Y: 0,
			Width: 50,
			Height: 100,
			Health: 75,
			Damage: 7,
			Block: 0,
			Speed: 0,
			Direction: 0,
			isPlayer: true,
			Color: rl.Yellow,
		},
	}
}
