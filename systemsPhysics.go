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
		position[2*i] += velocity[2*i]
		position[2*i+1] += velocity[2*i+1]
	}
}

//updateVelocity updates the velocity array from the accelearation array
func updateVelocity() {
	for i := 0; i < maxUsedEntity+1; i++ {

		velocity[2*i] += acceleration[2*i]
		velocity[2*i+1] += acceleration[2*i+1]

	}
}

//updatePlayerVelocity updates the player position based on current input
func updatePlayerVelocity() {
	if currentInputState.left == true {
		velocity[0] = -1
	} else if currentInputState.right == true {
		velocity[0] = 1
	} else {
		velocity[0] = 0
	}
	if currentInputState.up == true {
		velocity[1] = -1
	} else if currentInputState.down == true {
		velocity[1] = 1
	} else {
		velocity[1] = 0
	}
}

//sets are used to init PVA values.
func setPosition(x, y float64, gid uint16) {
	position[gid*2] = x
	position[gid*2+1] = y
}

func setVelocity(vx, vy float64, gid uint16) {
	velocity[gid*2] = vx
	velocity[gid*2+1] = vy
}

func setAcceleration(ax, ay float64, gid uint16) {
	acceleration[gid*2] = ax
	acceleration[gid*2+1] = ay
}

func setHitbox(x, y int8, gid uint16) {
	hitbox[gid*2] = x
	hitbox[gid*2+1] = y
}
