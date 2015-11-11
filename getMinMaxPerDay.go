package sensorserver

import (
	"container/heap"
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

	minMaxData := make([]minMaxData, days+2)
	// for iteration on array minMaxData
	i := 0

	// init first Element in the Array
	minMaxData[0].Timestamp = minTimestamp
	minMaxData[0].AvgValue = 0.0

	// init empty Heap
	min := &MinFloat32Heap{}
	heap.Init(min)

	for _, point := range points {

		// use Heap for Minimum and Maximum
		heap.Push(min, point.Value)

		// sum of all values
		minMaxData[i].AvgValue += point.Value

		// go to the next day
		if point.Timestamp > (minMaxData[i].Timestamp + halfDayInSeconds) {

			p99 := int(float32(min.Len()) * 0.99)
			min_p99 := min.Len() - p99
			if min.Len() > 1 {
				minMaxData[i].MinValue = (*min)[p99]
				minMaxData[i].MaxValue = (*min)[min_p99]
			} else {
				minMaxData[i].MinValue = (*min)[0]
				minMaxData[i].MaxValue = (*min)[0]
			}

			numberOfValues := min.Len() + 1
			minMaxData[i].AvgValue /= float32(numberOfValues)

			min = &MinFloat32Heap{}
			heap.Init(min)

			i++
			// init next Element
			minMaxData[i].Timestamp = minMaxData[i-1].Timestamp + (2 * halfDayInSeconds)
			minMaxData[i].AvgValue = 0.0
		}
	}

	return minMaxData[:len(minMaxData)-2]
}
