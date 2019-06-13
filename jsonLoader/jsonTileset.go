package jsonLoader

//TileSet
type TileSet struct {
	Columns      int     `json:"columns"`
	Grid         Grid    `json:"grid"`
	Margin       int     `json:"margin"`
	Name         string  `json:"name"`
	Spacing      int     `json:"spacing"`
	Tilecount    int     `json:"tilecount"`
	Tiledversion string  `json:"tiledversion"`
	Tileheight   int     `json:"tileheight"`
	Tiles        []Tiles `json:"tiles"`
	Tilewidth    int     `json:"tilewidth"`
	Type         string  `json:"type"`
	Version      float64 `json:"version"`
}
type Grid struct {
	Height      int    `json:"height"`
	Orientation string `json:"orientation"`
	Width       int    `json:"width"`
}
type Objects struct {
	Height   int    `json:"height"`
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Rotation int    `json:"rotation"`
	Type     string `json:"type"`
	Visible  bool   `json:"visible"`
	Width    int    `json:"width"`
	X        int    `json:"x"`
	Y        int    `json:"y"`
}
type Objectgroup struct {
	Draworder string    `json:"draworder"`
	Name      string    `json:"name"`
	Objects   []Objects `json:"objects"`
	Opacity   int       `json:"opacity"`
	Type      string    `json:"type"`
	Visible   bool      `json:"visible"`
	X         int       `json:"x"`
	Y         int       `json:"y"`
}
type Animation struct {
	Duration int `json:"duration"`
	Tileid   int `json:"tileid"`
}
type Tiles struct {
	ID          int         `json:"id"`
	Image       string      `json:"image"`
	Imageheight int         `json:"imageheight"`
	Imagewidth  int         `json:"imagewidth"`
	Objectgroup Objectgroup `json:"objectgroup,omitempty"`
	Type        string      `json:"type"`
	Animation   []Animation `json:"animation,omitempty"`
}
