package main

import rl "github.com/gen2brain/raylib-go/raylib"

type Player struct {
	X    float32
	Y    float32
	Size float32
	Game *Game
}

func NewPlayer(x float32, y float32, gameState *Game) Player {
	return Player{x, y, 10, gameState}
}

func (p *Player) PlayerTick() {

}

func (p *Player) PlayerCollision() bool {
	for i := 0; i < len(p.Game.Blocks); i++ {
		block := p.Game.Blocks[i]

		if RectangleCollision(rl.NewVector2(p.X, p.Y), rl.NewVector2(p.Size, p.Size), rl.NewVector2(block.X, block.Y), rl.NewVector2(block.Width, block.Height)) {
			return true
		}
	}
	return false
}

func (p *Player) DrawPlayer() {
	rl.DrawRectangle(int32(p.X-(p.Size/2)), int32(p.Y-(p.Size/2)), int32(p.Size), int32(p.Size), rl.DarkBlue)
}
