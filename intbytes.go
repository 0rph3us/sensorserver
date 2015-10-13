package sensorserver

import "encoding/binary"

func IntBytes(i int) []byte {
	bytes := make([]byte, 4)
	binary.LittleEndian.PutUint32(bytes, uint32(i))
	return bytes
}
