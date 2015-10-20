package sensorserver

import (
	"encoding/json"
	"github.com/0rph3us/astrotime"
	"time"
)

func (s *Sensorserver) GetSunriseAndSunset() string {
	var data []plotBands

	// A month has 744 hours
	now := time.Now().AddDate(0, 0, 1)
	date := time.Now().AddDate(0, -1, 0)

	for ; date.Before(now); date = date.AddDate(0, 0, 1) {
		sunrise := astrotime.CalcSunrise(date, s.conf.Latitude, s.conf.Longitude)
		from := sunrise.Unix() * 1000

		sunset := astrotime.CalcSunset(date, s.conf.Latitude, s.conf.Longitude)
		to := sunset.Unix() * 1000

		// put data into result set
		data = append(data, plotBands{from, to, "#FCFFC5"})
	}
	json, _ := json.Marshal(data)
	return string(json)
}
