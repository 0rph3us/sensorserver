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

	durationInSeconds := s.duration(c.Query("duration"))

	f := make(map[string]interface{})
	sensors := []string{"tmp_dth22", "p_sea", "humidity"}
	for _, sensor := range sensors {
		data, err := s.fetchLastData(sensor, durationInSeconds)
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
