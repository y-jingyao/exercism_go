package swiftscheduling

import (
	"strconv"
	"time"
)

func DeliveryDate(start, delivery string) string {
	// 注意：输入格式是 "2012-02-13T09:00:00"，用 T 分隔
	t, err := time.Parse("2006-01-02T15:04:05", start)
	if err != nil {
		return ""
	}

	switch delivery {
	case "NOW":
		result := t.Add(2 * time.Hour)
		return result.Format("2006-01-02T15:04:05")

	case "ASAP":
		hour := t.Hour()
		if hour < 13 {
			result := time.Date(t.Year(), t.Month(), t.Day(), 17, 0, 0, 0, time.UTC)
			return result.Format("2006-01-02T15:04:05")
		}
		tomorrow := t.AddDate(0, 0, 1)
		result := time.Date(tomorrow.Year(), tomorrow.Month(), tomorrow.Day(), 13, 0, 0, 0, time.UTC)
		return result.Format("2006-01-02T15:04:05")

	case "EOW":
		weekday := t.Weekday()
		if weekday >= time.Monday && weekday <= time.Wednesday {
			daysUntilFriday := int(time.Friday - weekday)
			friday := t.AddDate(0, 0, daysUntilFriday)
			result := time.Date(friday.Year(), friday.Month(), friday.Day(), 17, 0, 0, 0, time.UTC)
			return result.Format("2006-01-02T15:04:05")
		}
		daysUntilSunday := int(time.Sunday - weekday)
		if daysUntilSunday < 0 {
			daysUntilSunday += 7
		}
		sunday := t.AddDate(0, 0, daysUntilSunday)
		result := time.Date(sunday.Year(), sunday.Month(), sunday.Day(), 20, 0, 0, 0, time.UTC)
		return result.Format("2006-01-02T15:04:05")

	default:
		if len(delivery) > 1 && delivery[len(delivery)-1] == 'M' {
			n, err := strconv.Atoi(delivery[:len(delivery)-1])
			if err != nil || n < 1 || n > 12 {
				return ""
			}
			return handleMonth(t, n)
		}

		if len(delivery) > 1 && delivery[0] == 'Q' {
			n, err := strconv.Atoi(delivery[1:])
			if err != nil || n < 1 || n > 4 {
				return ""
			}
			return handleQuarter(t, n)
		}

		return ""
	}
}

func handleMonth(t time.Time, n int) string {
	currentMonth := int(t.Month())
	year := t.Year()

	if currentMonth >= n {
		year++
	}

	month := time.Month(n)
	day := 1
	for {
		candidate := time.Date(year, month, day, 8, 0, 0, 0, time.UTC)
		wd := candidate.Weekday()
		if wd >= time.Monday && wd <= time.Friday {
			return candidate.Format("2006-01-02T15:04:05")
		}
		day++
	}
}

func handleQuarter(t time.Time, n int) string {
	currentQuarter := (int(t.Month())-1)/3 + 1
	year := t.Year()

	if currentQuarter > n {
		year++
	}

	lastMonthOfQuarter := time.Month(n * 3)

	firstDayOfNextMonth := time.Date(year, lastMonthOfQuarter+1, 1, 8, 0, 0, 0, time.UTC)
	lastDay := firstDayOfNextMonth.AddDate(0, 0, -1)

	for {
		wd := lastDay.Weekday()
		if wd >= time.Monday && wd <= time.Friday {
			result := time.Date(lastDay.Year(), lastDay.Month(), lastDay.Day(), 8, 0, 0, 0, time.UTC)
			return result.Format("2006-01-02T15:04:05")
		}
		lastDay = lastDay.AddDate(0, 0, -1)
	}
}