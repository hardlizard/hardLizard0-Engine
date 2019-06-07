package main

import (
	//	"encoding/json"
	"fmt"
	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
	"log"
)
import "github.com/jfemory/goActionAdventure/loader"

func (state *gameState) initAssets() error {
	var err error
	//Old Player position code
	//TODO: refactor to spawn based on sprite layer

	state.player.img, _, err = ebitenutil.NewImageFromFile("assets/gopherTop.png", ebiten.FilterDefault)
	if err != nil {
		log.Fatal(err)
	}

	state.player.posX = windowWidth / 2
	state.player.posY = windowHeight / 2
	//////////////////
	//New world load code
	state.worldJSON, err = loader.LoadWorld("assets/AA.world")
	if err != nil {
		return err
	}
	if err := state.loadMaps(); err != nil {
		return err
	}
	//err = state.loadLayers()
	return err

}

//loadMaps takes a gamestate and loads the maps associated with its world.
func (state *gameState) loadMaps() error {
	var err error
	state.maps = make([]Map, 0, 0)
	for _, maps := range state.worldJSON.Maps {
		file := state.assetsDir + maps.FileName
		var newMap Map
		newMap.mapJSON, err = loader.LoadMap(file)
		newMap.offset = Offset{maps.X, maps.Y}
		fmt.Println(maps.X)
		fmt.Println(newMap)
		fmt.Println(err)
		if err != nil {
			return err
		}
		state.maps = append(state.maps, newMap)
		fmt.Println(state.maps)
		fmt.Println()
	}
	return err
}

/*
//loadLayers loads layers into a 2d slice of TileIDs from a 1d array in the JSON
func (state *gameState) loadLayers() error {
	numLayers := len(state.json.Layers)
	state.layers = make([]Layer, numLayers, numLayers)
	for i, val := range state.json.Layers {
		width := state.json.Layers[i].Width
		height := state.json.Layers[i].Height
		state.layers[i].data = make([][]TileID, width, width)
		for j := 0; j < width; j++ {
			state.layers[i].data[j] = make([]TileID, height, height)
		}
		for j := 0; j < width; j++ {
			for k := 0; k < height; k++ {
				index := (j * width) + k
				state.layers[i].data[j][k] = TileID(val.Data[index])
			}
		}
		fmt.Println(state.layers[i].data)
	}

	return nil
}


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
