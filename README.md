# olcMapFormat
A custom map format for the olcConsoleGameEngine by Javidx9

# About SPRMap Format
The SPRMap format or just _.map_ is a format which is a binary form of the currently existing _.txt_ map format in the olcConsoleGameEngine

# How it works
The format has a very basic struct layout that can be seen in the files _structs.go_ 

### **Format Breakdown**


_Header_ struct contains the following data
```golang
type header struct {
    // Is Size of Type which is the length
    Size    int `struc:"int8,little,sizeof=Type"` 
    // Always SPRMAP (making size 6)
    Type    string
    // Current Version 2
    Version int `struc:"int8"`
    // Physical map information such as tiles
    Data    mapData 
}
```

_MapData_ struct contains the information about the give map
```golang
type mapBase struct {
    // Size of the map name
    MapNameSize int `struc:"int32,little,sizeof=MapName"`
    // Map Name
    MapName     string
    // MapDesc Size
    MapDescSize int `struc:"int32,little,sizeof=MapDesc"`
    // Map description
    MapDesc     string
    // Map Width
    Width       int `struc:"int16"`
    // Map Height
    Height      int `struc:"int16"`
    // Total layers
    LayersCount int `struc:"int32,sizeof=Layers"`
    // Layer info
    Layers      []layer
}
```

_layerData_ is the base object of how a layer is stored which is a `Key -> array(tile)`
```golang
type layer struct {
    // Size of the name of the layer
    NameSize   int `struc:"int32,sizeof=Name"`
    // Layer name
    Name       string
    // Total tiles in layer
    TilesCount int `struc:"int32,sizeof=Tiles"`
    // Tiles that are within the layer
    Tiles      []tile
}
```

_TileData_ contains the given information on what a tile is in the engine such as TileID (the ID in the sprite sheet) and the Position which is the IDX of a 1D array, and the params
```golang
type tile struct {
    // ID from the Sprite Sheet
    TileID     int `struc:"int32"`
    // Total Params for this tile
    ParamsSize int `struc:"int32,little,sizeof=Params"`
    // Array of all Params (EX isSolid, isTrigger)
    Params     []param
}
```

_TileParams_ are the intresting field in this as they are the _Rules_ you set for the tile such as Trigger or Solid... These rules can be dyamic as are basically a simple form of `Key -> Value` design
```golang
type param struct {
    //Size of the Key 
    KeySize   int `struc:"int32,little,sizeof=Key"`
    // Key name (ex: isSolid, isTrigger)
    Key       string
    // Size of the value
    ValueSize int `struc:"int32,little,sizeof=Value"`
    // Value can be anything but always turned into a string
    Value     string
}
```

# Implementations
Coming Soon...