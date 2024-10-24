package main

import rl "github.com/gen2brain/raylib-go/raylib"

func main() {
	rl.InitWindow(800, 450, "Reflection")
	defer rl.CloseWindow()
	input := NewInput()
	game := NewGame()
	game.Blocks = append(game.Blocks, NewBlock(0, float32(rl.GetScreenHeight()-50), float32(rl.GetScreenWidth()), 50, &game))

	rl.SetTargetFPS(60)

	for !rl.WindowShouldClose() {
		rl.BeginDrawing()
		input.UpdateInput()
		rl.ClearBackground(rl.RayWhite)
		for i := 0; i < len(game.Blocks); i++ {
			game.Blocks[i].DrawBlock()
		}
		rl.EndDrawing()
	}
}
