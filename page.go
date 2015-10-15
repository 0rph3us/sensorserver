package sensorserver

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func (s *Sensorserver) Page(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", s.conf)

}
