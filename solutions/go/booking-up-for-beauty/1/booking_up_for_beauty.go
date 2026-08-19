package booking

import "time"

// Schedule returns a time.Time from a string containing a date.
// 输入格式: "7/25/2019 13:45:00"
func Schedule(date string) time.Time {
	layout := "1/2/2006 15:04:05"
	res, _ := time.Parse(layout, date)
	return res
}

// HasPassed returns whether a date has passed.
// 输入格式: "July 25, 2019 13:45:00"
func HasPassed(date string) bool {
	layout := "January 2, 2006 15:04:05"
	t, _ := time.Parse(layout, date)
	return time.Now().After(t)
}

// IsAfternoonAppointment returns whether a time is in the afternoon.
// 输入格式: "Thursday, July 25, 2019 13:45:00"
func IsAfternoonAppointment(date string) bool {
	layout := "Monday, January 2, 2006 15:04:05"
	t, _ := time.Parse(layout, date)
	h := t.Hour()
	return h >= 12 && h < 18
}

// Description returns a formatted string of the appointment time.
// 输入格式: "7/25/2019 13:45:00"
func Description(date string) string {
	t := Schedule(date)
	formatted := t.Format("Monday, January 2, 2006, at 15:04.")
	return "You have an appointment on " + formatted
}

// AnniversaryDate returns a Time with this year's anniversary.
func AnniversaryDate() time.Time {
	return time.Date(time.Now().Year(), time.September, 15, 0, 0, 0, 0, time.UTC)
}