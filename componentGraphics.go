package main

import (
	"github.com/hajimehoshi/ebiten"
	"image/color"
)

//flags stores various flags for entities.

//VRAM holds pointers to ebiten images.
var VRAM [256]*ebiten.Image
var SpriteRAM [64]*ebiten.Image

func populateVRAM() {
	VRAM[0], _ = ebiten.NewImage(16, 16, ebiten.FilterDefault)
	VRAM[1], _ = ebiten.NewImageFromImage(VRAM[0], ebiten.FilterDefault)
	VRAM[2], _ = ebiten.NewImageFromImage(VRAM[0], ebiten.FilterDefault)

	VRAM[0].Fill(color.Black)
	VRAM[1].Fill(color.RGBA{255, 0, 0, 255})
	VRAM[2].Fill(color.RGBA{0, 255, 0, 255})
}

func drawLevel(screen *ebiten.Image) {
	levelSize := 16
	for i, value := range levelMap {

		opts := &ebiten.DrawImageOptions{}
		opts.GeoM.Translate(float64(i%levelSize*levelSize), float64(i/levelSize*levelSize))
		screen.DrawImage(VRAM[value], opts)
	}
}

//renderer systems
func drawSprite(screen *ebiten.Image) {
	for i := 0; i < maxUsedEntity+1; i++ {
		opts := &ebiten.DrawImageOptions{}
		opts.GeoM.Translate(float64(game.position[i].x), float64(game.position[i].y))
		screen.DrawImage(SpriteRAM[i], opts)
	}
}
