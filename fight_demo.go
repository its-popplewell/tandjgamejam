package main

import rl "github.com/gen2brain/raylib-go/raylib"

func fightmain() {
	rl.InitWindow(WindowWidth, WindowHeight, FightDemoWindowTitle)
	defer rl.CloseWindow()

	rl.SetTargetFPS(TargetFPS)

	battle := NewBattle(defaultPlayer().dino)

	for !rl.WindowShouldClose() {
		battle.Update()

		rl.BeginDrawing()
		rl.ClearBackground(rl.Black)

		// battle.Draw()
		drawBattle(&battle)

		rl.EndDrawing()
	}
}
