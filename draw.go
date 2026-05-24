package main

func draw(state *State) GameAction {
	switch state.mode {
	case START:
		return ActionNone
	case SHOP:
		return drawShop(state)
	case FIGHT:
		return drawFight(state)
	case END:
		drawEnd(state)
	}

	return ActionNone
}
