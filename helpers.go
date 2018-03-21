package main

func createLayer(name string) layer {
	return layer{
		0, name, 0, make([]tile, 0),
	}
}

func createTile(pos int, tid int) tile {
	return tile{
		pos, tid, 0, make([]param, 0),
	}
}
