package main

import (
	"encoding/json"
	"fmt"
	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
	"image/color"
	"log"
	"os"
)

/////////////////////Globals//////////////////////////////////
const windowTitle = "goAA"
const scale = 1
const windowWidth = 1024
const windowHeight = 768

var gameMode = 0
var mc sprite

//////////////////////////////////////////////////////////////

type sprite struct {
	posX float64 //x position
	posY float64 //y position
	//	velocity vec  //Use this if you need velocity
	height int //height, used to determine collisions on multiheight maps
	img    *ebiten.Image
	//items
	//hold
}

//update is the main loop of the ebiten engine. Core loop is here.
func update(screen *ebiten.Image) error {
	//gameMode 0 initializes assets
	if gameMode == 0 {
		initAssets()
		gameMode = 1
	}
	if ebiten.IsDrawingSkipped() {
		return nil
	}
	//Do stuff goes here
	screen.Fill(color.RGBA{255, 0, 0, 255})
	mc.movement()
	mc.drawSprite(screen)
	return nil
}

func (s *sprite) movement() {
	//wasd controls
	if ebiten.IsKeyPressed(ebiten.KeyD) {
		s.posX = s.posX + 1
	}
	if ebiten.IsKeyPressed(ebiten.KeyA) {
		s.posX = s.posX - 1
	}
	if ebiten.IsKeyPressed(ebiten.KeyS) {
		s.posY = s.posY + 1
	}
	if ebiten.IsKeyPressed(ebiten.KeyW) {
		s.posY = s.posY - 1
	}
}

func (s sprite) drawSprite(screen *ebiten.Image) {
	opts := &ebiten.DrawImageOptions{}
	opts.GeoM.Translate(mc.posX, mc.posY)
	screen.DrawImage(mc.img, opts)
}

func main() {
	if err := ebiten.Run(update, windowWidth, windowHeight, scale, windowTitle); err != nil {
		log.Fatal(err)
	}
}

func initAssets() {
	var err error
	mc.img, _, err = ebitenutil.NewImageFromFile("gopherTop.png", ebiten.FilterDefault)
	if err != nil {
		log.Fatal(err)
	}
	mc.posX = windowWidth / 2
	mc.posY = windowHeight / 2
	world, _ := loadWorldMap("testMap.json")
	fmt.Println(world)
}

//Tile Engine

func loadWorldMap(file string) (WorldMap, error) {
	var world WorldMap
	openFile, err := os.Open(file)
	defer openFile.Close()
	if err != nil {
		return world, err
	}
	jsonParser := json.NewDecoder(openFile)
	err = jsonParser.Decode(&world)
	return world, err
}

type WorldMap struct {
	Height       int        `json:"height"`
	Infinite     bool       `json:"infinite"`
	Layers       []Layers   `json:"layers"`
	Nextlayerid  int        `json:"nextlayerid"`
	Nextobjectid int        `json:"nextobjectid"`
	Orientation  string     `json:"orientation"`
	Renderorder  string     `json:"renderorder"`
	Tiledversion string     `json:"tiledversion"`
	Tileheight   int        `json:"tileheight"`
	Tilesets     []Tilesets `json:"tilesets"`
	Tilewidth    int        `json:"tilewidth"`
	Type         string     `json:"type"`
	Version      float64    `json:"version"`
	Width        int        `json:"width"`
}
type Layers struct {
	Data    []int  `json:"data"`
	Height  int    `json:"height"`
	ID      int    `json:"id"`
	Name    string `json:"name"`
	Opacity int    `json:"opacity"`
	Type    string `json:"type"`
	Visible bool   `json:"visible"`
	Width   int    `json:"width"`
	X       int    `json:"x"`
	Y       int    `json:"y"`
}
type Grid struct {
	Height      int    `json:"height"`
	Orientation string `json:"orientation"`
	Width       int    `json:"width"`
}
type Animation struct {
	Duration int `json:"duration"`
	Tileid   int `json:"tileid"`
}
type Objects struct {
	Height   float64 `json:"height"`
	ID       int     `json:"id"`
	Name     string  `json:"name"`
	Rotation int     `json:"rotation"`
	Type     string  `json:"type"`
	Visible  bool    `json:"visible"`
	Width    float64 `json:"width"`
	X        float64 `json:"x"`
	Y        float64 `json:"y"`
}
type Objectgroup struct {
	Draworder string    `json:"draworder"`
	Name      string    `json:"name"`
	Objects   []Objects `json:"objects"`
	Opacity   int       `json:"opacity"`
	Type      string    `json:"type"`
	Visible   bool      `json:"visible"`
	X         int       `json:"x"`
	Y         int       `json:"y"`
}
type Tiles struct {
	Animation   []Animation `json:"animation,omitempty"`
	ID          int         `json:"id"`
	Image       string      `json:"image"`
	Imageheight int         `json:"imageheight"`
	Imagewidth  int         `json:"imagewidth"`
	Objectgroup Objectgroup `json:"objectgroup,omitempty"`
}
type Tilesets struct {
	Columns    int     `json:"columns"`
	Firstgid   int     `json:"firstgid"`
	Grid       Grid    `json:"grid"`
	Margin     int     `json:"margin"`
	Name       string  `json:"name"`
	Spacing    int     `json:"spacing"`
	Tilecount  int     `json:"tilecount"`
	Tileheight int     `json:"tileheight"`
	Tiles      []Tiles `json:"tiles"`
	Tilewidth  int     `json:"tilewidth"`
}
