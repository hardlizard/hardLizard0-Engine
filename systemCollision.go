package main

import "fmt"

//Collision ranges over all collidable objects and tests for collisions.
func collision() {
	for i := 0; i < maxUsedEntity; i++ {
		for j := i + 1; j < maxUsedEntity+1; j++ {
			if collidableF[i] == true && collidableF[j] == true {
				test, _ := collTest(hitbox[i], hitbox[j], position[i], position[j])
				if test == true {
					fmt.Println(health[0])
					heal(1, 0)
				}
			}
		}
	}
}

//collTest is a collision test for the main collision function.
//Remember, hb0 is the position
func collTest(hb0, hb1 circle, pos0, pos1 vecFloat) (bool, float32) {
	dx := (hb0.pos.x + pos0.x) - (hb1.pos.x + pos1.x)
	dy := (hb0.pos.x + pos0.y) - (hb1.pos.y + pos1.y)
	radii := hb0.size + hb1.size

	if (dx*dx)+(dy*dy) < (radii * radii) {
		offset := (radii * radii) - ((dx * dx) + (dy * dy))
		return true, offset
	}
	return false, 0
}

//func push(pusherGID, pusheeGID uint)

//TODO: For tile collisions w/ sprites, remember that the offset is zero when the hitbox is centered on the top left sprite corner. Remember to account for this when writing tile collision code. spritesize/2 offset.
