package main

import (
	ui "tandjgamejam/UI"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func drawFight(state *State) GameAction {
	drawBattle(state.battle)

	switch state.submode {
	case FIGHTREADY:
		startButton := ui.NewButton(float32(rl.GetScreenWidth()/2+StartButtonXOff), FightButtonY, StartButtonW, StartButtonH, "Start")
		if startButton.Draw() {
			return ActionStartFight
		}

	case FIGHTRESULT:
		shopButton := ui.NewButton(float32(rl.GetScreenWidth()/2+ShopButtonXOff), FightButtonY, ShopButtonW, ShopButtonH, "Shop")
		if shopButton.Draw() {
			return ActionGoShop
		}
		fightButton := ui.NewButton(float32(rl.GetScreenWidth()/2+FightAgainXOff), FightButtonY, FightAgainW, FightAgainH, "Fight Again")
		if fightButton.Draw() {
			return ActionFightAgain
		}
	}

	return ActionNone
}
