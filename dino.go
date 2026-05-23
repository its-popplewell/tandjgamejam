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
	X              int32
	Y              int32
	Width          int32
	Height         int32
	Health         int32
	Damage         int32
	Block          int32
	Speed          int32
	Direction      int32
	isPlayer       bool
	Color          rl.Color
	KnockbackTimer int32
	Helmet         Helmet
}

func newRandomDino(x int32, y int32, direction int32, isPlayer bool, color rl.Color) Dino {
	return Dino{
		X:         x,
		Y:         y,
		Width:     dinoWidth,
		Height:    dinoHeight,
		Health:    rand.Int31n(50) + 50, // Health between 50 and 99
		Damage:    rand.Int31n(5) + 5,   // dino base dmg
		Block:     0,                    // block amount, 0 for dinos but helmets add some amount
		Speed:     1,
		Direction: direction,
		isPlayer:  isPlayer,
		Color:     color,
	}
}

func (dino Dino) Bounds() rl.Rectangle {
	return rl.Rectangle{
		X:      float32(dino.X),
		Y:      float32(dino.Y),
		Width:  float32(dino.Width),
		Height: float32(dino.Height),
	}
}

func (dino *Dino) UpdateMovement() {
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

func (d Dino) Draw() {
	rl.DrawRectangle(d.X, d.Y, d.Width, d.Height, d.Color)
	// Draw health bar
	healthBarWidth := d.Width
	healthBarHeight := int32(10)
	healthPercentage := float32(d.Health) / 100.0
	currentHealthBarWidth := int32(float32(healthBarWidth) * healthPercentage)

	rl.DrawRectangle(d.X, d.Y-20, healthBarWidth, healthBarHeight, rl.Red)
	rl.DrawRectangle(d.X, d.Y-20, currentHealthBarWidth, healthBarHeight, rl.Green)
}
