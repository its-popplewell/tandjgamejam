package main

import (
	"math/rand"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type Dino struct {
	Health    int32
	MaxHealth int32
	Damage    int32
	Block     int32
	Color     rl.Color
	Helmet    Helmet
}

// battleDino is a copy of the persistent dino with additional fields for the battle
type BattleDino struct {
	Dino
	X              int32
	Y              int32
	Width          int32
	Height         int32
	Speed          int32 // at some point this might become a persistent stat on the dino
	Direction      int32
	isPlayer       bool
	KnockbackTimer int32 // how long the "knockback" is applied, i.e. how long they run backwards before they run forwards (in frames)
}

func NewRandomBattleDino(x int32, y int32, direction int32, isPlayer bool, color rl.Color) BattleDino {
	health := rand.Int31n(EnemyHealthRange) + EnemyMinHealth // Health between 50 and 99
	return BattleDino{
		Dino: Dino{
			Health:    health,
			MaxHealth: health,
			Damage:    rand.Int31n(EnemyDamageRange) + EnemyMinDamage, // dino base dmg
			Block:     0,                                              // block amount, 0 for dinos but helmets add some amount
			Color:     color,
		},
		X:         x,
		Y:         y,
		Width:     DinoWidth,
		Height:    DinoHeight,
		Speed:     DefaultDinoSpeed,
		Direction: direction,
		isPlayer:  isPlayer,
	}
}

func NewBattleDino(dino Dino, x int32, y int32, direction int32, isPlayer bool) BattleDino {
	return BattleDino{
		Dino:      dino,
		X:         x,
		Y:         y,
		Width:     DinoWidth,
		Height:    DinoHeight,
		Speed:     DefaultDinoSpeed,
		Direction: direction,
		isPlayer:  isPlayer,
	}
}

func (dino *BattleDino) Bounds() rl.Rectangle {
	return rl.Rectangle{
		X:      float32(dino.X),
		Y:      float32(dino.Y),
		Width:  float32(dino.Width),
		Height: float32(dino.Height),
	}
}

func (dino *BattleDino) UpdateMovement() {
	if dino.KnockbackTimer > 0 {
		dino.X -= dino.Speed * dino.Direction
		dino.KnockbackTimer--
		return
	}

	dino.X += dino.Speed * dino.Direction
}

func (dino Dino) getAttack() [2]int32 {
	outp := [2]int32{dino.Damage, dino.Block}
	helMod := dino.Helmet.getAttackModifier()
	outp[0] += helMod[0]
	outp[1] += helMod[1]

	return outp
}
