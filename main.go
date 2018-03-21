package main

import (
	"bytes"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/lunixbochs/struc"
)

var (
	inputMap   = flag.String("in", "", "Input map file")
	outPutFile = flag.String("out", "", "Output map file")
	validate   = flag.Bool("validate", true, "To validate the output")
)

func main() {
	log.SetPrefix("[olcMapFormat] ")
	flag.Parse()

	if len(*inputMap) <= 0 || !fileExist(*inputMap) {
		log.Fatalf("%s does not exist!\n", *inputMap)
		os.Exit(1)
	}

	if len(*outPutFile) <= 0 {
		out := fmt.Sprintf("%s.map", getInputMapName(*inputMap))
		outPutFile = &out
	}

	handleCreate(*inputMap, *outPutFile)
}

func handleCreate(input, output string) {
	lines, _, _ := readMapFile(input)

	mapName := getInputMapName(input)
	log.Printf("Reading Map: %s\n", mapName)
	log.Printf("Width x Height: %s x %s\n", lines[0][0], lines[0][1])
	log.Printf("Map Payload Size: %v\n", len(lines[1]))
	log.Printf("Map Object Count: %v\n", len(lines[1])/2)
	log.Printf("Map Collision Count: %v\n", len(lines[1])/2)

	mapData := mapData{
		padRight(mapName, 32, " "),
		toInt(lines[0][0]),
		toInt(lines[0][1]),
		0,
		make([]tileData, 0),
	}

	id := 0
	for i := 1; i < len(lines[1]); i += 2 {
		tile := tileData{
			id,
			toInt(lines[1][i-1]),
			0,
			make([]tileParams, 0),
		}

		word := "solid"
		tile.Params = append(tile.Params, tileParams{
			padRight(word, 32, " "), 0, lines[1][i],
		})

		mapData.Tiles = append(mapData.Tiles, tile)
		id++
	}

	sprHeader := &header{
		"SPRMAP", 2, mapData,
	}

	var buf bytes.Buffer
	err := struc.Pack(&buf, sprHeader)
	if err != nil {
		log.Fatalln(err)
	}

	fOpen, _ := os.Create(output)
	fOpen.Write(buf.Bytes())
	fOpen.Close()

	if *validate {
		log.Println("Validating File...")
		head := &header{}
		fRead, _ := os.Open(output)
		struc.Unpack(fRead, head)
		mData := head.Data

		log.Printf("TYPE: %s Version: %d", head.Type, head.Version)
		log.Printf("Map Name: %s\n", mData.Name)
		log.Printf("Map Size: %dx%d\n", mData.Width, mData.Height)
		log.Printf("Tile Size: %d\n", len(mData.Tiles))

		log.Println("Dumping Tile Data...")
		for _, tile := range mData.Tiles {
			log.Println(tile)
		}
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

func getInputMapName(fp string) string {
	basename := filepath.Base(*inputMap)

	return strings.TrimSuffix(basename, filepath.Ext(basename))
}

func readMapFile(fp string) ([][]string, []string, string) {
	by, _ := ioutil.ReadFile(fp)

	contents := string(by)

	lines := strings.Split(contents, "\n")
	values := make([][]string, len(lines))

	for l := 0; l < len(lines); l++ {
		values[l] = strings.Split(lines[l], " ")
	}

	return values, lines, contents
}
