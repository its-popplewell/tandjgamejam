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
	rl "github.com/gen2brain/raylib-go/raylib"
)


func newBattle() Battle {
	PlayerDino := newRandomDino(0, 0, 1, true, rl.Green)
	EnemyDino := newRandomDino(700, 0, -1, false, rl.Blue)

	return Battle{
		PlayerDino: PlayerDino,
		EnemyDino: EnemyDino,
		HitTimer: 0,
	}
}

func (battle *Battle) Update() {
	battle.PlayerDino.X += battle.PlayerDino.Speed * battle.PlayerDino.Direction
	battle.EnemyDino.X += battle.EnemyDino.Speed * battle.EnemyDino.Direction

	if battle.HitTimer > 0 {
		battle.HitTimer --
	}

	if rl.CheckCollisionRecs(battle.PlayerDino.Bounds(), battle.EnemyDino.Bounds()) && battle.HitTimer <=0 {
		battle.PlayerDino.Health -= battle.EnemyDino.Damage
		battle.EnemyDino.Health -= battle.PlayerDino.Damage
		battle.HitTimer = 30 // 30 frame cooldown for damage
	}
}
