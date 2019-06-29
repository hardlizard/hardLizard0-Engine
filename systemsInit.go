package main

import (
	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
	//	"image/color"
	"log"
)

//init systems
func initAssets() {
	img, _, err := ebitenutil.NewImageFromFile("assets/gopherTop.png", ebiten.FilterDefault)
	if err != nil {
		log.Fatal(err)
	}
	VRAM[0] = img
}
