package main

import (
	"fmt"
	"github.com/hajimehoshi/ebiten"
	//"github.com/hajimehoshi/ebiten/ebitenutil"
	"github.com/jfemory/goActionAdventure/game"
	//"github.com/jfemory/goActionAdventure/renderer"
	"image/color"
	"log"
)

/////////////////////Globals//////////////////////////////////
const windowTitle = "goAA"
const scale = 1
const windowWidth = 1024
const windowHeight = 768

var w *game.World

const assetsDir = "assets"

//////////////////////////////////////////////////////////////

//update is the main loop of the ebiten engine. Core loop is here.
func update(screen *ebiten.Image) error {
	//gameMode 0 initializes assets
	if w.GetMode() == 0 {
		w.SetMode(1)
	}
	if ebiten.IsDrawingSkipped() {
		return nil
	}
	//Do stuff goes here
	screen.Fill(color.RGBA{0, 128, 128, 255})
	return nil
}

func main() {
	var err error
	w, err = game.NewWorld("assets")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(w)
	if err := ebiten.Run(update, windowWidth, windowHeight, scale, windowTitle); err != nil {
		log.Fatal(err)

	}
}
