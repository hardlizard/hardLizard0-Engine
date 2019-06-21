package object

import (
	"github.com/hajimehoshi/ebiten"
)

type mover interface {
	Mover()
}

//Thing is a top level non tile object.
type Thing struct {
	Pos vec
}

//Entity is a Thing that has AI associated with it.
type Entity struct {
	Thing Thing
	//velocity vec     //Use this if you need velocity
	//height int //height, used to determine collisions on multiheight maps
	Img *ebiten.Image
	//items
	//hold
}

type Player struct {
	Entity Entity
	Button isPressed
}

type isPressed struct {
	up    bool
	down  bool
	left  bool
	right bool
	a     bool
	b     bool
	x     bool
	y     bool
	l     bool
	r     bool
	start bool
	sel   bool
}

//Mover takes an entity and updates its position.
func (s *Entity) Mover() {
	//wasd controls
	if ebiten.IsKeyPressed(ebiten.KeyS) {
		s.Thing.Pos.Y = s.Thing.Pos.Y + 1
	}
	if ebiten.IsKeyPressed(ebiten.KeyA) {
		s.Thing.Pos.X = s.Thing.Pos.X - 1
	}
	if ebiten.IsKeyPressed(ebiten.KeyD) {
		s.Thing.Pos.X = s.Thing.Pos.X + 1
	}
	if ebiten.IsKeyPressed(ebiten.KeyW) {
		s.Thing.Pos.Y = s.Thing.Pos.Y - 1
	}
}

//Extra types to make everything just work

type vec struct {
	X float64
	Y float64
}
