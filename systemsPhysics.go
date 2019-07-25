package main

import ()

//Physics Update systems.
//updatePVA updates position and velocity arrays from the acceration array.
func updatePVA() {
	updateVelocity()
	updatePlayerVelocity()

	collision()
	updatePosition()
}

//updatePosition updates the posisition array from the velocity array
func updatePosition() {
	for i := 0; i < maxUsedEntity+1; i++ {
		game.position[i].x += game.velocity[i].x
		game.position[i].y += game.velocity[i].y
	}
}

//updateVelocity updates the velocity array from the accelearation array
func updateVelocity() {
	for i := 0; i < maxUsedEntity+1; i++ {

		game.velocity[i].x += game.acceleration[i].x
		game.velocity[i].y += game.acceleration[i].y

	}
}

//updatePlayerVelocity updates the player position based on current input
func updatePlayerVelocity() {
	if currentInputState.left == true {
		game.velocity[0].x = -1
	} else if currentInputState.right == true {
		game.velocity[0].x = 1
	} else {
		game.velocity[0].x = 0
	}
	if currentInputState.up == true {
		game.velocity[0].y = -1
	} else if currentInputState.down == true {
		game.velocity[0].y = 1
	} else {
		game.velocity[0].y = 0
	}
}

//sets are used to init PVA values.
func setPosition(x, y float32, gid uint) {
	game.position[gid].x = x
	game.position[gid].y = y
}

func setVelocity(vx, vy float32, gid uint) {
	game.velocity[gid].x = vx
	game.velocity[gid].y = vy
}

func setAcceleration(ax, ay float32, gid uint) {
	game.acceleration[gid].x = ax
	game.acceleration[gid].y = ay
}

//setHitbox sets the size of a circular hitbox,
func setHitbox(offset vecFloat, size float32, gid uint) {
	game.hitbox[gid].pos = offset //Set offset
	game.hitbox[gid].size = size  //set size
}
