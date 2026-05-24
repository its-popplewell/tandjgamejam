package main

import rl "github.com/gen2brain/raylib-go/raylib"

func drawHelmetShop(h Helmet) {

}

func drawShop(state *State) {
	switch state.submode {
	case "BUY":
		var padding int32 = 40
		bgWidth := state.windowSize[0] - (padding)
		bgHeight := state.windowSize[1] - (padding)
		bg := rl.NewRectangle(
			float32(padding/2), float32(padding/2),
			float32(bgWidth), float32(bgHeight),
		)
		rl.DrawRectangleRec(bg, rl.Gray)

	case "EQUIP":
	}
}
