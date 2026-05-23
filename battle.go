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

type Battle struct {
	PlayerDino Dino
	EnemyDino  Dino
	HitTimer   int32
}

func newBattle() Battle {
	PlayerDino := newRandomDino(0, 400, 1, true, rl.Green)
	EnemyDino := newRandomDino(700, 400, -1, false, rl.Blue)

	return Battle{
		PlayerDino: PlayerDino,
		EnemyDino:  EnemyDino,
		HitTimer:   0,
	}
}

func (battle *Battle) Update() {
	// end battle if a dino is dead (health <= 0)
	if battle.PlayerDino.Health <= 0 || battle.EnemyDino.Health <= 0 {
		return
	}

	// update dino movement
	battle.PlayerDino.UpdateMovement()
	battle.EnemyDino.UpdateMovement()

	if battle.HitTimer > 0 {
		battle.HitTimer--
	}

	// if collision then do dino hit function
	if rl.CheckCollisionRecs(battle.PlayerDino.Bounds(), battle.EnemyDino.Bounds()) && battle.HitTimer <= 0 {
		battle.Hit()
	}
}

func (battle *Battle) Hit() {
	playerAttack := battle.PlayerDino.getAttack()
	enemyAttack := battle.EnemyDino.getAttack()

	playerDamage := playerAttack[0]
	playerBlock := playerAttack[1]

	enemyDamage := enemyAttack[0]
	enemyBlock := enemyAttack[1]

	// Calculate net damage after block
	// Damage cannot be negative, so we use max(0, damage - block)
	netPlayerDamage := max(0, enemyDamage-playerBlock)
	netEnemyDamage := max(0, playerDamage-enemyBlock)

	// Apply damage to health
	battle.PlayerDino.Health -= netPlayerDamage
	battle.EnemyDino.Health -= netEnemyDamage

	battle.resetHit()
}

func (b *Battle) resetHit() {
	b.PlayerDino.KnockbackTimer = 60
	b.EnemyDino.KnockbackTimer = 60
	b.HitTimer = 30
}

func (b Battle) Draw() {
	b.PlayerDino.Draw()
	b.EnemyDino.Draw()
}
