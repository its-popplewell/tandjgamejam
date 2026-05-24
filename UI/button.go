package ui

import (
	"math"

	rl "github.com/gen2brain/raylib-go/raylib"
)

// CornerToggle defines which corners should be rounded
type CornerToggle struct {
	TopLeft     bool
	TopRight    bool
	BottomRight bool
	BottomLeft  bool
}

// ButtonStyle allows you to customize the look of any button without changing game logic
type ButtonStyle struct {
	BgColor      rl.Color
	HoverBgColor rl.Color
	PressBgColor rl.Color
	TextColor    rl.Color
	BorderColor  rl.Color
	BorderWidth  float32
	CornerRadius float32      // Radius in pixels
	Corners      CornerToggle // Which specific corners to round
	FontSize     float32
	Font         rl.Font
}

// DefaultButtonStyle provides a nice, rounded gray baseline theme
func DefaultButtonStyle() ButtonStyle {
	return ButtonStyle{
		BgColor:      rl.Color{R: 200, G: 200, B: 200, A: 255},
		HoverBgColor: rl.Color{R: 230, G: 230, B: 230, A: 255},
		PressBgColor: rl.Gray,
		TextColor:    rl.DarkGray,
		BorderColor:  rl.Color{R: 150, G: 150, B: 150, A: 255},
		BorderWidth:  2.0,
		CornerRadius: 12.0,
		Corners: CornerToggle{
			TopLeft:     true,
			TopRight:    true,
			BottomRight: true,
			BottomLeft:  true,
		},
		FontSize: 20.0,
		Font:     rl.GetFontDefault(),
	}
}

// Button represents an interactive UI element
type Button struct {
	Rect  rl.Rectangle
	Label string
	Style ButtonStyle
}

// NewButton initializes a button with a default style
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
	isHovered := rl.CheckCollisionPointRec(mousePos, b.Rect)
	isPressed := isHovered && rl.IsMouseButtonDown(rl.MouseLeftButton)
	isClicked := isHovered && rl.IsMouseButtonReleased(rl.MouseLeftButton)

	// 1. Determine current background color based on state
	bgColor := b.Style.BgColor
	if isPressed {
		bgColor = b.Style.PressBgColor
	} else if isHovered {
		bgColor = b.Style.HoverBgColor
	}

	segments := int32(32)

	// 2. Draw Custom Background
	DrawRectangleCustomCorners(b.Rect, b.Style.CornerRadius, segments, b.Style.Corners, bgColor)

	// 3. Draw Custom Border
	if b.Style.BorderWidth > 0 {
		DrawRectangleCustomCornersLines(b.Rect, b.Style.CornerRadius, b.Style.BorderWidth, segments, b.Style.Corners, b.Style.BorderColor)
	}

	// 4. Draw Centered Text
	textSize := rl.MeasureTextEx(b.Style.Font, b.Label, b.Style.FontSize, 1.0)
	textX := b.Rect.X + (b.Rect.Width-textSize.X)/2
	textY := b.Rect.Y + (b.Rect.Height-textSize.Y)/2

	rl.DrawTextEx(b.Style.Font, b.Label, rl.NewVector2(textX, textY), b.Style.FontSize, 1.0, b.Style.TextColor)

	return isClicked
}

// ==========================================
// CUSTOM DRAWING HELPER FUNCTIONS
// ==========================================

// DrawRectangleCustomCorners draws a filled rectangle allowing you to smooth individual corners
func DrawRectangleCustomCorners(rect rl.Rectangle, radius float32, segments int32, corners CornerToggle, color rl.Color) {
	if radius > rect.Width/2 {
		radius = rect.Width / 2
	}
	if radius > rect.Height/2 {
		radius = rect.Height / 2
	}
	if radius <= 0 {
		rl.DrawRectangleRec(rect, color)
		return
	}

	left := rect.X + radius
	right := rect.X + rect.Width - radius
	top := rect.Y + radius
	bottom := rect.Y + rect.Height - radius

	rl.DrawRectangleV(rl.NewVector2(left, top), rl.NewVector2(rect.Width-(radius*2), rect.Height-(radius*2)), color)

	topLOffset, topROffset := float32(0), float32(0)
	if !corners.TopLeft {
		topLOffset = radius
	}
	if !corners.TopRight {
		topROffset = radius
	}
	rl.DrawRectangleV(rl.NewVector2(left-topLOffset, rect.Y), rl.NewVector2(rect.Width-(radius*2)+topLOffset+topROffset, radius), color)

	botLOffset, botROffset := float32(0), float32(0)
	if !corners.BottomLeft {
		botLOffset = radius
	}
	if !corners.BottomRight {
		botROffset = radius
	}
	rl.DrawRectangleV(rl.NewVector2(left-botLOffset, bottom), rl.NewVector2(rect.Width-(radius*2)+botLOffset+botROffset, radius), color)

	rl.DrawRectangleV(rl.NewVector2(rect.X, top), rl.NewVector2(radius, rect.Height-(radius*2)), color)
	rl.DrawRectangleV(rl.NewVector2(right, top), rl.NewVector2(radius, rect.Height-(radius*2)), color)

	if corners.TopLeft {
		rl.DrawCircleSector(rl.NewVector2(left, top), radius, 180, 270, segments, color)
	}
	if corners.TopRight {
		rl.DrawCircleSector(rl.NewVector2(right, top), radius, 270, 360, segments, color)
	}
	if corners.BottomRight {
		rl.DrawCircleSector(rl.NewVector2(right, bottom), radius, 0, 90, segments, color)
	}
	if corners.BottomLeft {
		rl.DrawCircleSector(rl.NewVector2(left, bottom), radius, 90, 180, segments, color)
	}
}

// DrawRectangleCustomCornersLines draws the outline of a rectangle with individual rounded corners
func DrawRectangleCustomCornersLines(rect rl.Rectangle, radius float32, lineThick float32, segments int32, corners CornerToggle, color rl.Color) {
	if radius > rect.Width/2 {
		radius = rect.Width / 2
	}
	if radius > rect.Height/2 {
		radius = rect.Height / 2
	}
	if radius <= 0 {
		rl.DrawRectangleLinesEx(rect, lineThick, color)
		return
	}

	left := rect.X + radius
	right := rect.X + rect.Width - radius
	top := rect.Y + radius
	bottom := rect.Y + rect.Height - radius

	rl.DrawLineEx(rl.NewVector2(left, rect.Y), rl.NewVector2(right, rect.Y), lineThick, color)
	rl.DrawLineEx(rl.NewVector2(left, rect.Y+rect.Height), rl.NewVector2(right, rect.Y+rect.Height), lineThick, color)
	rl.DrawLineEx(rl.NewVector2(rect.X, top), rl.NewVector2(rect.X, bottom), lineThick, color)
	rl.DrawLineEx(rl.NewVector2(rect.X+rect.Width, top), rl.NewVector2(rect.X+rect.Width, bottom), lineThick, color)

	if corners.TopLeft {
		drawArcLines(rl.NewVector2(left, top), radius, 180, 270, segments, lineThick, color)
	} else {
		rl.DrawLineEx(rl.NewVector2(rect.X, top), rl.NewVector2(rect.X, rect.Y), lineThick, color)
		rl.DrawLineEx(rl.NewVector2(rect.X, rect.Y), rl.NewVector2(left, rect.Y), lineThick, color)
	}

	if corners.TopRight {
		drawArcLines(rl.NewVector2(right, top), radius, 270, 360, segments, lineThick, color)
	} else {
		rl.DrawLineEx(rl.NewVector2(right, rect.Y), rl.NewVector2(rect.X+rect.Width, rect.Y), lineThick, color)
		rl.DrawLineEx(rl.NewVector2(rect.X+rect.Width, rect.Y), rl.NewVector2(rect.X+rect.Width, top), lineThick, color)
	}

	if corners.BottomRight {
		drawArcLines(rl.NewVector2(right, bottom), radius, 0, 90, segments, lineThick, color)
	} else {
		rl.DrawLineEx(rl.NewVector2(rect.X+rect.Width, bottom), rl.NewVector2(rect.X+rect.Width, rect.Y+rect.Height), lineThick, color)
		rl.DrawLineEx(rl.NewVector2(rect.X+rect.Width, rect.Y+rect.Height), rl.NewVector2(right, rect.Y+rect.Height), lineThick, color)
	}

	if corners.BottomLeft {
		drawArcLines(rl.NewVector2(left, bottom), radius, 90, 180, segments, lineThick, color)
	} else {
		rl.DrawLineEx(rl.NewVector2(left, rect.Y+rect.Height), rl.NewVector2(rect.X, rect.Y+rect.Height), lineThick, color)
		rl.DrawLineEx(rl.NewVector2(rect.X, rect.Y+rect.Height), rl.NewVector2(rect.X, bottom), lineThick, color)
	}
}

// drawArcLines connects line segments to form a smooth curve
func drawArcLines(center rl.Vector2, radius, startAngle, endAngle float32, segments int32, thickness float32, color rl.Color) {
	if segments < 4 {
		segments = 4
	}
	angleStep := (endAngle - startAngle) / float32(segments)
	degToRad := func(deg float32) float64 { return float64(deg * (math.Pi / 180.0)) }

	rad := degToRad(startAngle)
	prevX := center.X + float32(math.Cos(rad))*radius
	prevY := center.Y + float32(math.Sin(rad))*radius

	for i := int32(1); i <= segments; i++ {
		currentAngle := startAngle + float32(i)*angleStep
		rad = degToRad(currentAngle)
		currX := center.X + float32(math.Cos(rad))*radius
		currY := center.Y + float32(math.Sin(rad))*radius

		rl.DrawLineEx(rl.NewVector2(prevX, prevY), rl.NewVector2(currX, currY), thickness, color)
		prevX = currX
		prevY = currY
	}
}
