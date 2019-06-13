package main

import (
	//	"encoding/json"
	"fmt"
	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
	"log"
)
import "github.com/jfemory/goActionAdventure/jsonLoader"

func (state *gameState) initAssets() error {
	var err error
	//Old Player position code
	//TODO: refactor to spawn based on sprite layer

	state.player.img, _, err = ebitenutil.NewImageFromFile("assets/gopherTop.png", ebiten.FilterDefault)
	if err != nil {
		log.Fatal(err)
	}
	//TODO: Brittle code, replace w/ spawnpoint function
	state.player.pos.X = windowWidth / 2
	state.player.pos.Y = windowHeight / 2
	//////////////////
	//New world load code
	state.worldJSON, err = jsonLoader.LoadWorld("assets/AA.world")
	if err != nil {
		return err
	}
	if err := state.loadMaps(); err != nil {
		return err
	}
	//err = state.loadLayers()
	return err

}

//TODO: Move state.assetsDir to a env var or global.
//initGameState initializes state so the game can load.
func (state *gameState) initGameState() error {
	var err error
	state.gameMode = 0
	state.assetsDir = "assets/"
	return err
}

//loadMaps takes a gamestate and loads the maps associated with its world.
func (state *gameState) loadMaps() error {
	var err error
	state.maps = make([]Map, 0, 0)
	for _, maps := range state.worldJSON.Maps {
		file := state.assetsDir + maps.FileName
		var newMap Map
		newMap.mapJSON, err = jsonLoader.LoadMap(file)
		newMap.size = Offset{maps.X, maps.Y}
		if err != nil {
			return err
		}
		state.maps = append(state.maps, newMap)
		if state.maps[0].loadLayers() != nil {
			log.Fatal("Failed to load map")
		}

	}
	//fmt.Println(state.maps)
	//fmt.Println(state.maps[0])
	return err
}

//loadLayers loads layers into a 2d slice of TileIDs from a 1d array in the JSON
func (Map *Map) loadLayers() error {
	Map.size.X = Map.mapJSON.Width
	Map.size.Y = Map.mapJSON.Height
	length := len(Map.mapJSON.Layers)
	Map.layers = make([]Layer, length, length)
	for i, val := range Map.mapJSON.Layers {

		//TODO: Change to switch statement
		if Map.mapJSON.Type == "tilelayer" {
			Map.layers[i] = &TileLayer{}
		}
		if Map.mapJSON.Type == "objectgroup" {
			Map.layers[i] = &ObjLayer{}
		}
		Map.layers[i].makeLayer(val)
		fmt.Println(Map.layers)

	}
	return nil
}

func (layer *ObjLayer) makeLayer(val jsonLoader.Layers) error {
	return nil
}

func (layer *TileLayer) makeLayer(data jsonLoader.Layers) error {
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
	for i, set := range state.tileSetJSON.Tileset {
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
