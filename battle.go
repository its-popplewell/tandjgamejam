package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

type Battle struct {
	PlayerDino BattleDino
	EnemyDino  BattleDino
	HitTimer   int32
}

func NewBattle(playerDino Dino) Battle {
	screenWidth := int32(rl.GetScreenWidth())
	groundY := getGroundY()
	dinoY := groundY - dinoHeight

	PlayerDino := NewBattleDino(playerDino, screenWidth/4-dinoWidth/2, dinoY, 1, true)
	EnemyDino := NewRandomBattleDino(screenWidth*3/4-dinoWidth/2, dinoY, -1, false, rl.Blue)

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

func (battle *Battle) IsOver() bool {
	return battle.PlayerDino.Health <= 0 || battle.EnemyDino.Health <= 0
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
	b.PlayerDino.KnockbackTimer = KnockbackDuration
	b.EnemyDino.KnockbackTimer = KnockbackDuration
	b.HitTimer = HitCooldown
}

func getGroundY() int32 {
	return int32(rl.GetScreenHeight()) - 120
}
