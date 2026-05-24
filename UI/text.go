package ui

import (
	"strings"

	rl "github.com/gen2brain/raylib-go/raylib"
)

// --- TEXT STYLES ---

type TextStyle struct {
	Color    rl.Color
	FontSize int32
}

func DefaultHeaderStyle() TextStyle {
	return TextStyle{Color: rl.White, FontSize: 32}
}

func DefaultSubheaderStyle() TextStyle {
	return TextStyle{Color: rl.LightGray, FontSize: 24}
}

func DefaultBodyStyle() TextStyle {
	return TextStyle{Color: rl.RayWhite, FontSize: 18}
}

// --- LABEL COMPONENT (For Single Lines / Headers) ---

type Label struct {
	Text  string
	X, Y  float32
	Style TextStyle
}

// Quick constructors for consistency
func NewHeader(text string, x, y float32) Label {
	return Label{Text: text, X: x, Y: y, Style: DefaultHeaderStyle()}
}

func NewSubheader(text string, x, y float32) Label {
	return Label{Text: text, X: x, Y: y, Style: DefaultSubheaderStyle()}
}

func (l *Label) Draw() {
	rl.DrawText(l.Text, int32(l.X), int32(l.Y), l.Style.FontSize, l.Style.Color)
}

// --- PARAGRAPH COMPONENT (For Multi-line text with bounds) ---

type ParagraphStyle struct {
	Text        TextStyle
	BgColor     rl.Color // Use rl.Blank for a transparent background
	BorderColor rl.Color
	BorderWidth float32
	Roundness   float32
	Padding     float32
	LineSpacing float32
}

func DefaultParagraphStyle() ParagraphStyle {
	return ParagraphStyle{
		Text:        DefaultBodyStyle(),
		BgColor:     rl.Color{R: 40, G: 44, B: 48, A: 255}, // Dark slate background
		BorderColor: rl.Color{R: 70, G: 74, B: 78, A: 255},
		BorderWidth: 2.0,
		Roundness:   0.15,
		Padding:     12.0,
		LineSpacing: 4.0,
	}
}

type Paragraph struct {
	Text  string
	Rect  rl.Rectangle
	Style ParagraphStyle
}

func NewParagraph(text string, x, y, width, height float32) Paragraph {
	return Paragraph{
		Text:  text,
		Rect:  rl.NewRectangle(x, y, width, height),
		Style: DefaultParagraphStyle(),
	}
}

// Draw handles rendering the background and automatically wrapping the text
func (p *Paragraph) Draw() {
	// 1. Draw Background (if not completely transparent)
	if p.Style.BgColor.A > 0 {
		segments := int32(16)
		rl.DrawRectangleRounded(p.Rect, p.Style.Roundness, segments, p.Style.BgColor)

		if p.Style.BorderWidth > 0 {
			rl.DrawRectangleRoundedLines(p.Rect, p.Style.Roundness, segments, p.Style.BorderColor)
		}
	}

	// 2. Word Wrap Logic
	words := strings.Split(p.Text, " ")
	var currentLine string
	yOffset := p.Rect.Y + p.Style.Padding
	maxWidth := p.Rect.Width - (p.Style.Padding * 2)

	for _, word := range words {
		testLine := currentLine
		if len(testLine) > 0 {
			testLine += " "
		}
		testLine += word

		// Measure the line with the new word added
		textWidth := float32(rl.MeasureText(testLine, p.Style.Text.FontSize))

		// If it exceeds bounds, draw the old line and start a new one
		if textWidth > maxWidth && len(currentLine) > 0 {
			rl.DrawText(currentLine, int32(p.Rect.X+p.Style.Padding), int32(yOffset), p.Style.Text.FontSize, p.Style.Text.Color)
			currentLine = word
			yOffset += float32(p.Style.Text.FontSize) + p.Style.LineSpacing
		} else {
			currentLine = testLine
		}
	}

	// Draw the remaining words on the final line
	if len(currentLine) > 0 {
		rl.DrawText(currentLine, int32(p.Rect.X+p.Style.Padding), int32(yOffset), p.Style.Text.FontSize, p.Style.Text.Color)
	}
}
