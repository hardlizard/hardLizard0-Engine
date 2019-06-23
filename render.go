package main

import (
	"github.com/hajimehoshi/ebiten"
)

/*
func renderHelper(s *gameState, screen *ebiten.Image) {
	s.maps[0].tileLayers[0].drawLayer(screen, state)
}

/*
func (layer TileLayer) drawLayer(screen *ebiten.Image, state *gameState) error {
	for x, columns := range layer.data {
		for y, _ := range columns {
			localOffset := int(state.globalCounter / 3 % 5)
			opts := &ebiten.DrawImageOptions{}
			opts.GeoM.Translate(float64(x*layer.tileSize.X+localOffset), float64(y*layer.tileSize.Y))
			screen.DrawImage(state.player.Entity.Thing.Img, opts)
		}
	}
	return nil
}
*/

func drawSprite(s *gameState, screen *ebiten.Image) {
	opts := &ebiten.DrawImageOptions{}
	opts.GeoM.Translate(s.player.Entity.Thing.Pos.X, s.player.Entity.Thing.Pos.Y)
	screen.DrawImage(s.player.Entity.Thing.Img, opts)
}
