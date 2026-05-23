package main

import (
	// "fmt";
	rl "github.com/gen2brain/raylib-go/raylib";
)

type State struct {
	mode string
	player Player
	shop Shop
}

func newState() State {
	return State{
		mode: "START",
		player: defaultPlayer(),
		shop: generateShop(),
	}
}

func (self *State) updateState(newMode string) bool {
	if (self.mode == "START") {

	} else if (self.mode == "END") {

	} else if (self.mode == "SHOP") {

	} else if (self.mode == "FIGHT") {

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

	b := NewButton(10, 10, 50, 30, "TEST")

	for !rl.WindowShouldClose() {
		// LOGIC

		// Drawing!
		rl.BeginDrawing()
		rl.ClearBackground(rl.Black)

		if game.mode == "START" {
			// display title screen
			// wait for start button selected

			game.updateState("SHOP")

			// SEND TO SHOP
		} else if game.mode == "SHOP" {
			// buy or sell
			// equpt and upgrade and whatnot

			// b.Draw()
			if b.Draw() {
				game.updateState("FIGHT")
			}

			// fmt.Printf("%+v\n", game.shop.inventory)
			// SENDS TO FIGHT
		} else if game.mode == "FIGHT" {
			// go to battle screen
			// run fight
			// if player wins fight then they can go either to shop (heal and plus hits)
			// or they go to next battle (rewards mult)
			// if win:
			// send to fight or shop (player choice)
			// if lose:
			// send to end

			rl.ClearBackground(rl.Yellow)

			// SEND TO SHOP OR FIGHT OR END
		} else if game.mode == "END" {
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
