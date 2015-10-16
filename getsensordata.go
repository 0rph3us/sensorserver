package sensorserver

import (
	"errors"
	"github.com/boltdb/bolt"
	"github.com/gin-gonic/gin"
	"net/http"
	"sort"
)

type ByTimestamp []highchartData

func (a ByTimestamp) Len() int           { return len(a) }
func (a ByTimestamp) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByTimestamp) Less(i, j int) bool { return a[i].T < a[j].T }

func (s *Sensorserver) GetSensorData(c *gin.Context) {

	var data []highchartData

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
			T := int64(BytesToInt(k) * 1000) // highcharts need a int32 as Timestamp
			V := BytesToFloat32(v)
			data = append(data, highchartData{T, V})
		}

		return nil
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	} else {
		sort.Sort(ByTimestamp(data))

		c.IndentedJSON(http.StatusOK, data[len(data)-400:])
	}

}
