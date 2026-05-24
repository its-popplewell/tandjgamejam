package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

func drawEnd(s *State) {
	// Show death screen
	// Allow restart back to start
	screenWidth := s.windowSize[0]
	screenHeight := s.windowSize[1]

	rl.ClearBackground(rl.White)
	textWidth := rl.MeasureText(DeathText, DeathScreenFontSize)

	x := (int32(screenWidth) - textWidth) / 2
	y := (int32(screenHeight) - DeathScreenFontSize) / 2

	rl.DrawText(DeathText, x, y, DeathScreenFontSize, rl.Red)
	// SENDS TO START OR QUITS
}
