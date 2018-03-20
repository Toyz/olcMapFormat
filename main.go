package main

import (
	"bytes"
	"flag"
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

type sprHeader struct {
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

func main() {
	log.SetPrefix("[olcMapFormat] ")

	flag.Parse()

	lines, _, _ := readMapFile(*inputMap)

	mapName := getInputMapName(*inputMap)
	log.Printf("Reading Map: %s\n", mapName)
	log.Printf("Width x Height: %s x %s\n", lines[0][0], lines[0][1])
	log.Printf("Map Payload Size: %v\n", len(lines[1]))
	log.Printf("Map Object Count: %v\n", len(lines[1]))
	log.Printf("Map Collision Count: %v\n", len(lines[1]))

	header := &sprHeader{
		0,
		mapName,
		toInt(lines[0][0]),
		toInt(lines[0][1]),
		0,
		make([]tileData, 0),
	}

	for i := 1; i < len(lines[1]); i += 2 {
		tile := tileData{
			i - 1,
			toInt(lines[1][i-1]),
			0,
			make([]tileParams, 0),
		}

		tile.Params = append(tile.Params, tileParams{
			0, "solid", 0, lines[1][i],
		})

		header.Tiles = append(header.Tiles, tile)

	}

	var buf bytes.Buffer
	err := struc.Pack(&buf, header)
	if err != nil {
		log.Panicln(err)
	}

	fOpen, _ := os.Create(*outPutFile)
	fOpen.Write(buf.Bytes())
	fOpen.Close()

	if *validate {
		log.Println("Validating File...")
		head := &sprHeader{}
		fRead, _ := os.Open(*outPutFile)
		struc.Unpack(fRead, head)

		log.Printf("Map Name: %s\n", head.Name)
		log.Printf("Map Name Size: %d\n", head.Size)
		log.Printf("Map Size: %dx%d\n", head.Width, head.Height)
		log.Printf("Tile Size: %d\n", len(header.Tiles))

		log.Println("Dumping Tile Data...")
		for _, tile := range head.Tiles {
			log.Println(tile)
		}
	}
}

func toInt(value string) int {
	i, _ := strconv.Atoi(value)

	return i
}

func getInputMapName(fp string) string {
	basename := filepath.Base(*inputMap)

	return strings.TrimSuffix(basename, filepath.Ext(basename))
}

func readMapFile(fp string) (map[int][]string, []string, string) {
	by, _ := ioutil.ReadFile(fp)

	contents := string(by)

	lines := strings.Split(contents, "\n")
	values := make(map[int][]string)

	for l := 0; l < len(lines); l++ {
		values[l] = strings.Split(lines[l], " ")
	}

	return values, lines, contents
}
