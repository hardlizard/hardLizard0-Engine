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
		position[i].x += velocity[i].x
		position[i].y += velocity[i].y
	}
}

//updateVelocity updates the velocity array from the accelearation array
func updateVelocity() {
	for i := 0; i < maxUsedEntity+1; i++ {

		velocity[i].x += acceleration[i].x
		velocity[i].y += acceleration[i].y

	}
}

//updatePlayerVelocity updates the player position based on current input
func updatePlayerVelocity() {
	if currentInputState.left == true {
		velocity[0].x = -1
	} else if currentInputState.right == true {
		velocity[0].x = 1
	} else {
		velocity[0].x = 0
	}
	if currentInputState.up == true {
		velocity[0].y = -1
	} else if currentInputState.down == true {
		velocity[0].y = 1
	} else {
		velocity[0].y = 0
	}
}

//sets are used to init PVA values.
func setPosition(x, y float32, gid uint) {
	position[gid].x = x
	position[gid].y = y
}

func setVelocity(vx, vy float32, gid uint) {
	velocity[gid].x = vx
	velocity[gid].y = vy
}

func setAcceleration(ax, ay float32, gid uint) {
	acceleration[gid].x = ax
	acceleration[gid].y = ay
}

//setHitbox sets the size of a circular hitbox,
func setHitbox(offset vecFloat, size float32, gid uint) {
	hitbox[gid].pos = offset //Set offset
	hitbox[gid].size = size  //set size
}
