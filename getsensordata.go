package sensorserver

import (
	"bytes"
	"errors"
	"github.com/boltdb/bolt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (s *Sensorserver) GetSensorData(c *gin.Context) {

	var data []singleData

	sensor := c.Param("name")

	// return all Values, without parameter
	err := s.boltdb.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(sensor))
		if b == nil {
			msg := "Can't get Data for " + sensor
			return errors.New(msg)
		}

		c := b.Cursor()

		for k, v := c.First(); k != nil; k, v = c.Next() {
			T := BytesToInt(k)
			V := BytesToFloat32(v)
			data = append(data, singleData{T, V})
		}

		return nil
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	} else {
		c.IndentedJSON(http.StatusOK, s.reduceData(data))
	}

}

func (s *Sensorserver) getSensorData(sensor string, duration int) (data []singleData, err error) {

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
			T := BytesToInt(k)
			V := BytesToFloat32(v)
			data = append(data, singleData{T, V})
		}

		return nil
	})
	if err != nil {
		return
	}
	return
}
