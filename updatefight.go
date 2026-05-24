package main

func updateFight(state *State) {
	switch state.submode {
	case FIGHTRUNNING:
		state.battle.Update()
		if state.battle.IsOver() {
			state.player.dino.Health = max(0, state.battle.PlayerDino.Health)
			if state.player.dino.Health <= 0 {
				state.updateState(END)
				return
			}
			state.submode = FIGHTRESULT
		}
	}
}
