package sensorserver

import (
	"bytes"
	"encoding/binary"
	"log"
)

// BytesToFloat32 convert a byte Array to float32
func BytesToFloat32(b []byte) (f float32) {
	buf := bytes.NewReader(b)
	err := binary.Read(buf, binary.LittleEndian, &f)
	if err != nil {
		log.Println("binary.Read failed:", err)
	}

	return
}
