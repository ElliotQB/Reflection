package main

import rl "github.com/gen2brain/raylib-go/raylib"

type Input struct {
	Left        bool
	Right       bool
	Jump        bool
	JumpInstant bool
}

func NewInput() Input {
	return Input{false, false, false, false}
}

func (i *Input) UpdateInput() {
	if rl.IsKeyDown(rl.KeyA) {
		i.Left = true
	} else {
		i.Left = false
	}
	if rl.IsKeyDown(rl.KeyD) {
		i.Right = true
	} else {
		i.Right = false
	}
	if rl.IsKeyDown(rl.KeySpace) {
		i.Jump = true
	} else {
		i.Jump = false
	}
	if rl.IsKeyPressed(rl.KeySpace) {
		i.JumpInstant = true
	} else {
		i.JumpInstant = false
	}
}
