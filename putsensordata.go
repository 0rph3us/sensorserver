package sensorserver

import (
	"github.com/boltdb/bolt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (s *Sensorserver) PutSensorData(c *gin.Context) {
	var data []singleData

	sensor := []byte(c.Param("name"))

	if c.BindJSON(&data) == nil {
		// get array of singleData
		err := s.boltdb.Batch(func(tx *bolt.Tx) error {

			// create the bucket
			bucket, err := tx.CreateBucketIfNotExists(sensor)
			if err != nil {
				return err
			}

			// put all data in the bucket
			for _, d := range data {
				if err := bucket.Put(IntBytes(d.Timestamp), Float32Bytes(d.Value)); err != nil {
					return err
				}
			}
			return nil
		})
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		} else {
			msg := "Put all Data in bucket " + string(sensor)
			c.JSON(http.StatusOK, gin.H{"status": msg, "elements": len(data)})
		}

	} else {
		// can't parse data
		c.JSON(http.StatusBadRequest, gin.H{"error": "Mal formed JSON"})
	}
}
