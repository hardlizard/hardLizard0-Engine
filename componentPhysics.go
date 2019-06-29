package main

//Components
const maxEntity = 10                    //the maximum number of entities we can handle
var maxUsedEntity int                   //Change value if more than 65536 entities.
var freeList [maxEntity]bool            //freeList sores free gid values less than the maxUsedEntity
var hitbox [maxEntity * 2]int8          //hitbox stores dimensions of the hitbox in int8 pixels
var position [maxEntity * 2]float64     //position holds the x, y coords of an entity in a float32, this is global
var velocity [maxEntity * 2]float64     //velocity holds the velocity components, similarly to position
var acceleration [maxEntity * 2]float64 //acceleration holds... you get the idea
