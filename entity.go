package main

func initPlayer() {
	gid := uint(0)
	//has position
	setPosition(100.0, 100.0, gid)
	setVelocity(0.0, 0.0, gid)
	setHitbox(vecFloat{0, 0}, 16, gid)
	initHealth(1, 100, gid)
	game.collidableF[gid] = true
	//	setAcceleration(0.0, 0.0, gid)
}

func newMob() {
	gid := uint(1)
	setPosition(50.0, 50.0, gid)
	setVelocity(0.0, 0.0, gid)
	setHitbox(vecFloat{0, 0}, 16, gid)
	initHealth(100, 100, gid)
	game.collidableF[gid] = true
	game.pushableF[gid] = true
	maxUsedEntity++
}
