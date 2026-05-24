package main

import (
	"fmt"
	ui "tandjgamejam/UI"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func drawHelmetShop(h Helmet) {
}

type ShopStyle struct {
	BgColor     rl.Color
	BorderColor rl.Color
	BorderWidth float32
	Roundness   float32
	NumberItems int32
}

func DefaultShopStyle() ShopStyle {
	return ShopStyle{
		BgColor:     rl.Color{R: 30, G: 34, B: 38, A: 255},
		BorderColor: rl.Color{R: 60, G: 64, B: 68, A: 255},
		BorderWidth: 2.0,
		Roundness:   0.05,
		NumberItems: 5,
	}
}

type ShopGui struct {
	Rect  rl.Rectangle
	Style ShopStyle
}

func NewShopGui(x, y, width, height float32) ShopGui {
	return ShopGui{
		Rect:  rl.NewRectangle(x, y, width, height),
		Style: DefaultShopStyle(),
	}
}

func RoundnessToRadius(rect rl.Rectangle, roundness float32) float32 {
	if roundness < 0.0 {
		roundness = 0.0
	}
	if roundness > 1.0 {
		roundness = 1.0
	}

	shortestSide := rect.Width
	if rect.Height < rect.Width {
		shortestSide = rect.Height
	}

	return shortestSide * (roundness / 2.0)
}

func (s *ShopGui) Draw() {
	segments := int32(32)

	rl.DrawRectangleRounded(s.Rect, s.Style.Roundness, segments, s.Style.BgColor)

	ui.DrawRectangleCustomCorners(
		rl.NewRectangle(s.Rect.X, s.Rect.Y, 250.0, s.Rect.Height),
		ui.DefaultButtonStyle().CornerRadius,
		int32(32),
		ui.CornerToggle{
			TopLeft: true, TopRight: false,
			BottomLeft: true, BottomRight: false,
		},
		rl.Yellow,
	)

	var buttons [ShopCapacity]ui.Button
	var b ui.Button
	left_bar_height := s.Rect.Height / ShopCapacity

	for i := range ShopCapacity {
		b = ui.NewButton(
			s.Rect.X, s.Rect.Y+(float32(i)*left_bar_height),
			250, left_bar_height,
			fmt.Sprintf("Item %d", i+1),
		)
		b.Style.BorderWidth = 0

		switch i {
		case 0:
			b.Style.Corners = ui.CornerToggle{
				TopLeft: true, TopRight: false,
				BottomLeft: false, BottomRight: false,
			}
		case (ShopCapacity - 1):
			b.Style.Corners = ui.CornerToggle{
				TopLeft: false, TopRight: false,
				BottomLeft: true, BottomRight: false,
			}
		default:
			b.Style.Corners = ui.CornerToggle{
				TopLeft: false, TopRight: false,
				BottomLeft: false, BottomRight: false,
			}
		}

		buttons[i] = b
		if b.Draw() {
			println("BUTTON CLICKED: ", i)
		}
	}

	// b = ui.NewButton(s.Rect.X, s.Rect.Y, 250, s.Rect.Height/5, "shop")
	// b.Style.Corners = ui.CornerToggle{
	// 	TopLeft:     true,
	// 	TopRight:    false,
	// 	BottomRight: false,
	// 	BottomLeft:  false,
	// }
	// // b.Style.BorderColor = rl.Red
	// b.Style.BorderWidth = 0

	b.Draw()

	if s.Style.BorderWidth > 0 {
		rl.DrawRectangleRoundedLinesEx(s.Rect, s.Style.Roundness, segments, s.Style.BorderWidth, s.Style.BorderColor)
	}
}

func drawShop(state *State) GameAction {
	action := ActionNone

	switch state.submode {
	case BUY:
		var padding int32 = ShopBGPadding
		bgWidth := state.windowSize[0] - (padding)
		bgHeight := state.windowSize[1] - (padding)
		bg := rl.NewRectangle(
			float32(padding/2), float32(padding/2),
			float32(bgWidth), float32(bgHeight),
		)
		rl.DrawRectangleRec(bg, rl.Gray)

		rl.DrawText("SHOP", 24, 24, 32, rl.White)

		fightButton := NewButton(24, 72, 120, 42, "Fight")
		if fightButton.Draw() {
			action = ActionOpenFight
		}

	case EQUIP:

		s := NewShopGui(
			ShopBGPadding/2, ShopBGPadding/2,
			float32(state.windowSize[0]-(ShopBGPadding)),
			float32(state.windowSize[1]-40),
		)
		s.Draw()

		// h := ui.NewHeader("HELLO WORLD", 400, 300)
		// h.Draw()
		// p := ui.NewPanel(10, 20, 300, 50)
		// p.Draw()
		// b := ui.NewButton(10, 400, 50, 20, "Click me!")
		// b.Draw()

	}

	// --- Navigation Tabs ---
	buyTab := NewButton(800-136, 24, 100, 30, "Buy")
	equipTab := NewButton(800-246, 24, 100, 30, "Equip")

	if buyTab.Draw() {
		action = ActionShowBuy
	}
	if equipTab.Draw() {
		action = ActionShowEquip
	}

	return action
}
