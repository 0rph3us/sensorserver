package sensorserver

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func (s *Sensorserver) GetChart(c *gin.Context) {

	// configure sub title
	duration := c.Query("duration")
	var durationInSeconds int
	s.conf.Duration = duration
	switch duration {
	case "3h":
		{
			s.conf.SubTitle = "der letzten 3 Stunden"
			durationInSeconds = 3600 * 3
		}
	case "24h":
		{
			s.conf.SubTitle = "der letzten 24 Stunden"
			durationInSeconds = 3600 * 24
		}
	case "1w":
		{
			s.conf.SubTitle = "der letzten Woche"
			durationInSeconds = 3600 * 24 * 7
		}
	case "4w":
		{
			s.conf.SubTitle = "des letzten Monats"
			durationInSeconds = 3600 * 24 * 7 * 4
		}
	default:
		{
			s.conf.SubTitle = "der letzten 24 Stunden"
			durationInSeconds = 3600 * 24
		}
	}

	f := make(map[string][]highchartData)
	f["dht22"], _ = s.getSensorData("tmp_dth22", durationInSeconds)
	f["p_sea"], _ = s.getSensorData("p_sea", durationInSeconds)
	f["humi"], _ = s.getSensorData("humidity", durationInSeconds)

	c.Header("Content-Type", "application/javascript")
	c.HTML(http.StatusOK, "chart.js", f)

}
