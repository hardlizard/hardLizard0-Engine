package game

import (
	//"fmt"
	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
	"strings"
)

//thing is a top level non tile object.
type thing struct {
	img *ebiten.Image
	x   float32
	y   float32
	vx  float32
	vy  float32
}

//entity is a thing that has AI associated with it.
type entity struct {
	thing thing
	//velocity vecFloat     //Use this if you need velocity
	//height int //height, used to determine collisions on multiheight maps
	//items
	//hold
}

//player is an entity that is controllable by a human
type player struct {
	entity entity
	Button isPressed
}

//sprite is a struct containing all the information needed to draw the thing to the screen.
type spriteObj struct {
	//[]*ebiten.Image
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
func (p *player) Inputer() {
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

//Constructors for player, entity, and thing.

//Newplayer creates a new player struct and returns a pointer.
func Newplayer() *player {
	var p player
	return &p
}

//Newentity creates a new entity struct and returns a pointer.
func Newentity() *entity {
	var e entity
	return &e
}

//Newthing creates a new thing and returns a pointer.
func Newthing() *thing {
	var t thing
	return &t
}

//Methods for setting ebiten images in objects.

//ImgSetter sets an ebiten image in a thing.
func (t *thing) ImgSetter(folder, filename string) error {
	path := strings.Join([]string{folder, "/", filename}, "")
	//fmt.Println(path)
	img, _, err := ebitenutil.NewImageFromFile(path, ebiten.FilterDefault)
	t.img = img
	return err
}

//ImgSetter sets an ebiten Image in an entity.
func (e *entity) ImgSetter(folder, filename string) error {
	err := e.thing.ImgSetter(folder, filename)
	return err
}

//ImgSetter sets an ebiten Image in a player.
func (p *player) ImgSetter(folder, filename string) error {
	err := p.entity.ImgSetter(folder, filename)
	return err
}

//Methods for moving objects on the mapBuffer.

//Mover takes a player and updates its position.
func (p *player) Mover() {
	//wasd controls
	p.Inputer()
	var offset vecFloat // player to be moved by the amount in offset.
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
	p.entity.Mover(offset)
}

//Mover takes an entity and updates its position.
func (e *entity) Mover(offset vecFloat) {
	e.thing.Mover(offset)
}

//Mover takes a thing and updates its position.
func (t *thing) Mover(offset vecFloat) {
	t.x = t.x + offset.X*t.vx
	t.y = t.y + offset.Y*t.vy
}

//Extra types to make things work

type vecFloat struct {
	X float32
	Y float32
}
