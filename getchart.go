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

	s.conf.Type = c.Param("type")

	durationInSeconds := s.duration(c.Query("duration"))

	f := make(map[string]interface{})
	for _, sensor := range s.conf.Sensors {
		data, err := s.fetchLastData(sensor, durationInSeconds)
		if err != nil {
			log.Println(err.Error())
			c.String(http.StatusInternalServerError, err.Error())
			return
		}

		if s.conf.Type == "single" {
			f[sensor] = s.reduceData(data)
		} else {
			f[sensor] = s.getMinMaxPerDay(data)
		}
	}

	f["plotBands"] = s.GetSunriseAndSunset()
	f["type"] = s.conf.Type

	c.Header("Content-Type", "application/javascript; charset=utf-8")
	t.Execute(c.Writer, f)
}
