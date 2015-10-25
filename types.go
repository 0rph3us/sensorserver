package sensorserver

import (
	"github.com/boltdb/bolt"
)

type singleData struct {
	Timestamp int
	Value     float32
}

type minMaxData struct {
	Timestamp int
	MinValue  float32
	MaxValue  float32
	AvgValue  float32
}

type multidata struct {
	Timestamp int
	Sensors   map[string]float32
}

type plotBands struct {
	From  int64  `json:"from"`
	To    int64  `json:"to"`
	Color string `json:"color"`
}

type Config struct {
	Database  string   `toml:"database"`
	Title     string   `toml:"title"`
	Caption   string   `toml:"caption"`
	Latitude  float64  `toml:"latitude"`
	Longitude float64  `toml:"longitude"`
	MaxPoints int      `toml:"maxPoints"`
	Sensors   []string `toml:"sensors"`
	SubTitle  string
	Duration  string
}

type Sensorserver struct {
	boltdb *bolt.DB
	conf   Config
}
