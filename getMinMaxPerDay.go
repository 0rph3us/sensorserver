package sensorserver

import (
	"math"
	"time"
)

// return the minimum, maximum and average values per day.
func (s *Sensorserver) getMinMaxPerDay(points []singleData) []minMaxData {

	location, _ := time.LoadLocation("Local")

	// 12 hours in seconds
	halfDayInSeconds := 43200

	// calculate the minimum Timestamp
	start := time.Unix(int64(points[0].Timestamp), 0)
	year, month, day := start.Date()
	startTime := time.Date(year, month, day, 12, 0, 0, 0, location)
	minTimestamp := int(startTime.Unix())

	maxTimestamp := points[len(points)-1].Timestamp

	// number of days
	days := (maxTimestamp - minTimestamp) / (halfDayInSeconds * 2)

	minMaxData := make([]minMaxData, days+1)
	// for iteration on array minMaxData
	i := 0

	// init first Element in the Array
	minMaxData[0].Timestamp = minTimestamp
	minMaxData[0].MinValue = math.MaxFloat32
	minMaxData[0].MaxValue = -math.MaxFloat32
	minMaxData[0].AvgValue = 0.0

	numberOfValues := 0

	for _, point := range points {
		// find minimum value
		if minMaxData[i].MinValue > point.Value {
			minMaxData[i].MinValue = point.Value
		}

		// find maximum value
		if minMaxData[i].MaxValue < point.Value {
			minMaxData[i].MaxValue = point.Value
		}

		// sum of all values
		minMaxData[i].AvgValue += point.Value

		numberOfValues++

		// go to the next day
		if point.Timestamp > (minMaxData[i].Timestamp + halfDayInSeconds) {

			minMaxData[i].AvgValue /= float32(numberOfValues)

			i++
			// init next Element
			minMaxData[i].Timestamp = minMaxData[i-1].Timestamp + (2 * halfDayInSeconds)
			minMaxData[i].MinValue = math.MaxFloat32
			minMaxData[i].MaxValue = -math.MaxFloat32
			minMaxData[i].AvgValue = 0.0

			numberOfValues = 0
		}
	}

	return minMaxData
}
