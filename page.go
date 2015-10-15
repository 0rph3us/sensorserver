package sensorserver

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func (s *Sensorserver) Page(c *gin.Context) {

	s.conf.SubTitle = "der letzten 24 Stunden"

	c.HTML(http.StatusOK, "index.html", s.conf)

}
