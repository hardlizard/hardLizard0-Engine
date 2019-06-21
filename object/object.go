package object

import (
	"fmt"
	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
	"strings"
)

type mover interface {
	Mover()
}

type imgSetter interface {
	ImgSetter()
}

//Thing is a top level non tile object.
type Thing struct {
	Img *ebiten.Image
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

//ImgSetter sets an ebiten image in a Thing.
func (t *Thing) ImgSetter(folder, filename string) error {
	path := strings.Join([]string{folder, "/", filename}, "")
	fmt.Println(path)
	img, _, err := ebitenutil.NewImageFromFile(path, ebiten.FilterDefault)
	t.Img = img
	return err
}

//ImgSetter sets an ebiten Image in an Entity.
func (e *Entity) ImgSetter(folder, filename string) error {
	err := e.Thing.ImgSetter(folder, filename)
	return err
}

//ImgSetter sets an ebiten Image in a Player.
func (p *Player) ImgSetter(folder, filename string) error {
	err := p.Entity.ImgSetter(folder, filename)
	return err
}

//Mover takes an entity and updates its position.
func (e *Entity) Mover() {
	//wasd controls
	if ebiten.IsKeyPressed(ebiten.KeyS) {
		e.Thing.Pos.Y = e.Thing.Pos.Y + 1
	}
	if ebiten.IsKeyPressed(ebiten.KeyA) {
		e.Thing.Pos.X = e.Thing.Pos.X - 1
	}
	if ebiten.IsKeyPressed(ebiten.KeyD) {
		e.Thing.Pos.X = e.Thing.Pos.X + 1
	}
	if ebiten.IsKeyPressed(ebiten.KeyW) {
		e.Thing.Pos.Y = e.Thing.Pos.Y - 1
	}
}

//Extra types to make everything just work

type vec struct {
	X float64
	Y float64
}
