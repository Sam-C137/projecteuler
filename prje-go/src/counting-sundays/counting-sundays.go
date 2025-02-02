package counting_sundays

import (
	"time"
)

func Run() int {
	year := 1901
	dayOfMonth := 1
	month := 1
	dayOfWeek := 2
	var sundays int

	for year <= 2000 {
		if dayOfMonth == 1 && dayOfWeek == 7 {
			sundays++
		}

		dayOfWeek++

		if dayOfWeek > 7 {
			dayOfWeek = 1
		}

		dayOfMonth++

		daysInMonth := getDaysInMonth(month, year)
		if dayOfMonth > daysInMonth {
			dayOfMonth = 1
			month++
		}

		if month > 12 {
			month = 1
			year++
		}
	}

	return sundays
}

func getDaysInMonth(month, year int) int {
	switch month {
	case 2:
		if year%4 == 0 {
			return 29
		} else {
			return 28
		}
	case 4, 6, 9, 11:
		return 30
	default:
		return 31
	}
}

// ----------------------------------------------

func RunBetter() int {
	var sundays int
	start := time.Date(1901, 1, 1, 0, 0, 0, 0, time.UTC)
	end := time.Date(2000, 12, 31, 0, 0, 0, 0, time.UTC)

	for curr := start; curr.Before(end) || curr.Equal(end); curr = curr.AddDate(0, 1, 0) {
		if curr.Day() == 1 && curr.Weekday() == time.Sunday {
			sundays++
		}
	}

	return sundays
}
