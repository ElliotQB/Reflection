package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

type Block struct {
	X              float32
	Y              float32
	Width          float32
	Height         float32
	Game           *Game
	Reflect        bool
	DrawReflection bool
}

func NewBlock(x float32, y float32, width float32, height float32, game *Game, reflect bool) Block {
	return Block{x, y, width, height, game, reflect, false}
}

func (b *Block) DrawBlock() {
	x := b.X
	reflectX := float32(rl.GetScreenWidth()-int(b.X)) - b.Width
	color := rl.DarkPurple
	if b.Reflect {
		x = reflectX
		color = rl.Purple
	}

	rl.DrawRectangle(int32(x), int32(b.Y), int32(b.Width), int32(b.Height), color)

	if b.DrawReflection {
		if b.Reflect {
			rl.DrawRectangle(int32(b.X), int32(b.Y), int32(b.Width), int32(b.Height), rl.NewColor(rl.LightGray.R, rl.LightGray.G, rl.LightGray.B, 140))
		} else {
			rl.DrawRectangle(int32(reflectX), int32(b.Y), int32(b.Width), int32(b.Height), rl.NewColor(rl.LightGray.R, rl.LightGray.G, rl.LightGray.B, 60))
		}
	}
}
