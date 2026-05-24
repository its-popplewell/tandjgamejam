package main

import (
	// "fmt"
	rl "github.com/gen2brain/raylib-go/raylib"
)

const (
	modeStart = "START"
	modeShop  = "SHOP"
	modeFight = "FIGHT"
	modeEnd   = "END"

	submodeBuy          = "BUY"
	submodeFightReady   = "READY"
	submodeFightRunning = "RUNNING"
	submodeFightResult  = "RESULT"
)

type State struct {
	mode    string
	submode string
	player  Player
	shop    Shop
	battle  *Battle
}

func newState() State {
	return State{
		mode:   modeStart,
		player: defaultPlayer(),
		shop:   generateShop(),
	}
}

func (self *State) updateState(newMode string) bool {
	if self.mode == modeStart {
		if newMode != modeShop {
			return false
		}
		self.submode = submodeBuy
	} else if self.mode == modeEnd {
		return false
	} else if self.mode == modeShop {
		if newMode != modeFight {
			return false
		}
		battle := newBattle(self.player.dino)
		self.battle = &battle
		self.submode = submodeFightReady
	} else if self.mode == modeFight {
		if newMode == modeEnd {
			self.battle = nil
			self.submode = ""
		} else if newMode == modeShop {
			self.battle = nil
			self.submode = submodeBuy
		} else {
			return false
		}
	}
	self.mode = newMode
	return true
}

func main() {
	// fightmain()

	game := newState()
	rl.InitWindow(800, 600, "Dino Game")
	defer rl.CloseWindow()

	rl.SetTargetFPS(60)

	for !rl.WindowShouldClose() {
		// LOGIC

		// Drawing!
		rl.BeginDrawing()
		rl.ClearBackground(rl.Black)

		if game.mode == modeStart {
			// display title screen
			// wait for start button selected

			game.updateState(modeShop)

			// SEND TO SHOP
		} else if game.mode == modeShop {
			// buy or sell
			// equpt and upgrade and whatnot

			drawShop(&game)

			// fmt.Printf("%+v\n", game.shop.inventory)
			// SENDS TO FIGHT
		} else if game.mode == modeFight {
			// go to battle screen
			// run fight
			// if player wins fight then they can go either to shop (heal and plus hits)
			// or they go to next battle (rewards mult)
			// if win:
			// send to fight or shop (player choice)
			// if lose:
			// send to end

			drawFight(&game)

			// SEND TO SHOP OR FIGHT OR END
		} else if game.mode == modeEnd {
			// Show death screen
			// Allow restart back to start
			screenWidth := rl.GetScreenWidth()
			screenHeight := rl.GetScreenHeight()

			rl.ClearBackground(rl.White)
			text := "YOU DIED"
			fontSize := int32(60)

			textWidth := rl.MeasureText(text, fontSize)

			x := (int32(screenWidth) - textWidth) / 2
			y := (int32(screenHeight) - fontSize) / 2

			rl.DrawText(text, x, y, fontSize, rl.Red)
			// SENDS TO START OR QUITS
		}

		rl.EndDrawing()
	}

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
