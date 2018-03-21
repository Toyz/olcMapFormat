package main

type header struct {
	Type    string `struc:"[6]int16"`
	Version int    `struc:"int8"`
	Data    mapData
}

type mapData struct {
	Name       string `struc:"[32]int16"` // Mac length is 32 bytes (Padded to fill)
	Width      int    `struc:"int16"`
	Height     int    `struc:"int16"`
	TilesCount int    `struc:"int32,little,sizeof=Tiles"`
	Tiles      []tileData
}

type tileData struct {
	Position   int `struc:"int16"`
	TileID     int `struc:"int16"`
	ParamsSize int `struc:"int32,little,sizeof=Params"`
	Params     []tileParams
}

type tileParams struct {
	ParamName string `struc:"[32]int16"` // Max length is 32 bytes (Padded to fill all)
	MetaSize  int    `struc:"int32,little,sizeof=Meta"`
	Meta      string // String is required for this... So in the engine you must handle this as is
}
