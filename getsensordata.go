package sensorserver

import (
	"github.com/gin-gonic/gin"
	"math"
	"net/http"
)

// return data from a sensor in json format
// This function call reduceData!
func (s *Sensorserver) GetSensorData(c *gin.Context) {
	// maximum Timestamp is MaxInt32
	// fetchLastData return all records, if duration is
	// also MaxInt32
	//  -> MaxInt32 - MaxInt32 = 0 -> get all data since 01.0.1.1970
	const duration int = math.MaxInt32

	pretty := c.Request.URL.Query().Get("pretty")

	sensor := c.Param("name")
	data, err := s.fetchLastData(sensor, duration)

	if err != nil {
		if pretty == "true" {
			c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
	} else {
		if pretty == "true" {
			c.IndentedJSON(http.StatusOK, s.reduceData(data))
		} else {
			c.JSON(http.StatusOK, s.reduceData(data))
		}
	}

}
