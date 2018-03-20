package buffer

import (
	"bytes"
	"encoding/binary"
	"log"
	"os"
)

type buf struct {
	w    *bytes.Buffer
	file *os.File
}

func New(file string) *buf {
	fOpen, _ := os.Create(file)

	return &buf{
		new(bytes.Buffer),
		fOpen,
	}
}

func (buf *buf) Save() {
	buf.file.Write(buf.w.Bytes())
}

func (buf *buf) PutString(data string) {
	bytes := []byte(data)
	size := len(bytes)
	buf.PutInt32(int32(size))
	buf.Write(bytes)
}

func (buf *buf) PutInt8(data int8) {
	buf.Write(data)
}

func (buf *buf) PutInt16(data int16) {
	buf.Write(data)
}

func (buf *buf) PutInt32(data int32) {
	buf.Write(data)
}

func (buf *buf) PutInt64(data int64) {
	buf.Write(data)
}

func (buf *buf) Write(data interface{}) {
	err := binary.Write(buf.w, binary.BigEndian, data)
	if err != nil {
		log.Fatalln(err)
	}
}
