package main

import (
	"math"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type Game struct {
	Blocks       []Block
	Player       Player
	Input        Input
	LineWidth    float32
	CurrentLevel int
	RespawnTime  float32
	GameState    float32
	DM           float32
}

func NewGame() Game {
	return Game{[]Block{}, Player{}, Input{}, 20, 0, 0, 0, 1}
}

func RectangleCollision(pos1 rl.Vector2, size1 rl.Vector2, pos2 rl.Vector2, size2 rl.Vector2) bool {
	return pos1.X+size1.X > pos2.X && pos1.X < pos2.X+size2.X && pos1.Y+size1.Y > pos2.Y && pos1.Y < pos2.Y+size2.Y
}

func CircleCollision(circle1Pos rl.Vector2, circle1Radius float32, circle2Pos rl.Vector2, circle2Radius float32) bool {
	dist := math.Sqrt(math.Pow(float64(circle2Pos.X-circle1Pos.X), 2) + math.Pow(float64(circle2Pos.Y-circle1Pos.Y), 2))
	return dist <= float64(circle1Radius)+float64(circle2Radius)
}

func CircleRectangleCollision(pos1 rl.Vector2, size1 rl.Vector2, pos2 rl.Vector2, radius2 float32) bool {
	closest := rl.NewVector2(rl.Clamp(pos2.X, pos1.X, pos1.X+size1.X), rl.Clamp(pos2.Y, pos1.Y, pos1.Y+size1.Y))
	return CircleCollision(closest, 0, pos2, radius2)
}

func BoolToInt(val bool) int {
	if val {
		return 1
	} else {
		return 0
	}
}

func FloatToBool(val float32) bool {
	if val == 0 {
		return false
	} else {
		return true
	}
}

func Sign(val float32) float32 {
	if val > 0 {
		return 1
	} else if val < 0 {
		return -1
	} else {
		return 0
	}
}

func MoveValue(val float32, dest float32, step float32) float32 {
	orig := dest-val > 0
	if dest-val > 0 {
		val += step
	} else {
		val -= step
	}
	if (dest-val > 0) != orig {
		return dest
	} else {
		return val
	}
}
