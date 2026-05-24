package main

import (
	// "fmt"
	rl "github.com/gen2brain/raylib-go/raylib"
)

type GameState string
type SubGameState string
type GameAction string

const (
	START GameState = "START"
	SHOP  GameState = "SHOP"
	FIGHT GameState = "FIGHT"
	END   GameState = "END"

	BUY          SubGameState = "BUY"
	EQUIP        SubGameState = "EQUIP"
	FIGHTREADY   SubGameState = "READY"
	FIGHTRUNNING SubGameState = "RUNNING"
	FIGHTRESULT  SubGameState = "RESULT"
	EMPTY        SubGameState = "EMPTY"

	ActionNone       GameAction = ""
	ActionOpenFight  GameAction = "OPEN_FIGHT"
	ActionStartFight GameAction = "START_FIGHT"
	ActionFightAgain GameAction = "FIGHT_AGAIN"
	ActionGoShop     GameAction = "GO_SHOP"
	ActionShowBuy    GameAction = "SHOW_BUY"
	ActionShowEquip  GameAction = "SHOW_EQUIP"
)

type Vec2 = rl.Vector2

type State struct {
	mode       GameState
	submode    SubGameState
	player     Player
	shop       Shop
	windowSize [2]int32
	battle     *Battle
}

func NewState() State {
	return State{
		mode:       START,
		submode:    EMPTY,
		player:     defaultPlayer(),
		shop:       generateShop(),
		windowSize: [2]int32{800, 600},
	}
}

func (s *State) updateState(stateNew GameState) bool {
	switch s.mode {
	case START:
		if stateNew != SHOP {
			return false
		}
		s.submode = BUY

	case END:
		return false

	case SHOP:
		if stateNew != FIGHT {
			return false
		}
		battle := NewBattle(s.player.dino)
		s.battle = &battle
		s.submode = FIGHTREADY

	case FIGHT:
		switch stateNew {
		case END:
			s.battle = nil
			s.submode = ""
		case SHOP:
			s.battle = nil
			s.submode = BUY
		default:
			return false
		}
	}

	s.mode = stateNew
	return true
}

func main() {
	// fightmain()

	game := NewState()

	rl.SetConfigFlags(rl.FlagMsaa4xHint)

	rl.InitWindow(game.windowSize[0], game.windowSize[1], "Dino Game")
	defer rl.CloseWindow()

	rl.SetTargetFPS(60)

	for !rl.WindowShouldClose() {
		update(&game)

		rl.BeginDrawing()
		rl.ClearBackground(rl.Black)
		action := draw(&game)
		rl.EndDrawing()

		handleAction(&game, action)
	}

	// for !rl.WindowShouldClose() {
	// 	// LOGIC
	//
	// 	// Drawing!
	// 	rl.BeginDrawing()
	// 	rl.ClearBackground(rl.Black)
	//
	// 	if game.mode == START {
	// 		// display title screen
	// 		// wait for start button selected
	//
	// 		game.updateState(SHOP)
	//
	// 		// SEND TO SHOP
	// 	} else if game.mode == SHOP {
	// 		// buy or sell
	// 		// equpt and upgrade and whatnot
	//
	// 		drawShop(&game)
	//
	// 		// fmt.Printf("%+v\n", game.shop.inventory)
	// 		// SENDS TO FIGHT
	// 	} else if game.mode == FIGHT {
	// 		// go to battle screen
	// 		// run fight
	// 		// if player wins fight then they can go either to shop (heal and plus hits)
	// 		// or they go to next battle (rewards mult)
	// 		// if win:
	// 		// send to fight or shop (player choice)
	// 		// if lose:
	// 		// send to end
	//
	// 		drawFight(&game)
	//
	// 		// SEND TO SHOP OR FIGHT OR END
	// 	} else if game.mode == END {
	// 		// Show death screen
	// 		// Allow restart back to start
	// 		screenWidth := game.windowSize[0]
	// 		screenHeight := game.windowSize[1]
	//
	// 		rl.ClearBackground(rl.White)
	// 		text := "YOU DIED"
	// 		fontSize := int32(60)
	//
	// 		textWidth := rl.MeasureText(text, fontSize)
	//
	// 		x := (int32(screenWidth) - textWidth) / 2
	// 		y := (int32(screenHeight) - fontSize) / 2
	//
	// 		rl.DrawText(text, x, y, fontSize, rl.Red)
	// 		// SENDS TO START OR QUITS
	// 	}
	//
	// 	rl.EndDrawing()
	// }

	// for continueGame {
	// 	if game.mode == "START" {
	// 		// display title screen
	// 		// wait for start button selected

	// 		game.updateState("SHOP")

	// 		// SEND TO SHOP
	// 	} else if game.mode == "SHOP" {
	// 		// buy or sell
	// 		// equpt and upgrade and whatnot

	// 		fmt.Printf("%+v\n", game.shop.inventory)
	// 		// SENDS TO FIGHT
	// 	} else if game.mode == "FIGHT" {
	// 		// go to battle screen
	// 		// run fight
	// 		// if player wins fight then they can go either to shop (heal and plus hits)
	// 		// or they go to next battle (rewards mult)
	// 		// if win:
	// 		// send to fight or shop (player choice)
	// 		// if lose:
	// 		// send to end

	// 		// SEND TO SHOP OR FIGHT OR END
	// 	} else if game.mode == "END" {
	// 		// Show death screen
	// 		// Allow restart back to start

	// 		// SENDS TO START OR QUITS
	// 	}
	// }
}
