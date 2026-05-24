package ui

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

// ButtonStyle allows you to customize the look of any button without changing game logic
type ButtonStyle struct {
	BgColor      rl.Color
	HoverBgColor rl.Color
	TextColor    rl.Color
	BorderColor  rl.Color
	BorderWidth  float32
	Roundness    float32 // 0.0 (square) to 1.0 (pill shape)
	FontSize     int32
}

// DefaultButtonStyle provides a nice, rounded gray baseline theme
func DefaultButtonStyle() ButtonStyle {
	return ButtonStyle{
		BgColor:      rl.Color{R: 200, G: 200, B: 200, A: 255}, // Soft gray
		HoverBgColor: rl.Color{R: 230, G: 230, B: 230, A: 255}, // Lighter gray on hover
		TextColor:    rl.DarkGray,
		BorderColor:  rl.Color{R: 150, G: 150, B: 150, A: 255},
		BorderWidth:  2.0,
		Roundness:    0.3, // Nice rounded corners
		FontSize:     20,
	}
}

type Button struct {
	Rect  rl.Rectangle
	Label string
	Style ButtonStyle
}

// NewButton now initializes with a default style that can be changed later
func NewButton(x, y, width, height float32, label string) Button {
	return Button{
		Rect:  rl.NewRectangle(x, y, width, height),
		Label: label,
		Style: DefaultButtonStyle(),
	}
}

// Draw renders the button and returns true if clicked
func (b *Button) Draw() bool {
	mousePos := rl.GetMousePosition()
	hovered := rl.CheckCollisionPointRec(mousePos, b.Rect)
	clicked := hovered && rl.IsMouseButtonPressed(rl.MouseLeftButton)

	bgColor := b.Style.BgColor
	if hovered {
		bgColor = b.Style.HoverBgColor
	}

	// 16 is the number of segments used to draw the curve. Higher = smoother.
	segments := int32(16)

	// Draw Button Background
	rl.DrawRectangleRounded(b.Rect, b.Style.Roundness, segments, bgColor)

	// Draw Button Border
	if b.Style.BorderWidth > 0 {
		rl.DrawRectangleRoundedLines(b.Rect, b.Style.Roundness, segments, b.Style.BorderColor)
	}

	// Draw Centered Text
	textWidth := rl.MeasureText(b.Label, b.Style.FontSize)
	textX := int32(b.Rect.X) + (int32(b.Rect.Width)-textWidth)/2
	textY := int32(b.Rect.Y) + (int32(b.Rect.Height)-b.Style.FontSize)/2

	rl.DrawText(b.Label, textX, textY, b.Style.FontSize, b.Style.TextColor)

	return clicked
}
