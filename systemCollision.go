package main

import "fmt"

//Collision ranges over all collidable objects and tests for collisions.
func collision() {
	for i := 0; i < maxUsedEntity; i++ {
		for j := i + 1; j < maxUsedEntity+1; j++ {
			test := collTest(hitbox[i], hitbox[j], position[i], position[j])
			if test == true {
				fmt.Println(health[0])
				heal(1, 0)
			}
		}
	}
}

//collTest is a collision test for the main collision function.
//Remember, hb0 is the position
func collTest(hb0, hb1 circle, pos0, pos1 vecFloat) bool {
	dx := (hb0.pos.x + pos0.x) - (hb1.pos.x + pos1.x)
	dy := (hb0.pos.x + pos0.y) - (hb1.pos.y + pos1.y)
	radii := hb0.size + hb1.size
	if (dx*dx)+(dy*dy) < (radii * radii) {
		return true
	}
	return false
}
