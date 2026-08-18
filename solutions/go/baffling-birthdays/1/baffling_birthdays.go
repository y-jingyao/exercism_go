package bafflingbirthdays

import (
	"math/rand"
	"time"
)

const numTrials = 10000

func SharedBirthday(dates []time.Time) bool {
	seen := make(map[[2]int]bool)
	for _, d := range dates {
		key := [2]int{int(d.Month()), d.Day()}
		if seen[key] {
			return true
		}
		seen[key] = true
	}
	return false
}

func RandomBirthdates(size int) []time.Time {
	dates := make([]time.Time, size)
	base := time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC)

	for i := 0; i < size; i++ {
		dayOfYear := rand.Intn(365) 
		birthday := base.AddDate(0, 0, dayOfYear)

		year := rand.Intn(101) + 1900
		for isLeapYear(year) {
			year = rand.Intn(101) + 1900
		}

		dates[i] = time.Date(year, birthday.Month(), birthday.Day(), 0, 0, 0, 0, time.UTC)
	}
	return dates
}

func EstimatedProbability(size int) float64 {
	if size <= 1 {
		return 0.0
	}

	matches := 0
	for i := 0; i < numTrials; i++ {
		dates := RandomBirthdates(size)
		if SharedBirthday(dates) {
			matches++
		}
	}

	return float64(matches) / float64(numTrials) * 100.0
}
