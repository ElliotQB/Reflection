package main

import rl "github.com/gen2brain/raylib-go/raylib"

type Player struct {
	X    float32
	Y    float32
	Size float32
	game *Game
}

func NewPlayer(x float32, y float32, gameState *Game) Player {
	return Player{x, y, 10, gameState}
}

func (p *Player) PlayerTick() {

}

func (p *Player) PlayerCollision() {

}

func (p *Player) DrawPlayer() {
	rl.DrawRectangle(int32(p.X-(p.Size/2)), int32(p.Y-(p.Size/2)), int32(p.Size), int32(p.Size), rl.DarkBlue)
}
