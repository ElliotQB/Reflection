package main

import rl "github.com/gen2brain/raylib-go/raylib"

type Block struct {
	X      float32
	Y      float32
	Width  float32
	Height float32
}

func NewBlock(x float32, y float32, width float32, height float32) Block {
	return Block{x, y, width, height}
}

func (b *Block) DrawBlock() {
	rl.DrawRectangle(int32(b.X), int32(b.Y), int32(b.Width), int32(b.Height), rl.Pink)
}
