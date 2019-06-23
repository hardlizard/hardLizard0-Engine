package main

import (
	"github.com/jfemory/goActionAdventure/obj"
	//"gonum.org/v1/gonum/stat"
	//	"encoding/json"
	//"fmt"
	"github.com/hajimehoshi/ebiten"
	//"github.com/hajimehoshi/ebiten/ebitenutil"
	"github.com/jfemory/goActionAdventure/jsonLoader"
	//"github.com/jfemory/goActionAdventure/renderer"
	"image/color"
	"log"
)

/////////////////////Globals//////////////////////////////////
const windowTitle = "goAA"
const scale = 1
const windowWidth = 1024
const windowHeight = 768

var gameMode = 0

const assetsDir = "assets"

//////////////////////////////////////////////////////////////

//update is the main loop of the ebiten engine. Core loop is here.
func (s *gameState) update(screen *ebiten.Image) error {
	//gameMode 0 initializes assets
	if s.gameMode == 0 {
		if err := s.initAssets(vecInt{windowWidth, windowHeight}); err != nil {
			log.Fatal(err)
		}
		s.gameMode = 1
	}
	if ebiten.IsDrawingSkipped() {
		return nil
	}
	//Do stuff goes here
	screen.Fill(color.RGBA{255, 0, 0, 255})

	globalCount(s)
	s.player.Mover()
	//renderHelper(s, screen)
	drawSprite(s, screen)
	return nil
}

func main() {
	var s gameState
	err := s.initGameState()
	if err != nil {
		log.Fatal("failed to initialize game state. ")
	}
	if err := ebiten.Run(s.update, windowWidth, windowHeight, scale, windowTitle); err != nil {
		log.Fatal(err)
	}
}

func globalCount(state *gameState) {
	if state.globalCounter < 0xffffffffffffffff {
		state.globalCounter++
		//fmt.Println(state.globalCounter)
	}
}

//initGameState initializes state so the game can load.
func (s *gameState) initGameState() error {
	var err error
	s.gameMode = 0
	s.assetsDir = assetsDir
	return err
}

//initAssets initializes everything needed form the assets directory.
func (s *gameState) initAssets(winDim vecInt) error {
	var err error
	s.player = obj.NewPlayer()
	//state.Player.Entity.Thing.Img, _, err = ebitenutil.NewImageFromFile("assets/gopherTop.png", ebiten.FilterDefault)
	s.player.ImgSetter(assetsDir, "gopherTop.png")
	if err != nil {
		log.Fatal(err)
	}
	s.player.Entity.Thing.MovSpd.X, s.player.Entity.Thing.MovSpd.Y = 1.0, 1.0
	//TODO: Brittle code, replace w/ spawnpoint function
	s.player.Entity.Thing.Pos.X = float64(winDim.X) / 2
	s.player.Entity.Thing.Pos.Y = float64(winDim.Y) / 2
	//////////////////
	//New world load code
	s.WorldJSON, err = jsonLoader.LoadWorld("assets/AA.world")
	if err != nil {
		return err
	}
	/*	if err := state.loadMaps(); err != nil {
			return err
		}
	*/
	//err = state.loadLayers()
	return err

}

//gameState holds the current state of the game. Different fields are poniters to different subsystems.
type gameState struct {
	worldName     string
	gameMode      int
	assetsDir     string
	player        *obj.Player
	globalCounter uint
	//JSON fields
	WorldJSON jsonLoader.World //world map raw data, use only on first load.
}

//Vec holds the absolute screen position in floats.
type Vec struct {
	X float64
	Y float64
}

//vecInt holds a position in ints.
type vecInt struct {
	X int
	Y int
}
