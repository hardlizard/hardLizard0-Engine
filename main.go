package main

import (
	//"fmt"
	"github.com/hajimehoshi/ebiten"
	"image/color"
	"log"
)

//Globals
const windowTitle = "goAA"
const scale = 1
const windowWidth = 1024
const windowHeight = 768
const assetsDir = "assets"

//Components
const maxEntity = 100    //the maximum number of entities we can handle
var maxUsedEntity uint16 //Change value if more than 65536 entities.
var freeList [maxEntity]bool

//position holds positions, ordered x, y, as pairs of entries in the array
var position [maxEntity * 2]float32

//update is the main loop of the ebiten engine. Core loop is here.
func update(screen *ebiten.Image) error {
	//gameMode 0 initializes assets
	//TODO: Insert game mode code.
	if ebiten.IsDrawingSkipped() {
		return nil
	}
	//Do stuff goes here
	screen.Fill(color.RGBA{0, 128, 128, 255})
	return nil
}

func main() {
	if err := ebiten.Run(update, windowWidth, windowHeight, scale, windowTitle); err != nil {
		log.Fatal(err)

	}
}
