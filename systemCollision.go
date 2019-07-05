package main

import "fmt"

func collision() {
	for i := 1; i < maxUsedEntity+1; i++ {
		dx := (hitbox[0].pos.x + position[0].x) - (hitbox[i].pos.x + position[i].x)
		dy := (hitbox[0].pos.y + position[0].y) - (hitbox[i].pos.y + position[i].y)
		radii := hitbox[0].size + hitbox[i].size
		if (dx*dx)+(dy*dy) < (radii * radii) {
			fmt.Println(health[0])
			health[0]--
		}
	}
}
