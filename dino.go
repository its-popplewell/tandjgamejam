package main

import (
	// "fmt"
	// "io"
	// "log"
	// "net/http"
	// "os"
	// "path/filepath"
	// "sort"
	// "time"
	"math/rand"

	rl "github.com/gen2brain/raylib-go/raylib"
)

const dinoWidth int32 = 50
const dinoHeight int32 = 100

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

func newRandomBattleDino(x int32, y int32, direction int32, isPlayer bool, color rl.Color) BattleDino {
	health := rand.Int31n(50) + 50 // Health between 50 and 99
	return BattleDino{
		Dino: Dino{
			Health:    health,
			MaxHealth: health,
			Damage:    rand.Int31n(5) + 5, // dino base dmg
			Block:     0,                  // block amount, 0 for dinos but helmets add some amount
			Color:     color,
		},
		X:         x,
		Y:         y,
		Width:     dinoWidth,
		Height:    dinoHeight,
		Speed:     2,
		Direction: direction,
		isPlayer:  isPlayer,
	}
}

func newBattleDino(dino Dino, x int32, y int32, direction int32, isPlayer bool) BattleDino {
	return BattleDino{
		Dino:      dino,
		X:         x,
		Y:         y,
		Width:     dinoWidth,
		Height:    dinoHeight,
		Speed:     2,
		Direction: direction,
		isPlayer:  isPlayer,
	}
}

func (dino BattleDino) Bounds() rl.Rectangle {
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

func (d BattleDino) Draw() {
	rl.DrawRectangle(d.X, d.Y, d.Width, d.Height, d.Color)
	// Draw health bar
	healthBarWidth := d.Width
	healthBarHeight := int32(10)
	healthPercentage := float32(max(0, d.Health)) / float32(d.MaxHealth)
	currentHealthBarWidth := int32(float32(healthBarWidth) * healthPercentage)

	rl.DrawRectangle(d.X, d.Y-20, healthBarWidth, healthBarHeight, rl.Red)
	rl.DrawRectangle(d.X, d.Y-20, currentHealthBarWidth, healthBarHeight, rl.Green)
}
