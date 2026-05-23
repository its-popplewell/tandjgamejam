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

func main() {
	// INITIALIZATION
	rl.InitWindow(800, 600, "WPC GIF Player – click‑drag to pan")
	defer rl.CloseWindow()

	// FLAGS
	rl.SetWindowState(rl.FlagWindowResizable)

	// OBV
	rl.SetTargetFPS(60)

	// var x = 0
	// var x int = 0
	var x int32 = 0
	var y int32 = 0

	// RENDERING LOOP
	for !rl.WindowShouldClose() {
		if rl.IsKeyDown(rl.KeyD) {
			x += 1
		}
		if rl.IsKeyDown(rl.KeyA) {
			x -= 1
		}
		if rl.IsKeyDown(rl.KeyW) {
			y -= 1
		}
		if rl.IsKeyDown(rl.KeyS) {
			y += 1
		}

		rl.BeginDrawing()
		rl.ClearBackground(rl.Black)

		rl.DrawRectangle(x, y, 30, 30, rl.Red)

		rl.EndDrawing()
	}
}
