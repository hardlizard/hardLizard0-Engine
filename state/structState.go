package state

import (
	"github.com/hajimehoshi/ebiten"
	"github.com/jfemory/goActionAdventure/jsonLoader"
	"github.com/jfemory/goActionAdventure/maps"
	"github.com/jfemory/goActionAdventure/obj"
)

type GameState struct {
	WorldName     string
	GameMode      int
	AssetsDir     string
	Player        obj.Player
	Maps          []maps.Map
	Atlas         Atlas
	Camera        Vec
	GlobalCounter uint
	// raw JSON information
	
	TileSetJSON []jsonLoader.TileSet //json tileset info, use on first load to build tiles in atlas.
	WindowDim   VecInt
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


