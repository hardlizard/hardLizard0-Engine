package main

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

//heal restores a set amount of health to an entity
func heal(healAmount int, gid uint) {
	max := maxHealth[gid]
	current := health[gid]
	if max < healAmount+current {
		health[gid] = max
	} else {
		health[gid] += healAmount
	}
}
