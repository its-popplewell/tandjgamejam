package main

import rl "github.com/gen2brain/raylib-go/raylib"

func drawFight(state *State) GameAction {
	drawBattle(state.battle)

	switch state.submode {
	case FIGHTREADY:
		startButton := NewButton(float32(rl.GetScreenWidth()/2-70), 260, 140, 48, "Start")
		if startButton.Draw() {
			return ActionStartFight
		}

	case FIGHTRESULT:
		shopButton := NewButton(float32(rl.GetScreenWidth()/2-150), 260, 120, 48, "Shop")
		if shopButton.Draw() {
			return ActionGoShop
		}
		fightButton := NewButton(float32(rl.GetScreenWidth()/2+30), 260, 150, 48, "Fight Again")
		if fightButton.Draw() {
			return ActionFightAgain
		}
	}

	return ActionNone
}
