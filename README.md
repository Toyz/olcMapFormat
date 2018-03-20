# olcMapFormat
A custom map format for the olcConsoleGameEngine by Javidx9

# About SPRMap Format
The SPRMap format or just _.map_ is a format which is a binary form of the currently existing _.txt_ map format in the olcConsoleGameEngine

# How it works
The format has a very basic struct layout that can be seen in the files _structs.go_ 
```golang
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
	Meta          string 
}
```

_Header_ struct contains the following data
```golang
type header struct {
	Size    int `struc:"int8,little,sizeof=Type"` // Is Size of Type which is the length
	Type    string // Always SPRMAP (making size 6)
	Version int `struc:"int8"` // Current version 1
	Data    mapData // the data of the physical map
}
```

_MapData_ struct contains the information about the give map

_TileData_ contains the given information on what a tile is in the engine such as TileID (the ID in the sprite sheet) and the Position which is the IDX of a 1D array, and the params

_TileParams_ are the intresting field in this as they are the _Rules_ you set for the tile such as Trigger or Solid... These rules can be dyamic as are basically a simple form of Key -> Value design

# Implementations
Coming Soon...