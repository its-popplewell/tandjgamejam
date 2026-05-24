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
	text := "YOU DIED"
	fontSize := int32(60)

	textWidth := rl.MeasureText(text, fontSize)

	x := (int32(screenWidth) - textWidth) / 2
	y := (int32(screenHeight) - fontSize) / 2

	rl.DrawText(text, x, y, fontSize, rl.Red)
	// SENDS TO START OR QUITS
}
