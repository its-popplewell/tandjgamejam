package main

import (
	ui "tandjgamejam/UI"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func drawHelmetShop(h Helmet) {
}

// func drawShop(state *State) {
// 	rl.DrawText("SHOP", 24, 24, 32, rl.White)
//
// 	fightButton := NewButton(24, 72, 120, 42, "Fight")
// 	if fightButton.Draw() {
// 		state.updateState(modeFight)
// 	}
// }

func drawShop(state *State) {

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
			state.updateState(FIGHT)
		}

	case EQUIP:
		h := ui.NewHeader("HELLO WORLD", 400, 300)
		h.Draw()

		p := ui.NewPanel(10, 20, 300, 50)
		p.Draw()

		b := ui.NewButton(10, 400, 50, 20, "Click me!")
		b.Draw()
	}

	// --- Navigation Tabs ---
	buyTab := NewButton(24, 24, 100, 30, "Buy")
	equipTab := NewButton(134, 24, 100, 30, "Equip")

	if buyTab.Draw() {
		state.submode = BUY
	}
	if equipTab.Draw() {
		state.submode = EQUIP
	}
}
