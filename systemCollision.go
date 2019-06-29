package main

func collision() {
	for i := 1; i < maxUsedEntity+1; i++ {
		if position[0]+float64(hitbox[0]) > position[i] && position[1] > position[i]-float64(hitbox[i]) {
			position[1] = position[i] + position[0] + float64(hitbox[0])
		}
	}
}
