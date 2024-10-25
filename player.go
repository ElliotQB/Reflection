package main

import (
	"math"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type Player struct {
	X             float32
	Y             float32
	Hsp           float32
	Vsp           float32
	MaxHsp        float32
	AccelXGround  float32
	DecelXGround  float32
	AccelXAir     float32
	DecelXAir     float32
	Grv           float32
	OnGround      bool
	OnGroundCT    bool
	CTTime        float32
	CTCurrentTime float32
	JumpStrength  float32
	JumpBuffer    float32
	SetJumpBuffer float32
	Size          float32
	Game          *Game
}

func NewPlayer(x float32, y float32, gameState *Game) Player {
	grv := float32(0.35)
	size := float32(50)
	jumpstrength := float32(-12)
	accelXGround := float32(1)
	decelXGround := float32(1)
	accelXAir := float32(0.4)
	decelXAir := float32(0.2)
	maxHsp := float32(7)
	jumpBuffer := float32(10)
	coyoteTime := float32(8)

	return Player{x, y, 0, 0, maxHsp, accelXGround, decelXGround, accelXAir, decelXAir, grv, false, false, coyoteTime, 0, jumpstrength, 0, jumpBuffer, size, gameState}
}

func (p *Player) PlayerTick() {

	// check if the player is on the ground
	p.OnGround = p.PlayerCollision(p.X, p.Y+1)

	// coyote time
	if p.OnGround {
		p.OnGroundCT = true
		p.CTCurrentTime = p.CTTime
	}
	if p.CTCurrentTime == 0 && !p.OnGround {
		p.OnGroundCT = false
	}
	p.CTCurrentTime = max(0, p.CTCurrentTime-p.Game.DM)

	// get input for horizontal movement
	moveX := BoolToInt(p.Game.Input.Right) - BoolToInt(p.Game.Input.Left)

	// clean accel/decel variables
	accelX := p.AccelXAir
	decelX := p.DecelXAir
	if p.OnGround {
		accelX = p.AccelXGround
		decelX = p.DecelXGround
	}

	// accellerate or decelerate player based on what direction they're pressing
	if moveX == 1 {
		p.Hsp = min(p.MaxHsp, p.Hsp+(accelX*p.Game.DM))
	} else if moveX == -1 {
		p.Hsp = max(-p.MaxHsp, p.Hsp-(accelX*p.Game.DM))
	} else {
		p.Hsp = MoveValue(p.Hsp, 0, (decelX * p.Game.DM))
	}

	// pull the player down with gravity
	p.Vsp += (p.Grv * p.Game.DM)

	// jump
	p.JumpBuffer = max(0, p.JumpBuffer-p.Game.DM)
	if p.Game.Input.JumpInstant {
		p.JumpBuffer = p.SetJumpBuffer
	}
	if p.JumpBuffer > 0 && p.OnGroundCT {

		p.Vsp = p.JumpStrength
		p.OnGroundCT = false
		p.OnGround = false
	}

	// horizontal collision
	if p.PlayerCollision(p.X+(p.Hsp*p.Game.DM), p.Y) {
		p.X = float32(math.Round(float64(p.X)))
		for !p.PlayerCollision(p.X+Sign(p.Hsp), p.Y) {
			p.X += Sign(p.Hsp)
		}
		p.Hsp = 0
	}

	// vertical collision
	if p.PlayerCollision(p.X, p.Y+(p.Vsp*p.Game.DM)) {
		p.Y = float32(math.Round(float64(p.Y)))
		for !p.PlayerCollision(p.X, p.Y+Sign(p.Vsp)) {
			p.Y += Sign(p.Vsp)
		}
		p.Vsp = 0
	}

	// unstuck
	for p.PlayerCollision(p.X, p.Y) && p.Hsp == 0 && p.Vsp == 0 {
		p.Y--
	}

	// clamp player inside screen
	if p.X < 0 {
		p.X = 0
		p.Hsp = 0
	}
	if p.X > float32(rl.GetScreenWidth()/2)-(p.Game.LineWidth/2)-p.Size {
		p.X = float32(rl.GetScreenWidth()/2) - (p.Game.LineWidth / 2) - p.Size
		p.Hsp = 0
	}

	// apply speeds
	p.X += p.Hsp * p.Game.DM
	p.Y += p.Vsp * p.Game.DM
}

func (p *Player) PlayerCollision(x float32, y float32) bool {
	for i := 0; i < len(p.Game.Blocks); i++ {
		block := p.Game.Blocks[i]

		if RectangleCollision(rl.NewVector2(x, y), rl.NewVector2(p.Size, p.Size), rl.NewVector2(block.X, block.Y), rl.NewVector2(block.Width, block.Height)) {
			return true
		}
	}
	return false
}

func (p *Player) PlayerInstancePlace(x float32, y float32) *Block {
	for i := 0; i < len(p.Game.Blocks); i++ {
		block := &p.Game.Blocks[i]

		if RectangleCollision(rl.NewVector2(x, y), rl.NewVector2(p.Size, p.Size), rl.NewVector2(block.X, block.Y), rl.NewVector2(block.Width, block.Height)) {
			return block
		}
	}
	return nil
}

func (p *Player) DrawPlayer() {
	x := float32(rl.GetScreenWidth()) - p.X - p.Size
	rl.DrawRectangle(int32(x), int32(p.Y), int32(p.Size), int32(p.Size), rl.DarkBlue)
	rl.DrawRectangle(int32(p.X), int32(p.Y), int32(p.Size), int32(p.Size), rl.SkyBlue)
}
