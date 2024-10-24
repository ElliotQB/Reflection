package main

import (
	"math"
	"math/rand/v2"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func main() {
	rl.InitWindow(1280, 720, "Reflection")
	defer rl.CloseWindow()

	game := NewGame()

	game.Input = NewInput()

	game.Player = NewPlayer(float32(rl.GetScreenWidth()/4), float32(rl.GetScreenHeight()-100-50), &game)

	game.Blocks = append(game.Blocks, NewBlock(0, float32(rl.GetScreenHeight()-50), float32(rl.GetScreenWidth()), 50, &game, false))

	for i := 0; i < 50; i++ {
		blockWidth := 100
		blockHeight := 30
		game.Blocks = append(game.Blocks, NewBlock(rand.Float32()*float32((rl.GetScreenWidth()/2)-blockWidth), float32(rl.GetScreenHeight()-210-(150*i)), float32(blockWidth), float32(blockHeight), &game, FloatToBool(float32(math.Round(rand.Float64())))))
	}

	camera := rl.NewCamera2D(rl.NewVector2(0, float32(-rl.GetScreenHeight()/2)), rl.NewVector2(0, 0), 0, 1)
	cameraY := float32(0)

	rl.SetTargetFPS(60)

	for !rl.WindowShouldClose() {

		// game logic
		game.Input.UpdateInput()
		game.Player.PlayerTick()

		cameraY = float32(math.Min(0, float64(game.Player.Y-float32(rl.GetScreenHeight()))))
		camera.Target.Y = cameraY

		// drawing
		rl.BeginDrawing()
		rl.BeginMode2D(camera)
		rl.ClearBackground(rl.RayWhite)
		for i := 0; i < len(game.Blocks); i++ {
			game.Blocks[i].DrawBlock()
		}
		game.Player.DrawPlayer()
		rl.DrawRectangle(int32(rl.GetScreenWidth()/2)-int32(game.LineWidth)/2, int32(cameraY), int32(game.LineWidth), int32(rl.GetScreenHeight()), rl.Pink)
		rl.EndMode2D()
		rl.EndDrawing()
	}
}
