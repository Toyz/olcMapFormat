package main

type header struct {
	Size    int `struc:"int8,little,sizeof=Type"`
	Type    string
	Version int `struc:"int8"`
	Data    mapData
}

type mapData struct {
	Size       int `struc:"int32,little,sizeof=Name"`
	Name       string
	Width      int `struc:"int32"`
	Height     int `struc:"int32"`
	TilesCount int `struc:"int32,little,sizeof=Tiles"`
	Tiles      []tileData
}

type tileData struct {
	Position   int `struc:"int32"`
	TileID     int `struc:"int32"`
	ParamsSize int `struc:"int32,little,sizeof=Params"`
	Params     []tileParams
}

type tileParams struct {
	ParamNameSize int `struc:"int32,little,sizeof=ParamName"`
	ParamName     string
	MetaSize      int    `struc:"int32,little,sizeof=Meta"`
	Meta          string // String is required for this... So in the engine you must handle this as is
}
