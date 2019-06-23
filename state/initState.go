package state

import (
	//	"encoding/json"
	//"fmt"
	//"github.com/hajimehoshi/ebiten"
	//"github.com/hajimehoshi/ebiten/ebitenutil"
	//obj "github.com/jfemory/goActionAdventure/object"
	//"github.com/jfemory/goActionAdventure/maps"
	"log"
)
import "github.com/jfemory/goActionAdventure/jsonLoader"

//TODO: Move state.assetsDir to a env var or global.

/*
//loadMaps takes a gamestate and loads the maps associated with its world.
func (state *gameState) loadMaps() error {
	var err error
	state.maps = make([]maps.Map, 0, 0)
	for _, maps := range state.worldJSON.Maps {
		file := state.assetsDir + maps.FileName
		var newMap maps.Map

		newMap.mapJSON, err = jsonLoader.LoadMap(file)
		newMap.size = Offset{maps.X, maps.Y}
		if err != nil {
			return err
		}
		state.maps = append(state.maps, newMap)

	}
	if state.maps[0].loadLayers() != nil {
		return err
	}
	//fmt.Println(state.maps)
	//fmt.Println(state.maps[1])
	return err
}

func makeEntityLayer(val jsonLoader.Layers, Map *Map) error {
	Map.entLayers = append(Map.entLayers, EntLayer{val.Name, nil})
	return nil
}

func makeObjLayer(val jsonLoader.Layers, Map *Map) error {
	//spin up layer
	fmt.Println("obj")
	Map.objLayers = append(Map.objLayers, ObjLayer{})
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

	for i, set := range state.tileSetJSON.{
		for j, val := range set.Tiles {

		}

	}
	return err
}

/*
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
