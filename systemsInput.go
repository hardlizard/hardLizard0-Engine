package main

import (
	"github.com/hajimehoshi/ebiten"

	"fmt"
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
	//circle button
	if ebiten.IsKeyPressed(ebiten.KeyRight) {
		newInputState.circle = true
		fmt.Println("cirlce")
	} else {
		newInputState.circle = false
	}
	//triangle button
	if ebiten.IsKeyPressed(ebiten.KeyUp) {
		newInputState.triangle = true
		fmt.Println("triangle")
	} else {
		newInputState.triangle = false
	}
	//square button
	if ebiten.IsKeyPressed(ebiten.KeyLeft) {
		newInputState.square = true
		fmt.Println("square")
	} else {
		newInputState.square = false
	}
	//cross button
	if ebiten.IsKeyPressed(ebiten.KeyDown) {
		newInputState.cross = true
		fmt.Println("cross")
	} else {
		newInputState.cross = false
	}
	//start button
	if ebiten.IsKeyPressed(ebiten.KeyLeft) {
		newInputState.square = true
	} else {
		newInputState.square = false
	}
	currentInputState = newInputState
}
