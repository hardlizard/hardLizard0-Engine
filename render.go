package main

import (
	"github.com/hajimehoshi/ebiten"
)

func (state *gameState) drawSprite(screen *ebiten.Image) {
	opts := &ebiten.DrawImageOptions{}
	opts.GeoM.Translate(state.player.posX, state.player.posY)
	screen.DrawImage(state.player.img, opts)
}
