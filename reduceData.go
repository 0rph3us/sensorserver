package sensorserver

func (s *Sensorserver) reduceData(data []singleData) []singleData {

	maxPoints := s.conf.MaxPoints

	if len(data) <= maxPoints {
		return data
	}

	min := data[0].Timestamp
	max := data[len(data)-1].Timestamp
	stepInSeconds := (max - min) / maxPoints

	numberOfValues := 0

	reducedData := make([]singleData, maxPoints+1)
	// for iteration on array reducedData
	i := 0

	// init first Element in the Array
	reducedData[0].Timestamp = min
	reducedData[0].Value = 0.0

	for _, element := range data {
		// sum of all values
		reducedData[i].Value += element.Value

		numberOfValues++

		// go to the next interval
		if element.Timestamp > (reducedData[i].Timestamp + stepInSeconds) {
			reducedData[i].Value /= float32(numberOfValues)

			i++
			// init next Element
			reducedData[i].Timestamp = reducedData[i-1].Timestamp + stepInSeconds
			reducedData[i].Value = 0.0

			numberOfValues = 0
		}
	}

	// return the first i elements
	// i can be smaler than maxPoints, if large time gaps in the data array
	return reducedData[:i]
}
