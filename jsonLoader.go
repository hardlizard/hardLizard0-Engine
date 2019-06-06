package main

import (
	"encoding/json"
	"os"
)

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
