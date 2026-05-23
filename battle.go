package main

import (
	"fmt"
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
	screenWidth := int32(rl.GetScreenWidth())
	groundY := getGroundY()
	dinoY := groundY - dinoHeight

	PlayerDino := newRandomDino(screenWidth/4-dinoWidth/2, dinoY, 1, true, rl.Green)
	EnemyDino := newRandomDino(screenWidth*3/4-dinoWidth/2, dinoY, -1, false, rl.Blue)

	return Battle{
		PlayerDino: PlayerDino,
		EnemyDino:  EnemyDino,
		HitTimer:   0,
	}
}

func (battle *Battle) Update() {
	if rl.IsKeyPressed(rl.KeyR) || restartButtonPressed() {
		*battle = newBattle()
		return
	}

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
	screenWidth := int32(rl.GetScreenWidth())

	rl.ClearBackground(rl.SkyBlue)
	drawBackground(screenWidth)
	drawStatPanel(24, 24, "PLAYER DINO", b.PlayerDino)
	drawStatPanel(screenWidth-244, 24, "ENEMY DINO", b.EnemyDino)
	drawRestartButton()

	b.PlayerDino.Draw()
	b.EnemyDino.Draw()

	if b.PlayerDino.Health <= 0 || b.EnemyDino.Health <= 0 {
		message := "PLAYER WINS"
		if b.PlayerDino.Health <= 0 {
			message = "ENEMY WINS"
		}
		textWidth := rl.MeasureText(message, 36)
		rl.DrawText(message, screenWidth/2-textWidth/2, 150, 36, rl.White)
		rl.DrawText("Press R or click Restart", screenWidth/2-120, 192, 20, rl.White)
	}
}

func drawBackground(screenWidth int32) {
	groundY := getGroundY()

	rl.DrawRectangle(0, groundY, screenWidth, 120, rl.DarkGreen)
	rl.DrawRectangle(0, groundY+28, screenWidth, 92, rl.Brown)
}

func getGroundY() int32 {
	return int32(rl.GetScreenHeight()) - 120
}

func drawStatPanel(x int32, y int32, label string, d Dino) {
	panelWidth := int32(220)
	panelHeight := int32(130)
	attack := d.getAttack()

	rl.DrawRectangle(x, y, panelWidth, panelHeight, rl.Color{R: 28, G: 32, B: 36, A: 220})
	rl.DrawRectangleLines(x, y, panelWidth, panelHeight, rl.White)
	rl.DrawText(label, x+12, y+10, 18, rl.White)
	drawPanelHealthBar(x+12, y+38, panelWidth-24, 14, d)
	rl.DrawText(fmt.Sprintf("HP: %d / %d", max(0, d.Health), d.MaxHealth), x+12, y+60, 16, rl.White)
	rl.DrawText(fmt.Sprintf("DMG: %d", attack[0]), x+12, y+82, 16, rl.White)
	rl.DrawText(fmt.Sprintf("BLOCK: %d", attack[1]), x+110, y+82, 16, rl.White)
	rl.DrawText(fmt.Sprintf("SPEED: %d", d.Speed), x+12, y+104, 16, rl.White)
}

func drawPanelHealthBar(x int32, y int32, width int32, height int32, d Dino) {
	healthPercentage := float32(max(0, d.Health)) / float32(d.MaxHealth)
	currentWidth := int32(float32(width) * healthPercentage)

	rl.DrawRectangle(x, y, width, height, rl.Color{R: 94, G: 38, B: 38, A: 255})
	rl.DrawRectangle(x, y, currentWidth, height, rl.Color{R: 86, G: 201, B: 106, A: 255})
	rl.DrawRectangleLines(x, y, width, height, rl.White)
}

func drawRestartButton() {
	button := restartButtonBounds()
	color := rl.Color{R: 45, G: 51, B: 58, A: 255}
	if rl.CheckCollisionPointRec(rl.GetMousePosition(), button) {
		color = rl.Color{R: 65, G: 74, B: 84, A: 255}
	}

	rl.DrawRectangleRec(button, color)
	rl.DrawRectangleLines(int32(button.X), int32(button.Y), int32(button.Width), int32(button.Height), rl.White)
	rl.DrawText("Restart", int32(button.X)+18, int32(button.Y)+10, 20, rl.White)
}

func restartButtonPressed() bool {
	return rl.IsMouseButtonPressed(rl.MouseLeftButton) &&
		rl.CheckCollisionPointRec(rl.GetMousePosition(), restartButtonBounds())
}

func restartButtonBounds() rl.Rectangle {
	return rl.Rectangle{
		X:      float32(rl.GetScreenWidth()/2 - 60),
		Y:      24,
		Width:  120,
		Height: 42,
	}
}
