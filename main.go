package main

import (
	//	"encoding/json"
	//"fmt"
	"github.com/hajimehoshi/ebiten"
	//"github.com/hajimehoshi/ebiten/ebitenutil"
	"image/color"
	"log"
)

/////////////////////Globals//////////////////////////////////
const windowTitle = "goAA"
const scale = 1
const windowWidth = 1024
const windowHeight = 768

var gameMode = 0

//////////////////////////////////////////////////////////////

//update is the main loop of the ebiten engine. Core loop is here.
func (state *gameState) update(screen *ebiten.Image) error {
	//gameMode 0 initializes assets
	if state.gameMode == 0 {
		if err := state.initAssets(); err != nil {
			log.Fatal(err)
		}
		state.gameMode = 1
	}
	if ebiten.IsDrawingSkipped() {
		return nil
	}
	//Do stuff goes here
	screen.Fill(color.RGBA{255, 0, 0, 255})
	state.player.movement()
	state.drawSprite(screen)
	return nil
}

/*
//loadAtlas loads pngs of tiles into the atlas
func (state *gameState) loadAtlas() error {
	var err error
	state.atlas.VRAM = make([]*ebiten.Image, 0, 0)

	return err
}


//loadTileSet loads a tileset into VRAM with a blank 0 tile of the dimensions of the first tileset.
func (state *gameState) loadTileSet() error {
	//initialize blank tile, index 0
	err := state.loadBlankTile(0) //argument is the size of the tiles in the 0th tileset.
	if err != nil {
		return err
	}
	for i, set := range state.json.Tilesets {
		for j, val := range set.Tiles {

		}

	}
	return err
}


//loadBlankTile loads a blank tile of given dimensions into VRAM at the current end of the slice.
func (state *gameState) loadBlankTile(i int) error {
	tileWidth := state.json.Tilesets[i].Tilewidth
	tileHeight := state.json.Tilesets[i].Tileheight
	blank, err := ebiten.NewImage(tileWidth, tileHeight, ebiten.FilterDefault)
	blank.Fill(color.RGBA{0, 0, 0, 0})
	state.atlas.VRAM = append(state.atlas.VRAM, blank)
	return nil
}

*/

func main() {
	var state gameState
	err := state.initGameState()
	if err != nil {
		log.Fatal("failed to initialize game state. ")
	}
	if err := ebiten.Run(state.update, windowWidth, windowHeight, scale, windowTitle); err != nil {
		log.Fatal(err)
	}
}
