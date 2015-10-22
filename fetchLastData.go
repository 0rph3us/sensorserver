package sensorserver

import (
	"bytes"
	"errors"
	"github.com/boltdb/bolt"
)

// return all data from max(Timestamp) - duration for a specific sensor
func (s *Sensorserver) fetchLastData(sensor string, duration int) (data []singleData, err error) {

	// return all Values, without parameter
	err = s.boltdb.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(sensor))
		if b == nil {
			msg := "Can't get Data for " + sensor
			return errors.New(msg)
		}

		// get Cursor on the Bucket
		c := b.Cursor()

		// fetch latest entry
		maxByte, _ := c.Last()

		// beginning must be positive
		min := BytesToInt(maxByte) - duration
		minBytes := IntBytes(min)
		if min < 0 {
			minBytes = IntBytes(0)
		}

		for k, v := c.Seek(minBytes); k != nil && bytes.Compare(k, maxByte) <= 0; k, v = c.Next() {
			Timestamp := BytesToInt(k)
			Value := BytesToFloat32(v)
			data = append(data, singleData{Timestamp, Value})
		}

		return nil
	})
	if err != nil {
		return
	}
	return
}
