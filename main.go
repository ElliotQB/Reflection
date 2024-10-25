package main

import (
	"math"
	"math/rand/v2"
	"strconv"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func main() {
	rl.InitWindow(1280, 720, "Reflection")
	defer rl.CloseWindow()

	groundHeight := 50
	firstHeight := 210
	gapSize := 150

	game := NewGame()

	game.Input = NewInput()

	game.Player = NewPlayer(float32(rl.GetScreenWidth()/4), float32(rl.GetScreenHeight()-100-50), &game)

	game.Blocks = append(game.Blocks, NewBlock(0, float32(rl.GetScreenHeight()-groundHeight), float32(rl.GetScreenWidth()), float32(groundHeight), &game, false))

	numberBlocks := 50

	for i := 0; i < numberBlocks; i++ {
		blockWidth := 100
		blockHeight := 30
		game.Blocks = append(game.Blocks, NewBlock(rand.Float32()*float32((rl.GetScreenWidth()/2)-blockWidth), float32(rl.GetScreenHeight()-firstHeight-(gapSize*i)), float32(blockWidth), float32(blockHeight), &game, FloatToBool(float32(math.Round(rand.Float64())))))
	}

	camera := rl.NewCamera2D(rl.NewVector2(0, float32(-rl.GetScreenHeight()/2)), rl.NewVector2(0, 0), 0, 1)
	cameraY := float32(0)
	tweenCameraY := float32(0)
	cameraLowerBound := float64(rl.GetScreenHeight() / 2)

	respawnTime := 45

	firstHeightAbs := float32(rl.GetScreenHeight()) - float32(groundHeight) - float32(firstHeight) + float32(game.Player.Size)

	goal := NewGoal(float32(rl.GetScreenWidth()/4), firstHeightAbs-(float32(gapSize)*float32(numberBlocks)))

	lastStoodOn := &game.Blocks[0]

	rl.SetTargetFPS(60)

	for !rl.WindowShouldClose() {

		// game logic
		game.DM = rl.GetFrameTime() * 60
		if game.GameState == 0 {
			game.Input.UpdateInput()
			game.Player.PlayerTick()

			if game.CurrentLevel > 5 && game.Player.Y > float32(cameraLowerBound+float64(rl.GetScreenHeight()/2)+10) {
				game.GameState = 1
				game.RespawnTime = float32(respawnTime)
			}

		} else if game.GameState == 1 {
			game.RespawnTime = max(0, game.RespawnTime-game.DM)
			if game.RespawnTime == 0 {
				game.GameState = 0
				game.Player.X = lastStoodOn.X + (lastStoodOn.Width / 2) - (game.Player.Size / 2)
				game.Player.Y = lastStoodOn.Y - (game.Player.Size)
				game.Player.Hsp = 0
				game.Player.Vsp = 0
			}
		}

		cameraY = float32(math.Min(cameraLowerBound, float64(float32(int32(game.Player.Y))-float32(rl.GetScreenHeight())+float32(rl.GetScreenHeight()))))
		tweenCameraY = tweenCameraY + (cameraY-tweenCameraY)*(0.1*game.DM)
		camera.Target.Y = tweenCameraY - float32(rl.GetScreenHeight())

		if game.Player.Y < float32(firstHeightAbs)-float32(gapSize*game.CurrentLevel) {
			game.CurrentLevel++
			if game.CurrentLevel > 5 {
				cameraLowerBound = float64(firstHeightAbs) - (float64(gapSize) * float64(game.CurrentLevel-1))
			}
		}

		if game.Player.OnGround {
			ground := game.Player.PlayerInstancePlace(game.Player.X, game.Player.Y+2)

			if ground != nil && ground.Y <= lastStoodOn.Y {
				lastStoodOn = ground
			}
		}

		// drawing
		rl.BeginDrawing()
		rl.BeginMode2D(camera)
		rl.ClearBackground(rl.RayWhite)
		for i := 0; i < len(game.Blocks); i++ {
			game.Blocks[i].DrawBlock()
		}
		goal.DrawGoal()
		game.Player.DrawPlayer()
		rl.DrawRectangle(int32(rl.GetScreenWidth()/2)-int32(game.LineWidth)/2, int32(tweenCameraY-float32(rl.GetScreenHeight()/2)), int32(game.LineWidth), int32(rl.GetScreenHeight()), rl.Pink)
		//rl.DrawCircle(int32(lastStoodOn.X), int32(lastStoodOn.Y), 10, rl.Red)
		rl.EndMode2D()
		rl.DrawText(strconv.Itoa(game.CurrentLevel), 15, 15, 20, rl.Beige)
		rl.EndDrawing()
	}
}
