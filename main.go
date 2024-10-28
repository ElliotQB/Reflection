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

	// basic game values
	groundHeight := 50
	firstHeight := 210
	gapSize := 150
	numberBlocks := 60
	respawnTime := 45

	// create gameobjects
	game := NewGame()
	game.Input = NewInput()
	game.Player = NewPlayer(float32(rl.GetScreenWidth()/4), float32(rl.GetScreenHeight()-100-50), &game)

	// spawn blocks (and floor)
	game.SpawnBlocks(numberBlocks, float32(firstHeight), float32(gapSize), float32(groundHeight))

	// setup camera
	camera := rl.NewCamera2D(rl.NewVector2(0, float32(-rl.GetScreenHeight()/2)), rl.NewVector2(0, 0), 0, 1)
	game.CameraY = float32(0)
	game.CameraSpeed = float32(0.1)
	game.CameraMaxSpeed = float32(1000000)
	game.TweenCameraY = float32(0)
	game.CameraLowerBound = float64(rl.GetScreenHeight() / 2)

	firstHeightAbs := float32(rl.GetScreenHeight()) - float32(groundHeight) - float32(firstHeight) + float32(game.Player.Size)

	goal := NewGoal(float32(rl.GetScreenWidth()/4), firstHeightAbs-(float32(gapSize)*(float32(numberBlocks)+0.4)))

	game.LastStoodOn = &game.Blocks[0]

	rl.SetTargetFPS(60)

	for !rl.WindowShouldClose() {

		// << game logic >>
		game.DM = rl.GetFrameTime() * 60

		// state machine
		if game.GameState == 0 {
			// state 0: game running

			game.Input.UpdateInput()
			game.Player.PlayerTick()

			if game.CurrentLevel > 5 && game.Player.Y > float32(game.CameraLowerBound+float64(rl.GetScreenHeight()/2)+10) {
				game.GameState = 1
				game.RespawnTime = float32(respawnTime)
			}

		} else if game.GameState == 1 {
			// state 1: player died and waiting to respawn

			game.RespawnTime = max(0, game.RespawnTime-game.DM)
			if game.RespawnTime == 0 {
				game.GameState = 0
				game.Player.X = game.LastStoodOn.X + (game.LastStoodOn.Width / 2) - (game.Player.Size / 2)
				game.Player.Y = game.LastStoodOn.Y - (game.Player.Size)
				game.Player.Hsp = 0
				game.Player.Vsp = 0
			}
		} else if game.GameState == 2 {
			// state 2: game won and camera scrolling down

			if math.Abs(float64(game.TweenCameraY)-float64(game.CameraY)) < 1 {
				game.RespawnTime = max(0, game.RespawnTime-game.DM)
			}
			if game.RespawnTime == 0 {
				game.RespawnTime = 0
				game.SpawnBlocks(numberBlocks, float32(firstHeight), float32(gapSize), float32(groundHeight))
				game.GameState = 0
				game.LastStoodOn = &game.Blocks[0]
				game.CurrentLevel = 0
				game.CameraSpeed = 0.1
				game.CameraMaxSpeed = float32(math.Inf(1))
				game.Player = NewPlayer(float32(rl.GetScreenWidth()/4), float32(rl.GetScreenHeight()-100-50), &game)
			}
		}

		// tween camera and follow player above a certain point
		game.CameraY = float32(math.Min(game.CameraLowerBound, float64(float32(int32(game.Player.Y))-float32(rl.GetScreenHeight())+float32(rl.GetScreenHeight()))))
		game.TweenCameraY = game.TweenCameraY + rl.Clamp(game.CameraY-game.TweenCameraY, -game.CameraMaxSpeed, game.CameraMaxSpeed)*(game.CameraSpeed*game.DM)
		camera.Target.Y = game.TweenCameraY - float32(rl.GetScreenHeight())

		if game.Player.OnGround && game.Player.Y < float32(firstHeightAbs)-float32(gapSize*game.CurrentLevel) {
			game.CurrentLevel++
			if game.CurrentLevel > 5 {
				game.CameraLowerBound = float64(firstHeightAbs) - (float64(gapSize) * float64(game.CurrentLevel-1))
			}
		}

		// track the block the player was standing on last
		if game.Player.OnGround {
			ground := game.Player.PlayerInstancePlace(game.Player.X, game.Player.Y+2)

			if ground != nil && ground.Y <= game.LastStoodOn.Y {
				game.LastStoodOn = ground
			}
		}

		// illuminate platform on the side it's not displayed on when the player is standing on it
		for i := 0; i < len(game.Blocks); i++ {
			block := &game.Blocks[i]
			block.DrawReflection = false
			if game.Player.PlayerInstancePlace(game.Player.X, game.Player.Y+1) == block {
				block.DrawReflection = true
			}
		}

		// if the player collides with the goal, restart and reset the level
		if game.GameState != 2 && CircleRectangleCollision(rl.NewVector2(game.Player.X, game.Player.Y), rl.NewVector2(game.Player.Size, game.Player.Size), rl.NewVector2(goal.X, goal.Y), goal.Radius) {
			game.GameState = 2
			game.CameraLowerBound = float64(rl.GetScreenHeight() / 2)
			game.RespawnTime = 30
			game.Player.Y = game.Blocks[0].Y - game.Player.Size
			game.CameraSpeed = 0.1
			game.CameraMaxSpeed = 600
		}

		// save game
		if rl.IsKeyPressed(rl.KeyOne) {
			game.SaveGame()
		}

		// load game
		if rl.IsKeyPressed(rl.KeyTwo) {
			game.LoadGame()
		}

		// << drawing >>
		rl.BeginDrawing()
		rl.BeginMode2D(camera)
		rl.ClearBackground(rl.RayWhite)
		for i := 0; i < len(game.Blocks); i++ {
			game.Blocks[i].DrawBlock()
		}
		goal.DrawGoal()

		if game.GameState != 2 {
			game.Player.DrawPlayer()
		}
		rl.DrawRectangle(int32(rl.GetScreenWidth()/2)-int32(game.LineWidth)/2, int32(game.TweenCameraY-float32(rl.GetScreenHeight()/2)), int32(game.LineWidth), int32(rl.GetScreenHeight()), rl.Pink)
		rl.EndMode2D()

		text := strconv.Itoa(min(game.CurrentLevel, numberBlocks))
		textSize := rl.MeasureText(text, 70)
		rl.DrawText(text, 15, 15, 70, rl.DarkGray)
		rl.DrawText(text, int32(rl.GetScreenWidth())-15-textSize, 15, 70, rl.LightGray)

		rl.EndDrawing()
	}
}

// create all blocks at the start of a game
func (g *Game) SpawnBlocks(amount int, firstHeight float32, gapSize float32, groundHeight float32) {
	g.Blocks = []Block{}

	for i := 0; i < 2; i++ {
		g.Blocks = append(g.Blocks, NewBlock(0, float32(rl.GetScreenHeight()-int(groundHeight)), float32(rl.GetScreenWidth()/2), float32(groundHeight), g, i == 0))
	}

	for i := 0; i < amount; i++ {
		blockWidth := 100
		blockHeight := 30
		g.Blocks = append(g.Blocks, NewBlock(rand.Float32()*float32((rl.GetScreenWidth()/2)-blockWidth), float32(rl.GetScreenHeight()-int(firstHeight)-(int(gapSize)*i)), float32(blockWidth), float32(blockHeight), g, FloatToBool(float32(math.Round(rand.Float64())))))
	}
}
