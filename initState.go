package main

import (
	//	"encoding/json"
	"fmt"
	//"github.com/hajimehoshi/ebiten"
	//"github.com/hajimehoshi/ebiten/ebitenutil"
	//obj "github.com/jfemory/goActionAdventure/object"
	"log"
)
import "github.com/jfemory/goActionAdventure/jsonLoader"

func (state *gameState) initAssets() error {
	var err error
	//Old Player position code
	//TODO: refactor to spawn based on sprite layer

	//state.player.Entity.Thing.Img, _, err = ebitenutil.NewImageFromFile("assets/gopherTop.png", ebiten.FilterDefault)
	state.player.ImgSetter("assets", "gopherTop.png")
	if err != nil {
		log.Fatal(err)
	}
	state.player.Entity.Thing.MovSpd.X, state.player.Entity.Thing.MovSpd.Y = 1.0, 1.0
	//TODO: Brittle code, replace w/ spawnpoint function
	state.player.Entity.Thing.Pos.X = windowWidth / 2
	state.player.Entity.Thing.Pos.Y = windowHeight / 2
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

	}
	if state.maps[0].loadLayers() != nil {
		return err
	}
	//fmt.Println(state.maps)
	//fmt.Println(state.maps[1])
	return err
}

//loadLayers loads layers into a 2d slice of TileIDs from a 1d array in the JSON
func (Map *Map) loadLayers() error {
	Map.size.X = Map.mapJSON.Width
	Map.size.Y = Map.mapJSON.Height
	//init counters for layers
	for _, val := range Map.mapJSON.Layers {
		//TODO: Change to switch statement
		if val.Type == "tilelayer" {
			makeTileLayer(val, Map)
			fmt.Println("make tilemap called")
		}
		if val.Type == "objectgroup" {
			makeObjLayer(val, Map)
		}
	}
	return nil
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

func makeTileLayer(val jsonLoader.Layers, Map *Map) error {
	//spin up layer
	tileMatrix := make([][]TileID, Map.size.X)
	i := 0 //i is the index of the tile array fromt the raw JSON data.
	for x := 0; x < Map.size.X; x++ {
		tileMatrix[x] = make([]TileID, Map.size.Y)
		for y := 0; y < Map.size.Y; y++ {
			tileMatrix[x][y] = TileID(val.Data[i])
			i++
		}
	}
	Map.tileLayers = append(Map.tileLayers, TileLayer{val.Name, tileMatrix, Map.size, Offset{32, 32}}) //TODO: Change hardcoded values to tilesize
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
