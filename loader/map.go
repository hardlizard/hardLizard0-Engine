package loader

////////////////////////////
//file.world related structs
////////////////////////////
//World holds a slice of map file names and the world name.
type World struct {
	Maps []Maps `json:"maps"`
	Type string `json:"type"`
}

//Maps holds the filename for each map and it's position in the world.
type Maps struct {
	FileName string `json:"fileName"`
	X        int    `json:"x"`
	Y        int    `json:"y"`
}

////////////////////////////
//Map file.json related structs
////////////////////////////
//Map holds map file related information.
type Map struct {
	Height       int        `json:"height"`
	Infinite     bool       `json:"infinite"`
	Layers       []Layers   `json:"layers"`
	Nextlayerid  int        `json:"nextlayerid"`
	Nextobjectid int        `json:"nextobjectid"`
	Orientation  string     `json:"orientation"`
	Renderorder  string     `json:"renderorder"`
	Tiledversion string     `json:"tiledversion"`
	Tileheight   int        `json:"tileheight"`
	Tilesets     []Tilesets `json:"tilesets"`
	Tilewidth    int        `json:"tilewidth"`
	Type         string     `json:"type"`
	Version      float64    `json:"version"`
	Width        int        `json:"width"`
}

//Layers holds layer data.
type Layers struct {
	Draworder string        `json:"draworder,omitempty"`
	ID        int           `json:"id"`
	Name      string        `json:"name"`
	Objects   []interface{} `json:"objects,omitempty"`
	Opacity   int           `json:"opacity"`
	Type      string        `json:"type"`
	Visible   bool          `json:"visible"`
	X         int           `json:"x"`
	Y         int           `json:"y"`
	Data      []int         `json:"data,omitempty"`
	Height    int           `json:"height,omitempty"`
	Width     int           `json:"width,omitempty"`
}

//Tilesets holds information on the tilesets used in the Map.
type Tilesets struct {
	Firstgid int    `json:"firstgid"`
	Source   string `json:"source"`
}
