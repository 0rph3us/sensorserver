package sensorserver

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func (s *Sensorserver) Page(c *gin.Context) {

	// configure sub title
	duration := c.Query("duration")
	s.conf.Duration = duration
	switch duration {
	case "3h":
		{
			s.conf.SubTitle = "der letzten 3 Stunden"
		}
	case "24h":
		{
			s.conf.SubTitle = "der letzten 24 Stunden"
		}
	case "1w":
		{
			s.conf.SubTitle = "der letzten Woche"
		}
	case "4w":
		{
			s.conf.SubTitle = "des letzten Monats"
		}
	default:
		{
			s.conf.SubTitle = "der letzten 24 Stunden"
		}
	}

	c.HTML(http.StatusOK, "index.html", s.conf)

}
