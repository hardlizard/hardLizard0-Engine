package main

import (
	//"github.com/hajimehoshi/ebiten"
	"github.com/jfemory/goActionAdventure/jsonLoader"
)

//Map holds map file data
type Map struct {
	size    Offset         //Global map offset
	layers  []Layer        //Layer data, matrices of TileIDs
	mapJSON jsonLoader.Map //Raw JSON, use only of first load
}

//Layer interface holds layers of tiles or layers of sprites.
type Layer interface {
	makeLayer(jsonLoader.Layers) error
}

//ObjLayer holds sprite layer data
type ObjLayer struct {
	name string
	id   int
	data []Entity
	size Offset
}

//TileLayer holds layer data in a matrix of tile IDs.
type TileLayer struct {
	name string
	id   int
	data [][]TileID //The numerical ID of a tile's index in the atlas
	size Offset
}

//TileID is the ID of the tile as referenced in the Layer data
type TileID int
