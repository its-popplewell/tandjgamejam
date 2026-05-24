package main

import (
	"fmt"
	rl "github.com/gen2brain/raylib-go/raylib"
)

func drawBattle(b *Battle) {
	screenWidth := int32(rl.GetScreenWidth())

	rl.ClearBackground(rl.SkyBlue)
	drawBackground(screenWidth)
	drawStatPanel(24, 24, "PLAYER DINO", &b.PlayerDino)
	drawStatPanel(screenWidth-244, 24, "ENEMY DINO", &b.EnemyDino)

	b.PlayerDino.Draw()
	b.EnemyDino.Draw()

	drawBattleResult(b)
}

func drawBattleResult(b *Battle) {
	if b.PlayerDino.Health > 0 && b.EnemyDino.Health > 0 {
		return
	}
	message := "PLAYER WINS"
	if b.PlayerDino.Health <= 0 {
		message = "ENEMY WINS"
	}
	screenWidth := int32(rl.GetScreenWidth())
	textWidth := rl.MeasureText(message, 36)
	rl.DrawText(message, screenWidth/2-textWidth/2, 150, 36, rl.White)
}

func drawBackground(screenWidth int32) {
	groundY := getGroundY()

	rl.DrawRectangle(0, groundY, screenWidth, 120, rl.DarkGreen)
	rl.DrawRectangle(0, groundY+28, screenWidth, 92, rl.Brown)
}

func drawStatPanel(x int32, y int32, label string, d *BattleDino) {
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

func drawPanelHealthBar(x int32, y int32, width int32, height int32, d *BattleDino) {
	healthPercentage := float32(max(0, d.Health)) / float32(d.MaxHealth)
	currentWidth := int32(float32(width) * healthPercentage)

	rl.DrawRectangle(x, y, width, height, rl.Color{R: 94, G: 38, B: 38, A: 255})
	rl.DrawRectangle(x, y, currentWidth, height, rl.Color{R: 86, G: 201, B: 106, A: 255})
	rl.DrawRectangleLines(x, y, width, height, rl.White)
}
