package game

import (
	"fmt"
	"github.com/jfemory/goActionAdventure/jsonLoader"
	"strings"
)

type mapFile struct {
	filename string
	x        int
	y        int
}

type localMap struct {
	backgroundLayers []tileLayer
	tileLayers       []tileLayer
	objLayers        []objLayer
	forgroundLayers  []tileLayer
	width            int
	height           int
	tileHeight       int
	tileWidth        int
}

type tileLayer struct {
	id      int
	data    []int
	height  int
	width   int
	offsetX int
	offsetY int
	opacity int
}

type objLayer struct {
	things []thing
	ents   []entity
}

//loadMap loads a new map into a map pointer.
func (m *localMap) loadMap(folder, filename string) error {
	path := strings.Join([]string{folder, "/", filename}, "")
	fmt.Println(path)
	jsonMap, err := jsonLoader.LoadMap(path)
	m.height = jsonMap.Height
	m.width = jsonMap.Width
	m.tileHeight = jsonMap.Tilewidth
	m.tileWidth = jsonMap.Tilewidth
	for _, val := range jsonMap.Layers {
		if val.Type == "tilelayer" {
			m.tileLayers = append(m.tileLayers, tileLayer{val.ID, val.Data, val.Height, val.Width, val.X, val.Y, val.Opacity})
		} else if val.Type == "objectgroup" {

		}
	}
	return err
}
