package sensorserver

import "encoding/binary"

// IntBytes convert a int into a byte Array
func IntBytes(i int) []byte {
	bytes := make([]byte, 4)
	binary.BigEndian.PutUint32(bytes, uint32(i))
	return bytes
}
