package game

import (
	//"github.com/hajimehoshi/ebiten"
	"errors"
)

//World holds the current game world data, including maps, player position, entities, etc.
type World struct {
	player   player
	mapFiles []mapFile
	oldMap   localMap
	newMap   localMap
	folder   string
	//Game Mode: 0 - init state
	//			 1 - run
	//			-1 - Error
	mode          int
	cameraX       float32
	cameraY       float32
	globalCounter uint
}

//SetMode sets the world's gameMode to an integer. Returns an error.
func (w *World) SetMode(mode int) error {
	w.mode = mode
	if mode != w.mode {
		return errors.New("Failed to assign game mode. Missmatch occured. ")
	}
	return nil
}

//GetMode returns the int value of the world's game mode.
func (w *World) GetMode() int {
	return w.mode
}
