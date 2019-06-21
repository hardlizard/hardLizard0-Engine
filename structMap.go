package main

import (
	//"github.com/hajimehoshi/ebiten"
	"github.com/jfemory/goActionAdventure/jsonLoader"
	obj "github.com/jfemory/goActionAdventure/object"
)

//Map holds map file data
type Map struct {
	size       Offset //Global map offset
	tileLayers []TileLayer
	objLayers  []ObjLayer
	entLayers  []EntLayer
	mapJSON    jsonLoader.Map //Raw JSON, use only of first load
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
	size     Offset
	tileSize Offset
	//size Offset
}

//TileID is the ID of the tile as referenced in the Layer data
type TileID int
