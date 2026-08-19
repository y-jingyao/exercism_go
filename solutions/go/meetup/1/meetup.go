package meetup

import "time"

type WeekSchedule int

const (
	First  WeekSchedule = iota
	Second
	Third
	Fourth
	Last
	Teenth
)

func Day(wSched WeekSchedule, wDay time.Weekday, month time.Month, year int) int {
	var matchingDays []int

	for day := 1; day <= 31; day++ {
		date := time.Date(year, month, day, 0, 0, 0, 0, time.UTC)
		
		if date.Month() != month {
			break
		}

		if date.Weekday() == wDay {
			matchingDays = append(matchingDays, day)
		}
	}

	switch wSched {
	case First:
		return matchingDays[0]
	case Second:
		return matchingDays[1]
	case Third:
		return matchingDays[2]
	case Fourth:
		return matchingDays[3]
	case Last:
		return matchingDays[len(matchingDays)-1]
	case Teenth:
		for _, d := range matchingDays {
			if d >= 13 && d <= 19 {
				return d
			}
		}
	}

	return -1
}
