package sensorserver

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func (s *Sensorserver) Stats(c *gin.Context) {
	prev := s.boltdb.Stats()

	pretty := c.Request.URL.Query().Get("pretty")

	if pretty == "true" {
		c.IndentedJSON(http.StatusOK, prev)
	} else {
		c.JSON(http.StatusOK, prev)
	}
}
