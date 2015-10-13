package sensorserver

import (
    "github.com/boltdb/bolt"
    "log"
)

func New(filename string) (s *Sensorserver, err error) {
    s = &Sensorserver{}
    s.boltdb, err = bolt.Open(filename, 0600, nil)
    if err != nil {
        log.Fatal(err)
    }
//    defer s.boltdb.Close()
    return
}
