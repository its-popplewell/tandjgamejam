package main

import rl "github.com/gen2brain/raylib-go/raylib"

func drawFight(state *State) {
	if state.submode == submodeFightRunning {
		state.battle.Update()
		if state.battle.IsOver() {
			state.player.dino.Health = max(0, state.battle.PlayerDino.Health)
			if state.player.dino.Health <= 0 {
				state.updateState(modeEnd)
				return
			}
			state.submode = submodeFightResult
		}
	}

	state.battle.Draw()

	if state.submode == submodeFightReady {
		startButton := NewButton(float32(rl.GetScreenWidth()/2-70), 260, 140, 48, "Start")
		if startButton.Draw() {
			state.submode = submodeFightRunning
		}
	}

	if state.submode == submodeFightResult {
		shopButton := NewButton(float32(rl.GetScreenWidth()/2-150), 260, 120, 48, "Shop")
		if shopButton.Draw() {
			state.updateState(modeShop)
		}

		fightButton := NewButton(float32(rl.GetScreenWidth()/2+30), 260, 150, 48, "Fight Again")
		if fightButton.Draw() {
			battle := newBattle(state.player.dino)
			state.battle = &battle
			state.submode = submodeFightReady
		}
	}
}
