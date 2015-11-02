package sensorserver

// return the duration in seconds and update the duration in config and set the subtitle
func (s *Sensorserver) duration(duration string) (durationInSeconds int) {
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
	case "12w":
		{
			s.conf.SubTitle = "der letzten 3 Monate"
			durationInSeconds = 3600 * 24 * 7 * 12
		}
	default:
		{
			s.conf.SubTitle = "der letzten 24 Stunden"
			durationInSeconds = 3600 * 24
			s.conf.Duration = "24h"
		}
	}
	return
}
