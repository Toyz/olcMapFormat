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
		msg := fmt.Sprintf("%s does not exist!\n", *inputMap)
		if len(*inputMap) <= 0 {
			msg = "Input file cannot be empty"
		}
		log.Fatal(msg)
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

	mapData := mapBase{}
	mapData.MapNameSize = 0
	mapData.MapName = mapName
	mapData.MapDescSize = 0
	mapData.MapDesc = fmt.Sprintf("This map is called: %s", mapName)
	mapData.Width = toInt(lines[0][0])
	mapData.Height = toInt(strings.TrimSpace(lines[0][1]))
	mapData.LayersCount = 0
	mapData.Layers = make([]layer, 0)
	mapData.Layers = append(mapData.Layers, createLayer("background"))
	mapData.Layers = append(mapData.Layers, createLayer("collision"))

	pos := 0
	for i := 1; i < len(lines[1]); i += 2 {
		tid := toInt(lines[1][i-1])
		mapData.Layers[0].Tiles = append(mapData.Layers[0].Tiles, createTile(pos, tid))

		if lines[1][i] == "1" {
			mapData.Layers[1].Tiles = append(mapData.Layers[1].Tiles, createTile(pos, -1))
		}
		pos++
	}
	sprHeader := &header{
		0, "SPRMAP", 2, mapData,
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
		log.Printf("Map Name: %s\n", mData.MapName)
		log.Printf("Map Desc: %s\n", mData.MapDesc)
		log.Printf("Map Size: %dx%d\n", mData.Width, mData.Height)
		log.Printf("Total Layer Size: %d\n", len(mData.Layers))

		log.Println("Dumping Layer Info")

		for _, layer := range mData.Layers {
			log.Printf("Data for layer %s total tiles %d", layer.Name, layer.TilesCount)
			for _, tile := range layer.Tiles {
				log.Printf("IDX: %d TileID: %d\n", tile.Position, tile.TileID)
				if tile.ParamsSize > 0 {
					for _, param := range tile.Params {
						log.Printf("%s -> %s\n", param.Key, param.Value)
					}
				}
			}
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
