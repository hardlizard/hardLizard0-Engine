package maps

import (
	//"github.com/hajimehoshi/ebiten"
	"github.com/jfemory/goActionAdventure/jsonLoader"
	"github.com/jfemory/goActionAdventure/obj"
)

//Map holds map file data
type Map struct {
	Size       vecInt //Global map offset
	TileLayers []TileLayer
	ObjLayers  []ObjLayer
	EntLayers  []EntLayer
	MapJSON    jsonLoader.Map //Raw JSON, use only of first load
}

//ObjLayer holds sprite layer data
type ObjLayer struct {
	name string
	objs []Object
}

type EntLayer struct {
	name     string
	entities []obj.Entity
}

type Object struct {
}

//TileLayer holds layer data in a matrix of tile IDs.
type TileLayer struct {
	name     string
	data     [][]TileID //The numerical ID of a tile's index in the atlas
	size     vecInt
	tileSize vecInt
	//size Offset
}

//TileID is the ID of the tile as referenced in the Layer data
type TileID int

type vecInt struct {
	X int
	Y int
}

//loadLayers loads layers into a 2d slice of TileIDs from a 1d array in the JSON
func (Map *Map) loadLayers() error {
	Map.Size.X = Map.MapJSON.Width
	Map.Size.Y = Map.MapJSON.Height
	//init counters for layers
	for _, val := range Map.MapJSON.Layers {
		//TODO: Change to switch statement
		if val.Type == "tilelayer" {
			makeTileLayer(val, Map)
		}
		//if val.Type == "objectgroup" {
		//		obj.makeObjLayer(val, Map)
		//		}
	}
	return nil
}

func makeTileLayer(val jsonLoader.Layers, Map *Map) error {
	//spin up layer
	tileMatrix := make([][]TileID, Map.Size.X)
	i := 0 //i is the index of the tile array fromt the raw JSON data.
	for x := 0; x < Map.Size.X; x++ {
		tileMatrix[x] = make([]TileID, Map.Size.Y)
		for y := 0; y < Map.Size.Y; y++ {
			tileMatrix[x][y] = TileID(val.Data[i])
			i++
		}
	}
	Map.TileLayers = append(Map.TileLayers, TileLayer{val.Name, tileMatrix, Map.Size, vecInt{32, 32}}) //TODO: Change hardcoded values to tilesize
	return nil
}
