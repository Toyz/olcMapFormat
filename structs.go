package main

type header struct {
	Size    int `struc:"int8,little,sizeof=Type"`
	Type    string
	Version int `struc:"int8"`
	Data    mapData
}

type mapData struct {
	MapNameSize int `struc:"int32,little,sizeof=MapName"`
	MapName     string
	MapDescSize int `struc:"int32,little,sizeof=MapDesc"`
	MapDesc     string
	Width       int `struc:"int16"`
	Height      int `struc:"int16"`
	TilesCount  int `struc:"int32,little,sizeof=Tiles"`
	Tiles       []tileData
}

type tileData struct {
	TileID     int `struc:"int32"`
	ParamsSize int `struc:"int32,little,sizeof=Params"`
	Params     []tileParams
}

type tileParams struct {
	Size      int `struc:"int32,little,sizeof=Key"`
	Key       string
	ValueSize int `struc:"int32,little,sizeof=Value"`
	Value     string
}
