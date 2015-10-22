package sensorserver

func (s *Sensorserver) reduceData(data []singleData) []singleData {

	maxPoints := s.conf.MaxPoints

	if len(data) <= maxPoints {
		return data
	}

	min := data[0].Timestamp
	max := data[len(data)-1].Timestamp
	step := (max - min) / maxPoints
	step_h := step / 2
	i := 0
	numberOfValues := 0

	newData := make([]singleData, maxPoints+1)
	v := singleData{min + step_h, 0}
	newData[0].Timestamp = min + step_h

	for _, value := range data {

		v.Value += value.Value
		numberOfValues++

		if value.Timestamp >= (newData[i].Timestamp + step_h) {
			v.Value /= float32(numberOfValues)
			newData[i] = v
			v.Timestamp = newData[i].Timestamp + step
			v.Value = 0
			newData[i+1].Timestamp = newData[i].Timestamp + step
			i++
			numberOfValues = 0
		}
	}

	// remove the last Element
	return newData[:len(newData)-1]
}
