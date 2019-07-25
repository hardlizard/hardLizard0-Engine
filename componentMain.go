package main

const maxEntity = 10         //the maximum number of entities we can handle.
var maxUsedEntity int        //Change value if more than 65536 entities.
var freeList [maxEntity]bool //freeList sores free gid values less than the maxUsedEntity

type gameState struct {
	//Entity Components
	hitbox       [maxEntity]circle   //hitbox stores dimensions of the hitbox in int8 pixels
	position     [maxEntity]vecFloat //position holds the x, y coords of an entity in a float32, this is global
	velocity     [maxEntity]vecFloat //velocity holds the velocity components, similarly to position
	acceleration [maxEntity]vecFloat //acceleration holds... you get the ideax
	zHeight      [maxEntity]int      //Height normal to the screen plane. This determines what layer something is on.
	health       [maxEntity]int      //Current entity health.
	maxHealth    [maxEntity]int      //Maximum health for an entity may have.
	currency     [maxEntity]int      //In game currency.

	//Entity Flags
	collidableF [maxEntity]bool //Flag for collidable.
	physicsF    [maxEntity]bool //Flag for physics engine.
	invincibleF [maxEntity]bool //Flag for invincibility.
	deadF       [maxEntity]bool //Flag when health goes below 0.
	projectileF [maxEntity]bool //Flag when entity is a projectile, run projectile logic on collision.
	pushableF   [maxEntity]bool //Flag when entity is pushable by the player.
}

//Map Components

//vecFloat holds a floating point vector in two dimensions.
type vecFloat struct {
	x float32
	y float32
}

//vecInt holds an integer vector in two dimensions.
type vecInt struct {
	x int
	y int
}

type circle struct {
	pos  vecFloat
	size float32
}
