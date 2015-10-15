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
	Caption  string `toml:"caption"`
	SubTitle string
}

type Sensorserver struct {
	boltdb *bolt.DB
	conf   Config
}
