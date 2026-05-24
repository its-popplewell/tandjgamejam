package main

import (
	"fmt"
	ui "tandjgamejam/UI"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func drawBattle(b *Battle) {
	screenWidth := int32(rl.GetScreenWidth())

	rl.ClearBackground(rl.SkyBlue)
	drawBackground(screenWidth)
	drawStatPanel(StatPanelMargin, StatPanelMargin, "PLAYER DINO", &b.PlayerDino)
	drawStatPanel(screenWidth-StatPanelWidth-StatPanelMargin, StatPanelMargin, "ENEMY DINO", &b.EnemyDino)

	drawBattleDino(&b.PlayerDino)
	drawBattleDino(&b.EnemyDino)

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
	textWidth := rl.MeasureText(message, BattleResultFontSize)
	rl.DrawText(message, screenWidth/2-textWidth/2, BattleResultY, BattleResultFontSize, rl.White)
}

func drawBackground(screenWidth int32) {
	groundY := getGroundY()

	rl.DrawRectangle(0, groundY, screenWidth, GroundHeight, rl.DarkGreen)
	rl.DrawRectangle(0, groundY+GroundDirtY, screenWidth, GroundDirtHeight, rl.Brown)
}

func drawBattleDino(d *BattleDino) {
	rl.DrawRectangle(d.X, d.Y, d.Width, d.Height, d.Color)
	ui.DrawHealthBar(d.X, d.Y+DinoHealthBarYOffset, d.Width, DinoHealthBarHeight, d.Health, d.MaxHealth)
}

func drawStatPanel(x int32, y int32, label string, d *BattleDino) {
	attack := d.getAttack()

	rl.DrawRectangle(x, y, StatPanelWidth, StatPanelHeight, rl.Color{R: 28, G: 32, B: 36, A: 220})
	rl.DrawRectangleLines(x, y, StatPanelWidth, StatPanelHeight, rl.White)
	rl.DrawText(label, x+StatPanelPadding, y+StatPanelTitleY, StatPanelFontSizeLarge, rl.White)
	ui.DrawHealthBar(x+StatPanelPadding, y+StatPanelHealthY, StatPanelWidth-2*StatPanelPadding, StatPanelHealthHeight, d.Health, d.MaxHealth)
	rl.DrawText(fmt.Sprintf("HP: %d / %d", max(0, d.Health), d.MaxHealth), x+StatPanelPadding, y+StatPanelHPTextY, StatPanelFontSizeSmall, rl.White)
	rl.DrawText(fmt.Sprintf("DMG: %d", attack[0]), x+StatPanelPadding, y+StatPanelStatTextY, StatPanelFontSizeSmall, rl.White)
	rl.DrawText(fmt.Sprintf("BLOCK: %d", attack[1]), x+StatPanelBlockX, y+StatPanelStatTextY, StatPanelFontSizeSmall, rl.White)
	rl.DrawText(fmt.Sprintf("SPEED: %d", d.Speed), x+StatPanelPadding, y+StatPanelSpeedTextY, StatPanelFontSizeSmall, rl.White)
}
