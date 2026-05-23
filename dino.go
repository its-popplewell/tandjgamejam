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
	X	int32
	Y 	int32
	Width int32
	Height int32
	Health int32
	Damage int32
	Speed int32
	Direction int32
	isPlayer bool
	Color rl.Color
}

type Battle struct {
	PlayerDino Dino
	EnemyDino Dino
	HitTimer int32
}

func newRandomDino(x int32, y int32, direction int32, isPlayer bool, color rl.Color) Dino {
	return Dino{
		X: x,
		Y: y,
		Width: dinoWidth,
		Height: dinoHeight,
		Health: rand.Int31n(50) + 50, // Health between 50 and 99
		Damage: rand.Int31n(5) + 5, // Damage between 5 and 10
		Speed: 5,
		Direction: direction,
		isPlayer: isPlayer,
	}
}

func (dino Dino) Bounds() rl.Rectangle {
	return rl.Rectangle{
		X: 		float32(dino.X),
		Y: 		float32(dino.Y),
		Width: 	float32(dino.Width),
		Height: float32(dino.Height),
	}
}