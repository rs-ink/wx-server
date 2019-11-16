package util

import "time"

func GetBetweenAndTimes(days int, targetDate ...time.Time) (start, end time.Time) {
	var tt time.Time
	if len(targetDate) > 0 {
		tt = targetDate[0]
	} else {
		tt = time.Now()
	}
	tt = time.Date(tt.Year(), tt.Month(), tt.Day(), 0, 0, 0, 0, time.Local)
	if days < 0 {
		start = tt.AddDate(0, 0, days)
		end = tt.AddDate(0, 0, 1)
	} else {
		start = tt
		end = tt.AddDate(0, 0, days+1)
	}
	return
}
