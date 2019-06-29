package main

import (
	"fmt"
	"github.com/hajimehoshi/ebiten"
)

func updateInputState() {
	var newInputState inputState
	//up button
	if ebiten.IsKeyPressed(ebiten.KeyW) && currentInputState.down == false {
		newInputState.up = true
	} else {
		newInputState.up = false
	}
	//down button
	if ebiten.IsKeyPressed(ebiten.KeyS) && currentInputState.up == false {
		newInputState.down = true
	} else {
		newInputState.down = false
	}
	//left button
	if ebiten.IsKeyPressed(ebiten.KeyA) && currentInputState.right == false {
		newInputState.left = true
	} else {
		newInputState.left = false
	}
	//right button
	if ebiten.IsKeyPressed(ebiten.KeyD) && currentInputState.left == false {
		newInputState.right = true
	} else {
		newInputState.right = false
	}
	fmt.Println(ebiten.IsKeyPressed(ebiten.KeyW))
	fmt.Println(ebiten.IsKeyPressed(ebiten.KeyS))
	fmt.Println(ebiten.IsKeyPressed(ebiten.KeyA))
	fmt.Println(ebiten.IsKeyPressed(ebiten.KeyD))
	currentInputState = newInputState
}