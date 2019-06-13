package main

import (
	"github.com/hajimehoshi/ebiten"
)

type Entity struct {
	pos Vec
	//velocity vec     //Use this if you need velocity
	//height int //height, used to determine collisions on multiheight maps
	img *ebiten.Image
	//items
	//hold
}

func (s *Entity) movement() {
	//wasd controls
	if ebiten.IsKeyPressed(ebiten.KeyD) {
		s.pos.X = s.pos.Y + 1
	}
	if ebiten.IsKeyPressed(ebiten.KeyA) {
		s.pos.X = s.pos.X - 1
	}
	if ebiten.IsKeyPressed(ebiten.KeyS) {
		s.pos.Y = s.pos.Y + 1
	}
	if ebiten.IsKeyPressed(ebiten.KeyW) {
		s.pos.Y = s.pos.Y - 1
	}
}
