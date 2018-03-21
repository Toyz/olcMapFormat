package main

import (
	"os"
	"strconv"
)

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

func times(str string, n int) (out string) {
	for i := 0; i < n; i++ {
		out += str
	}
	return
}

func padRight(str string, length int, pad string) string {
	return str + times(pad, length-len(str))
}

func fileExist(name string) bool {
	if _, err := os.Stat(name); err != nil {
		if os.IsNotExist(err) {
			return false
		}
	}
	return true
}

func toInt(value string) int {
	i, _ := strconv.Atoi(value)

	return i
}
