package sensorserver

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// render the HTML page
func (s *Sensorserver) Page(c *gin.Context) {
	// configure the sub title
	_ = s.duration(c.Query("duration"))

	c.HTML(http.StatusOK, "index.html", s.conf)
}
