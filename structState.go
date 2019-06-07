package main

import (
	"github.com/hajimehoshi/ebiten"
	"github.com/jfemory/goActionAdventure/loader"
)

type gameState struct {
	worldName string
	gameMode  int
	assetsDir string
	player    Sprite
	maps      []Map
	atlas     Atlas
	camera    Position
	// raw JSON information
	worldJSON loader.World

	tileSetJSON loader.TileSet
}

//Map holds map file data
type Map struct {
	offset  Offset
	layers  []Layer
	mapJSON loader.Map
}

//Atlas holds the tiles to be rendered based on layer data
type Atlas struct {
	VRAM  []*ebiten.Image
	tiles []Tile
}

//Offset holds an offset in pixels
type Offset struct {
	X int
	Y int
}

//Position holds the
type Position struct {
	X float32
	Y float32
}

/////////////Rework////////////
//Layer holds layer data in a matrix of tile IDs.
type Layer struct {
	data [][]TileID
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

//initGameState initializes state so the game can load.
func (state *gameState) initGameState() error {
	var err error
	state.gameMode = 0
	state.assetsDir = "assets/"
	return err
}

//TileID is the ID of the tile as referenced in the Layer data
type TileID int
