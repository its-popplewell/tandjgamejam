package main

import rl "github.com/gen2brain/raylib-go/raylib"

func main() {
	rl.InitWindow(800, 600, "Dino Fight Demo")
	defer rl.CloseWindow()

	rl.SetTargetFPS(60)

	battle := newBattle()

	for !rl.WindowShouldClose() {
		battle.Update()

		rl.BeginDrawing()
		rl.ClearBackground(rl.Black)

		battle.Draw()

		rl.EndDrawing()
	}
}
