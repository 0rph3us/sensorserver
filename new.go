package sensorserver

import (
	"github.com/BurntSushi/toml"
	"github.com/boltdb/bolt"
	"log"
)

func New(filename string) (s *Sensorserver, err error) {
	s = &Sensorserver{}

	// Read config
	_, err = toml.DecodeFile(filename, &s.conf)
	if err != nil {
		log.Fatal(err)
		return
	}

	s.boltdb, err = bolt.Open(s.conf.Database, 0600, nil)
	if err != nil {
		log.Fatal(err)
	}
	return
}
