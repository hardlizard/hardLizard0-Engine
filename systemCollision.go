package main

func collision() {
	for i := 1; i < maxUsedEntity+1; i++ {
		if position[0] < position[i]+float64(hitbox[i]) &&
			position[0]+float64(hitbox[0]) > position[i] &&
			position[1] < position[i+1]+float64(hitbox[i+1]) &&
			position[1]+float64(hitbox[1]) > position[i+1] {
			velocity[0] = 0
		}
	}
}
