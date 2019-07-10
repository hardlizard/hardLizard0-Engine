package main

//Components
const maxEntity = 10                 //the maximum number of entities we can handle
var maxUsedEntity int                //Change value if more than 65536 entities.
var freeList [maxEntity]bool         //freeList sores free gid values less than the maxUsedEntity
var hitbox [maxEntity]circle         //hitbox stores dimensions of the hitbox in int8 pixels
var position [maxEntity]vecFloat     //position holds the x, y coords of an entity in a float32, this is global
var velocity [maxEntity]vecFloat     //velocity holds the velocity components, similarly to position
var acceleration [maxEntity]vecFloat //acceleration holds... you get the idea

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
