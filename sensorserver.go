package sensorserver

import (
	"github.com/boltdb/bolt"
)

type Sensorserver struct {
	boltdb *bolt.DB
}
