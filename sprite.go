package main

import (
	"github.com/hajimehoshi/ebiten"
)

type Sprite struct {
	posX float64 //x position
	posY float64 //y position
	//	velocity vec  //Use this if you need velocity
	height int //height, used to determine collisions on multiheight maps
	img    *ebiten.Image
	//items
	//hold
}

func (s *Sprite) movement() {
	//wasd controls
	if ebiten.IsKeyPressed(ebiten.KeyD) {
		s.posX = s.posX + 1
	}
	if ebiten.IsKeyPressed(ebiten.KeyA) {
		s.posX = s.posX - 1
	}
	if ebiten.IsKeyPressed(ebiten.KeyS) {
		s.posY = s.posY + 1
	}
	if ebiten.IsKeyPressed(ebiten.KeyW) {
		s.posY = s.posY - 1
	}
}
