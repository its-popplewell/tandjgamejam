package ui

import rl "github.com/gen2brain/raylib-go/raylib"

func DrawHealthBar(x, y, width, height int32, current, max int32) {
	if max <= 0 {
		max = 1
	}
	if current < 0 {
		current = 0
	}
	if current > max {
		current = max
	}

	currentWidth := int32(float32(width) * (float32(current) / float32(max)))

	rl.DrawRectangle(x, y, width, height, rl.Color{R: 94, G: 38, B: 38, A: 255})
	rl.DrawRectangle(x, y, currentWidth, height, rl.Color{R: 86, G: 201, B: 106, A: 255})
	rl.DrawRectangleLines(x, y, width, height, rl.White)
}
