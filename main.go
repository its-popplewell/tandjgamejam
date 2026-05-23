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
	rl.InitWindow(800, 600, "WPC GIF Player – click‑drag to pan")
	defer rl.CloseWindow()
	rl.SetWindowState(rl.FlagWindowResizable)

	rl.SetTargetFPS(60)

	for !rl.WindowShouldClose() {
		rl.BeginDrawing()
		rl.ClearBackground(rl.Black)

		rl.DrawRectangle(0, 0, 30, 30, rl.Red)

		rl.EndDrawing()
	}
}
