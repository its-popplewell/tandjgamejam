package ui

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

// PanelStyle defines the look of static background blocks
type PanelStyle struct {
	BgColor     rl.Color
	BorderColor rl.Color
	BorderWidth float32
	Roundness   float32
}

func DefaultPanelStyle() PanelStyle {
	return PanelStyle{
		BgColor:     rl.Color{R: 30, G: 34, B: 38, A: 255}, // Nice dark gray/blue
		BorderColor: rl.Color{R: 60, G: 64, B: 68, A: 255},
		BorderWidth: 2.0,
		Roundness:   0.15,
	}
}

type Panel struct {
	Rect  rl.Rectangle
	Style PanelStyle
}

func NewPanel(x, y, width, height float32) Panel {
	return Panel{
		Rect:  rl.NewRectangle(x, y, width, height),
		Style: DefaultPanelStyle(),
	}
}

// Draw renders the background panel
func (p *Panel) Draw() {
	segments := int32(16)

	// Fill
	rl.DrawRectangleRounded(p.Rect, p.Style.Roundness, segments, p.Style.BgColor)

	// Outline
	if p.Style.BorderWidth > 0 {
		rl.DrawRectangleRoundedLines(p.Rect, p.Style.Roundness, segments, p.Style.BorderColor)
	}
}
