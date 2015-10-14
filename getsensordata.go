package sensorserver

import (
	"github.com/boltdb/bolt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (s *Sensorserver) GetSensorData(c *gin.Context) {
	var data []putdata

	sensor := []byte(c.Param("name"))

	// return all Values, without parameter
	err := s.boltdb.View(func(tx *bolt.Tx) error {
		b := tx.Bucket(sensor)
		b.ForEach(func(k, v []byte) error {
			data = append(data, putdata{BytesToInt(k), BytesToFloat32(v)})
			return nil
		})
		return nil
	})

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
	} else {
		c.JSON(http.StatusOK, data)
	}

}
