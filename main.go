package main

import (
	"github.com/hajimehoshi/ebiten"
	"image/color"
	"log"
)

//Globals
const windowTitle = "goAA"
const scale = 3
const windowWidth = 480
const windowHeight = 320
const assetsDir = "assets"

var game gameState
var gameMode = 0
var currentInputState inputState //TODO: Move to player singleton struct

//position holds positions, ordered x, y, as pairs of entries in the array
//var position [maxEntity * 2]float32

//update is the main loop of the ebiten engine. Core loop is here.
func update(screen *ebiten.Image) error {
	//gameMode 0 initializes assets
	if gameMode == -1 {
		panic("Game Over")
	}
	if gameMode == 0 {
		initAssets()
		initPlayer()
		newMob()
		gameMode = 1
	}

	if ebiten.IsDrawingSkipped() {
		return nil
	}
	//Do stuff goes here
	//physics
	updateInputState()
	updatePVA()
	//render
	screen.Fill(color.RGBA{0, 128, 128, 255})
	drawSprite(screen)
	death()
	//ebitenutil.DebugPrint(screen, out)
	return nil
}

func main() {
	if err := ebiten.Run(update, windowWidth, windowHeight, scale, windowTitle); err != nil {
		log.Fatal(err)

	}
}
