package game

import (
	"fmt"
	"github.com/jfemory/goActionAdventure/jsonLoader"
	"strings"
)

//NewWorld returns the address of a World in memory.
func NewWorld(folder string) (*World, error) {
	//TODO: refactor next line using path library
	path := strings.Join([]string{folder, "/", "world.world"}, "")
	jsonWorld, err := jsonLoader.LoadWorld(path)
	if err != nil {
		return nil, err
	}
	var w World
	w.mapFiles = make([]mapFile, len(jsonWorld.Maps))
	w.folder = folder

	//populate mapFiles
	for i, val := range jsonWorld.Maps {
		var currentMap mapFile
		currentMap.filename = val.FileName
		currentMap.x = val.X
		currentMap.y = val.Y
		w.mapFiles[i] = currentMap
	}
	w.folder = folder

	w.oldMap.loadMap(w.folder, w.mapFiles[0].filename)
	fmt.Println(w.oldMap)
	return &w, nil
}
