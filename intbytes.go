package sensorserver

import "encoding/binary"

func IntBytes(i int) []byte {
	bytes := make([]byte, 4)
	binary.BigEndian.PutUint32(bytes, uint32(i))
	return bytes
}
