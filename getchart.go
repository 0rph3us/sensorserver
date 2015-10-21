package sensorserver

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"text/template"
)

func (s *Sensorserver) GetChart(c *gin.Context) {

	// Load Template
	//  Can't use c.HTML(), because the html template function escape "
	t, _ := template.ParseFiles("templates/chart.js")

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

	f := make(map[string]interface{})
	sensors := []string{"tmp_dth22", "p_sea", "humidity"}
	for _, sensor := range sensors {
		data, err := s.getSensorData(sensor, durationInSeconds)
		if err != nil {
			log.Println(err.Error())
			c.String(http.StatusInternalServerError, err.Error())
			return
		}
		f[sensor] = s.reduceData(data)
	}

	f["plotBands"] = s.GetSunriseAndSunset()

	c.Header("Content-Type", "application/javascript")
	t.Execute(c.Writer, f)
}
