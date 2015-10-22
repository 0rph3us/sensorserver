package sensorserver

import (
    "github.com/boltdb/bolt"
    "errors"
    "bytes"
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

		c := b.Cursor()

		// fetch latest entry
		max, _ := c.Last()

		min := IntBytes(BytesToInt(max) - duration)

		for k, v := c.Seek(min); k != nil && bytes.Compare(k, max) <= 0; k, v = c.Next() {
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
