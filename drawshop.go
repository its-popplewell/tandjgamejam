package main

import rl "github.com/gen2brain/raylib-go/raylib"

func drawShop(state *State) {
	rl.DrawText("SHOP", 24, 24, 32, rl.White)

	fightButton := NewButton(24, 72, 120, 42, "Fight")
	if fightButton.Draw() {
		state.updateState(modeFight)
	}
}
