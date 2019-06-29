package main

func initPlayer() {
	gid := uint16(0)
	//has position
	setPosition(100.0, 100.0, gid)
	setVelocity(0.0, 0.0, gid)
	setHitbox(32, 32, gid)
	//	setAcceleration(0.0, 0.0, gid)
}

func newEnemy() {
	gid := uint16(1)
	setPosition(50.0, 50.0, gid)
	setVelocity(0.0, 0.0, gid)
	setHitbox(32, 32, gid)
	maxUsedEntity++
}
