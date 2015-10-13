package sensorserver

import (
	"encoding/binary"
	"math"
)

func Float32Bytes(float float32) []byte {
	bits := math.Float32bits(float)
	bytes := make([]byte, 4)
	binary.LittleEndian.PutUint32(bytes, bits)
	return bytes
}
