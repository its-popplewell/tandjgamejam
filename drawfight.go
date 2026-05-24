package main

import rl "github.com/gen2brain/raylib-go/raylib"

func drawFight(state *State) {
	drawBattle(state.battle)

	switch state.submode {
	case FIGHTREADY:
		startButton := NewButton(float32(rl.GetScreenWidth()/2-70), 260, 140, 48, "Start")
		if startButton.Draw() {
			state.submode = FIGHTRUNNING // minimal state change, acceptable
		}

	case FIGHTRESULT:
		shopButton := NewButton(float32(rl.GetScreenWidth()/2-150), 260, 120, 48, "Shop")
		if shopButton.Draw() {
			state.updateState(SHOP)
		}
		fightButton := NewButton(float32(rl.GetScreenWidth()/2+30), 260, 150, 48, "Fight Again")
		if fightButton.Draw() {
			battle := NewBattle(state.player.dino)
			state.battle = &battle
			state.submode = FIGHTREADY
		}
	}
}
