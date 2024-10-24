package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

type Block struct {
	X       float32
	Y       float32
	Width   float32
	Height  float32
	Game    *Game
	Reflect bool
}

func NewBlock(x float32, y float32, width float32, height float32, game *Game, reflect bool) Block {
	return Block{x, y, width, height, game, reflect}
}

func (b *Block) DrawBlock() {
	x := b.X
	color := rl.Purple
	if b.Reflect {
		x = float32(rl.GetScreenWidth()-int(b.X)) - b.Width
		color = rl.DarkPurple
	}

	rl.DrawRectangle(int32(x), int32(b.Y), int32(b.Width), int32(b.Height), color)
}
