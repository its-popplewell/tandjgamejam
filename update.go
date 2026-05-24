package main

func update(state *State) {
	switch state.mode {
	case START:
		state.updateState(SHOP)
	case FIGHT:
		updateFight(state)
	case SHOP:
		updateShop(state)
	}
}

func handleAction(state *State, action GameAction) {
	switch action {
	case ActionOpenFight:
		state.updateState(FIGHT)
	case ActionStartFight:
		if state.mode == FIGHT && state.submode == FIGHTREADY {
			state.submode = FIGHTRUNNING
		}
	case ActionFightAgain:
		if state.mode == FIGHT && state.submode == FIGHTRESULT {
			battle := NewBattle(state.player.dino)
			state.battle = &battle
			state.submode = FIGHTREADY
		}
	case ActionGoShop:
		state.updateState(SHOP)
	case ActionShowBuy:
		if state.mode == SHOP {
			state.submode = BUY
		}
	case ActionShowEquip:
		if state.mode == SHOP {
			state.submode = EQUIP
		}
	}
}
