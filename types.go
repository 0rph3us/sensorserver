package sensorserver

import (
	"github.com/boltdb/bolt"
)

type singledata struct {
	Timestamp int
	Value     float32
}

type multidata struct {
	Timestamp int
	Sensors   map[string]float32
}

type highchartData struct {
	T int64
	V float32
}

type plotBands struct {
	From  int64  `json:"from"`
	To    int64  `json:"to"`
	Color string `json:"color"`
}

type Config struct {
	Database  string  `toml:"database"`
	Title     string  `toml:"title"`
	Caption   string  `toml:"caption"`
	Latitude  float64 `toml:"latitude"`
	Longitude float64 `toml:"longitude"`
	SubTitle  string
	Duration  string
}

type Sensorserver struct {
	boltdb *bolt.DB
	conf   Config
}
