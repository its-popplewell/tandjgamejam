package main

func draw(state *State) {
	switch state.mode {
	case START:
		// drawStart(state)
		state.updateState(SHOP)
	case SHOP:
		drawShop(state)
	case FIGHT:
		drawFight(state)
	case END:
		drawEnd(state)
	}
}
