package main

import (
	"github.com/hajimehoshi/ebiten"
	"github.com/jfemory/goActionAdventure/jsonLoader"
	"github.com/jfemory/goActionAdventure/obj"
)

type gameState struct {
	worldName     string
	gameMode      int
	assetsDir     string
	player        obj.Player
	maps          []Map
	atlas         Atlas
	camera        Vec
	globalCounter uint
	// raw JSON information
	worldJSON   jsonLoader.World     //world map raw data, use only on first load.
	tileSetJSON []jsonLoader.TileSet //json tileset info, use on first load to build tiles in atlas.
}

//Atlas holds the tiles to be rendered based on layer data
type Atlas struct {
	VRAM  []*ebiten.Image //The ebiten images used to render tiles
	tiles []Tile          //Tile information, useful for dynamic tile loading
}

//Tile is the basic unit of the world map, they may be animated if bool is true.
type Tile struct {
	frames   []Frame
	animated bool
	size     int
	//col  collision //collision geometry goes here
}

//Frame holds a pointer to one frame of animatinon and its duraiton
type Frame struct {
	img      *ebiten.Image
	duraiton int //in ms
}

//Offset holds an offset in pixels.
type Offset struct {
	X int
	Y int
}

//Position holds the absolute screen position in a float.
type Vec struct {
	X float64
	Y float64
}
