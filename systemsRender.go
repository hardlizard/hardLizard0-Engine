package main

import (
	//"fmt"
	"github.com/hajimehoshi/ebiten"
)

//renderer systems
func drawSprite(screen *ebiten.Image) {
	for i := 0; i < maxUsedEntity+1; i++ {
		opts := &ebiten.DrawImageOptions{}
		opts.GeoM.Translate(float64(position[i].x), float64(position[i].y))
		screen.DrawImage(VRAM[i], opts)
	}
}
