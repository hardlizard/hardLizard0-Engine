package main

var health [maxEntity]int
var maxHealth [maxEntity]int
var dead [maxEntity]bool

//setHealth sets the current health for an entity.
func setHealth(currentHealth int, gid uint) {
	health[gid] = currentHealth
}

//setMaxHealth sets the max health of an entity.
func setMaxHealth(mh int, gid uint) {
	maxHealth[gid] = mh
}

//initHealth initializes both health and max health.
func initHealth(ch, mh int, gid uint) {
	setHealth(ch, gid)
	setMaxHealth(mh, gid)
}
