package main

const maxEntity = 10         //the maximum number of entities we can handle.
var maxUsedEntity int        //Change value if more than 65536 entities.
var freeList [maxEntity]bool //freeList sores free gid values less than the maxUsedEntity

//Entity Components
var hitbox [maxEntity]circle         //hitbox stores dimensions of the hitbox in int8 pixels
var position [maxEntity]vecFloat     //position holds the x, y coords of an entity in a float32, this is global
var velocity [maxEntity]vecFloat     //velocity holds the velocity components, similarly to position
var acceleration [maxEntity]vecFloat //acceleration holds... you get the ideax
var zHeight [maxEntity]int           //Height normal to the screen plane. This determines what layer something is on.
var health [maxEntity]int            //Current entity health.
var maxHealth [maxEntity]int         //Maximum health for an entity may have.
var currency [maxEntity]int          //In game currency.

//Entity Flags
var collidableF [maxEntity]bool //Flag for collidable.
var physicsF [maxEntity]bool    //Flag for physics engine.
var invincibleF [maxEntity]bool //Flag for invincibility.
var deadF [maxEntity]bool       //Flag when health goes below 0.
var projectileF [maxEntity]bool //Flag when entity is a projectile, run projectile logic on collision.

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
