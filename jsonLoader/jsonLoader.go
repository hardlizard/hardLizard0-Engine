package jsonLoader

import (
	"encoding/json"
	"os"
)

//loadWorldMap imports a world map from a json file. In this case, made with Tiled.
func LoadWorld(file string) (World, error) {
	var world World
	openFile, err := os.Open(file)
	defer openFile.Close()
	if err != nil {
		return world, err
	}
	jsonParser := json.NewDecoder(openFile)
	err = jsonParser.Decode(&world)
	return world, err
}

func LoadTileSet(file string) (TileSet, error) {
	var tileSet TileSet
	openFile, err := os.Open(file)
	defer openFile.Close()
	if err != nil {
		return tileSet, err
	}
	jsonParser := json.NewDecoder(openFile)
	err = jsonParser.Decode(&tileSet)
	return tileSet, err
}

func LoadMap(file string) (Map, error) {
	var loadedMap Map
	openFile, err := os.Open(file)
	defer openFile.Close()
	if err != nil {
		return loadedMap, err
	}
	jsonParser := json.NewDecoder(openFile)
	err = jsonParser.Decode(&loadedMap)
	return loadedMap, err
}
