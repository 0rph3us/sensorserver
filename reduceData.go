package sensorserver

func (s *Sensorserver) reduceData(data []highchartData) []highchartData {

	maxPoints := s.conf.MaxPoints

	if len(data) <= maxPoints {
		return data
	}

	min := data[0].T
	max := data[len(data)-1].T
	step := (max - min) / int64(maxPoints)
	step_h := step / 2
	i := 0
	numberOfValues := 0

	newData := make([]highchartData, maxPoints+1)
	v := highchartData{min + step_h, 0}
	newData[0].T = min + step_h

	for _, value := range data {

		v.V += value.V
		numberOfValues++

		if value.T >= (newData[i].T + step_h) {
			v.V /= float32(numberOfValues)
			newData[i] = v
			v.T = newData[i].T + step
			v.V = 0
			newData[i+1].T = newData[i].T + step
			i++
			numberOfValues = 0
		}
	}

	// remove the last Element
	return newData[:len(newData)-1]
}
