package main

import (
	"fmt"
	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
	"image/color"
	"log"
)
import "github.com/jfemory/goActionAdventure/loader"

/////////////////////Globals//////////////////////////////////
const windowTitle = "goAA"
const scale = 1
const windowWidth = 1024
const windowHeight = 768

var gameMode = 0

//////////////////////////////////////////////////////////////

//update is the main loop of the ebiten engine. Core loop is here.
func (state *gameState) update(screen *ebiten.Image) error {
	//gameMode 0 initializes assets
	if state.gameMode == 0 {
		state.initAssets()
		state.gameMode = 1
	}
	if ebiten.IsDrawingSkipped() {
		return nil
	}
	//Do stuff goes here
	screen.Fill(color.RGBA{255, 0, 0, 255})
	state.player.movement()
	state.drawSprite(screen)
	return nil
}

func (state *gameState) initAssets() error {
	var err error
	state.player.img, _, err = ebitenutil.NewImageFromFile("assets/gopherTop.png", ebiten.FilterDefault)
	if err != nil {
		log.Fatal(err)
	}
	state.player.posX = windowWidth / 2
	state.player.posY = windowHeight / 2
	state.json, err = loader.LoadWorldMap("assets/testMap.json")
	if err != nil {
		return err
	}
	err = state.loadLayers()
	return err
}

//loadLayers loads layers into a 2d slice from a 1d array in the JSON
func (state *gameState) loadLayers() error {
	numLayers := len(state.json.Layers)
	state.layers = make([]Layer, numLayers, numLayers)
	for i, val := range state.json.Layers {
		width := state.json.Layers[i].Width
		height := state.json.Layers[i].Height
		state.layers[i].data = make([][]TileID, width, width)
		for j := 0; j < width; j++ {
			state.layers[i].data[j] = make([]TileID, height, height)
		}
		for j := 0; j < width; j++ {
			for k := 0; k < height; k++ {
				index := (j * width) + k
				state.layers[i].data[j][k] = TileID(val.Data[index])
			}
		}
		fmt.Println(state.layers[i].data)
	}

	return nil
}

type gameState struct {
	gameMode  int
	assetsDir string
	layers    []Layer
	atlas     []Tile
	player    Sprite
	camera    Position
	json      loader.WorldMap
}

//Position holds the
type Position struct {
	X float32
	Y float32
}

type Layer struct {
	data [][]TileID
}

type Tile struct {
	img  *ebiten.Image
	size int
	//col  collision //collision geometry goes here
}

//initGameState initializes state so the game can load.
func (state *gameState) initGameState() error {
	var err error
	state.gameMode = 0
	state.assetsDir = "assets/"
	return err
}

type TileID int

func main() {
	var state gameState
	err := state.initGameState()
	if err != nil {
		log.Fatal("failed to initialize game state. ")
	}
	if err := ebiten.Run(state.update, windowWidth, windowHeight, scale, windowTitle); err != nil {
		log.Fatal(err)
	}
}
