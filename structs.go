package main

type header struct {
	TypeSize int `struc:"int8,sizeof=Type"`
	Type     string
	Version  int `struc:"int8"`
	Data     mapBase
}

type mapBase struct {
	MapNameSize int `struc:"int32,sizeof=MapName"`
	MapName     string
	MapDescSize int `struc:"int32,sizeof=MapDesc"`
	MapDesc     string
	Width       int `struc:"int32"`
	Height      int `struc:"int32"`
	LayersCount int `struc:"int32,sizeof=Layers"`
	Layers      []layer
}

type layer struct {
	NameSize   int `struc:"int32,sizeof=Name"`
	Name       string
	TilesCount int `struc:"int32,sizeof=Tiles"`
	Tiles      []tile
}

type tile struct {
	Position   int `struc:"int32"`
	TileID     int `struc:"int32"`
	ParamsSize int `struc:"int32,sizeof=Params"`
	Params     []params
}

type params struct {
	KeySize   int `struc:"int32,sizeof=Key"`
	Key       string
	ValueSize int `struc:"int32,sizeof=Value"`
	Value     string
}
