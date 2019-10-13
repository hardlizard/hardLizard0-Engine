package main

import "fmt"

//Collision ranges over all collidable objects and tests for collisions.
func collision() {
	for i := 0; i < maxUsedEntity; i++ {
		for j := i + 1; j < maxUsedEntity+1; j++ {
			if game.collidableF[i] == true && game.collidableF[j] == true {
				test, _ := collTest(game.hitbox[i], game.hitbox[j], game.position[i], game.position[j])
				fmt.Println(game.health[0])
				if test == true {
					//fmt.Println(game.health[0])
					heal(1, 0)
				}
				if tileCollision(game.hitbox[i]) == true {
					fmt.Println("collision")
					game.position[i].y = float32(int(game.position[i].y)) //hacky, replace with sane function to set y position, add setting x position, too.
				}
			}
		}
	}
}

//tileCollision checks collision of sprites with tiles, and moves the sprites if needed.
func tileCollision(hb circle) bool {
	collidables := []int8{1}
	for _, value := range collidables {
		if levelMap[int(hb.pos.x)+int(hb.pos.y)*16] == value {
			return true
		}
	}
	return false
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
