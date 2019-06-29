package main

import "fmt"

func initPlayer() {
	gid := uint16(0)
	//has position
	setPosition(100.0, 100.0, gid)
	setVelocity(0.0, 0.0, gid)
	//	setAcceleration(0.0, 0.0, gid)
	fmt.Println(velocity)
}

func newEnemy() {

}
