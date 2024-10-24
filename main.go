package main

import rl "github.com/gen2brain/raylib-go/raylib"

func main() {
	rl.InitWindow(800, 450, "Reflection")
	defer rl.CloseWindow()
	input := NewInput()
	rl.SetTargetFPS(60)

	for !rl.WindowShouldClose() {
		rl.BeginDrawing()
		input.UpdateInput()
		rl.ClearBackground(rl.RayWhite)

		rl.EndDrawing()
	}
}
