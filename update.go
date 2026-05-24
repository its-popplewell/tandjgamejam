package main

func update(state *State) {
	switch state.mode {
	case FIGHT:
		updateFight(state)
	case SHOP:
		updateShop(state)
	}
}
