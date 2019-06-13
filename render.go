package main

import (
	"github.com/hajimehoshi/ebiten"
)

func renderHelper(state *gameState, screen *ebiten.Image) {
	state.maps[0].tileLayers[0].drawLayer(screen, state)
}

func (layer TileLayer) drawLayer(screen *ebiten.Image, state *gameState) error {
	for x, columns := range layer.data {
		for y, _ := range columns {
			localOffset := int(state.globalCounter / 3 % 5)
			opts := &ebiten.DrawImageOptions{}
			opts.GeoM.Translate(float64(x*layer.tileSize.X+localOffset), float64(y*layer.tileSize.Y))
			screen.DrawImage(state.player.img, opts)
		}
	}
	return nil
}

func (state *gameState) drawSprite(screen *ebiten.Image) {
	opts := &ebiten.DrawImageOptions{}
	opts.GeoM.Translate(state.player.pos.X, state.player.pos.Y)
	screen.DrawImage(state.player.img, opts)
}
