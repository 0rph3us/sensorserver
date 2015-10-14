package sensorserver

import (
	"github.com/boltdb/bolt"
)

type putdata struct {
	Timestamp int
	Value     float32
}

type Config struct {
	Database string `toml:"database"`
	Title    string `toml:"title"`
}

type Sensorserver struct {
	boltdb *bolt.DB
	conf   Config
}
