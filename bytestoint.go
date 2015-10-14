package sensorserver

import (
	"bytes"
	"encoding/binary"
	"log"
)

func BytesToInt(b []byte) int {
	var i uint32
	buf := bytes.NewReader(b)
	err := binary.Read(buf, binary.LittleEndian, &i)
	if err != nil {
		log.Println("binary.Read failed:", err)
	}

	return int(i)
}
