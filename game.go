package main

import rl "github.com/gen2brain/raylib-go/raylib"

type Game struct {
	Blocks []Block
	Player Player
	Input  Input
}

func NewGame() Game {
	return Game{[]Block{}, Player{}, Input{}}
}

func RectangleCollision(pos1 rl.Vector2, size1 rl.Vector2, pos2 rl.Vector2, size2 rl.Vector2) bool {
	return pos1.X > pos2.X && pos1.X+size1.X < pos2.X+size2.X && pos1.Y > pos2.Y && pos1.Y+size1.Y < pos2.Y+size2.Y
}
