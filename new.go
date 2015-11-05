package sensorserver

import (
	"github.com/BurntSushi/toml"
	"github.com/boltdb/bolt"
	"log"
)

func New(filename, datadir string) (s *Sensorserver, port uint16, err error) {
	s = &Sensorserver{}

	// Read config
	_, err = toml.DecodeFile(filename, &s.conf)
	if err != nil {
		log.Fatal(err)
		return
	}

	// return port
	port = s.conf.Port

	// open BoltDB
	databasefile := datadir + s.conf.Database
	s.boltdb, err = bolt.Open(databasefile, 0600, nil)
	if err != nil {
		log.Fatal(err)
	}

	s.conf.Type = "single"

	return
}
