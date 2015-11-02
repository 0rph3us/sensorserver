package sensorserver

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// render the HTML page
func (s *Sensorserver) Page(c *gin.Context) {
	// configure the sub title
	_ = s.duration(c.Query("duration"))

	s.conf.Type = c.Query("type")
	if s.conf.Type == "" {
		s.conf.Type = "single"
	}

	c.HTML(http.StatusOK, "index.html", s.conf)
}
