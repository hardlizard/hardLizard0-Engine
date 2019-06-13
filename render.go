package main

import (
	"github.com/hajimehoshi/ebiten"
)

func (state *gameState) drawMap(screen *ebiten.Image) {

}

func (state *gameState) drawSprite(screen *ebiten.Image) {
	opts := &ebiten.DrawImageOptions{}
	opts.GeoM.Translate(state.player.pos.X, state.player.pos.Y)
	screen.DrawImage(state.player.img, opts)
}
