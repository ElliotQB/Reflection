package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

type Goal struct {
	X      float32
	Y      float32
	Radius float32
}

func NewGoal(x float32, y float32, radius float32) Goal {
	return Goal{x, y, 150}
}

func (g *Goal) DrawGoal() {
	x := float32(rl.GetScreenWidth()) - g.X
	rl.DrawCircle(int32(x), int32(g.Y), g.Radius, rl.DarkBlue)
	rl.DrawCircle(int32(g.X), int32(g.Y), g.Radius, rl.SkyBlue)
}
