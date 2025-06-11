package util

import "time"

func GetLastMidnightFrom(from time.Time) time.Time {
	year, month, day := from.Date()
	lastMidnight := time.Date(year, month, day, 0, 0, 0, 0, from.Location())
	return lastMidnight
}
