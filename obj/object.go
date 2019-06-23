package obj

import (
	"fmt"
	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
	"strings"
)

//Thing is a top level non tile object.
type Thing struct {
	Img    *ebiten.Image
	Pos    Vec
	MovSpd Vec //MovSpd sets a base movevemnt speed value for a Thing
}

//Entity is a Thing that has AI associated with it.
type Entity struct {
	Thing Thing
	//velocity Vec     //Use this if you need velocity
	//height int //height, used to determine collisions on multiheight maps
	//items
	//hold
}

//Player is an Entity that is controllable by a human
type Player struct {
	Entity Entity
	Button isPressed
}

//isPressed contains button state for the player. These are set with the Inputter method.
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

//Inputer uptdates the isPressed state for a player.
func (p *Player) Inputer() {
	//wasd controls
	if ebiten.IsKeyPressed(ebiten.KeyS) {
		p.Button.down = true
	} else {
		p.Button.down = false
	}
	if ebiten.IsKeyPressed(ebiten.KeyA) {
		p.Button.left = true
	} else {
		p.Button.left = false
	}
	if ebiten.IsKeyPressed(ebiten.KeyD) {
		p.Button.right = true
	} else {
		p.Button.right = false
	}
	if ebiten.IsKeyPressed(ebiten.KeyW) {
		p.Button.up = true
	} else {
		p.Button.up = false
	}
}

//ImgSetter methods
//ImgSetter sets an ebiten image in a Thing.
func (t *Thing) ImgSetter(folder, filename string) error {
	path := strings.Join([]string{folder, "/", filename}, "")
	fmt.Println(path)
	img, _, err := ebitenutil.NewImageFromFile(path, ebiten.FilterDefault)
	t.Img = img
	return err
}

//NewPlayer creates a new player struct and returns a pointer.
func NewPlayer() *Player {
	var p Player
	return &p
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

//Mover methods
//Mover takes a Player and updates its position.
func (p *Player) Mover() {
	//wasd controls
	p.Inputer()
	var offset Vec // Player to be moved by the amount in offset.
	if p.Button.up {
		offset.Y = -1.0
	}
	if p.Button.down {
		offset.Y = 1.0
	}
	if p.Button.left {
		offset.X = -1.0
	}
	if p.Button.right {
		offset.X = 1.0
	}
	p.Entity.Mover(offset)
}

//Mover takes an Entity and updates its position.
func (e *Entity) Mover(offset Vec) {
	e.Thing.Mover(offset)
}

//Mover takes a thing and updates its position.
func (t *Thing) Mover(offset Vec) {
	t.Pos.X = t.Pos.X + offset.X*t.MovSpd.X
	t.Pos.Y = t.Pos.Y + offset.Y*t.MovSpd.Y
}

//Extra types to make things work

type Vec struct {
	X float64
	Y float64
}
