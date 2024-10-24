package main

import rl "github.com/gen2brain/raylib-go/raylib"

func main() {
	rl.InitWindow(1280, 720, "Reflection")
	defer rl.CloseWindow()

	game := NewGame()

	game.Input = NewInput()

	game.Player = NewPlayer(float32(rl.GetScreenWidth()/4), float32(rl.GetScreenHeight()-100-50), &game)

	game.Blocks = append(game.Blocks, NewBlock(0, float32(rl.GetScreenHeight()-50), float32(rl.GetScreenWidth()), 50, &game, false))

	rl.SetTargetFPS(60)

	for !rl.WindowShouldClose() {
		// game logic
		game.Input.UpdateInput()
		game.Player.PlayerTick()

		// drawing
		rl.BeginDrawing()
		rl.ClearBackground(rl.RayWhite)
		for i := 0; i < len(game.Blocks); i++ {
			game.Blocks[i].DrawBlock()
		}
		game.Player.DrawPlayer()
		rl.DrawRectangle(int32(rl.GetScreenWidth()/2)-int32(game.LineWidth)/2, 0, int32(game.LineWidth), int32(rl.GetScreenHeight()), rl.Pink)
		rl.EndDrawing()
	}
}
