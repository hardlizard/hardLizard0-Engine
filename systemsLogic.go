package main

import "fmt"

//death checks to see if an entity's
//health is 0 or negative, killing that entity if so.
func death() {
	for i := 0; i < maxEntity; i++ {
		if game.health[i] <= 0 {
			game.deadF[i] = true
		}
	}
	if game.deadF[0] == true {
		playerDeath()
	}
}

func playerDeath() {
	gameMode = -1
	fmt.Println("player deadF")
}
