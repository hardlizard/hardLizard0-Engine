package main

import (
	//"fmt"
	"github.com/hajimehoshi/ebiten"
)

//renderer systems
func drawSprite(screen *ebiten.Image) {
	for i := 0; i < maxUsedEntity+1; i++ {
		opts := &ebiten.DrawImageOptions{}
		opts.GeoM.Translate(position[i*2], position[i*2+1])
		screen.DrawImage(VRAM[i], opts)
	}
}
