package sensorserver

import (
	"github.com/boltdb/bolt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (s *Sensorserver) PutMultiSensorData(c *gin.Context) {
	var data []multidata
	numOfData := 0

	if c.BindJSON(&data) == nil {
		err := s.boltdb.Update(func(tx *bolt.Tx) error {

			// iterate over all Timestamps
			for _, multi := range data {
				timestamp := IntBytes(multi.Timestamp)
				// iterate over the sensors
				for sensor, value := range multi.Sensors {

					numOfData++

					// create the bucket
					bucket, err := tx.CreateBucketIfNotExists([]byte(sensor))
					if err != nil {
						return err
					}

					// put Value in the Bucket
					if err := bucket.Put(timestamp, Float32Bytes(value)); err != nil {
						return err
					}
				}
			}
			return nil
		})
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		} else {
			msg := "Put all Data in the database"
			c.JSON(http.StatusOK, gin.H{"status": msg, "elements": numOfData})
		}
	} else {
		// can't parse data
		c.JSON(http.StatusBadRequest, gin.H{"error": "Mal formed JSON"})
	}
}
