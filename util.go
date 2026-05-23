package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

func FindFunc[T any](slice []T, match func(T) bool) int {
	for i, v := range slice {
		if match(v) {
			return i
		}
	}
	return -1
}

func DeleteFunc[T any](slice *[]T, match func(T) bool) {
	i := FindFunc(*slice, match)
	if i == -1 {
		return
	}
	*slice = append((*slice)[:i], (*slice)[i+1:]...)
}

type Button struct {
    rect  rl.Rectangle
    label string
    color rl.Color
}

func NewButton(x, y, width, height float32, label string) Button {
    return Button{
        rect:  rl.NewRectangle(x, y, width, height),
        label: label,
        color: rl.Gray,
    }
}

// Returns true if clicked
func (b *Button) Draw() bool {
    mousePos := rl.GetMousePosition()
    hovered := rl.CheckCollisionPointRec(mousePos, b.rect)
    clicked := hovered && rl.IsMouseButtonPressed(rl.MouseLeftButton)

    // Change color on hover
    color := b.color
    if hovered {
        color = rl.LightGray
    }

    rl.DrawRectangleRec(b.rect, color)
    rl.DrawRectangleLinesEx(b.rect, 2, rl.DarkGray) // border

    // Center the text
    fontSize := int32(20)
    textWidth := rl.MeasureText(b.label, fontSize)
    textX := int32(b.rect.X) + (int32(b.rect.Width)-textWidth)/2
    textY := int32(b.rect.Y) + (int32(b.rect.Height)-fontSize)/2
    rl.DrawText(b.label, textX, textY, fontSize, rl.Black)

    return clicked
}
