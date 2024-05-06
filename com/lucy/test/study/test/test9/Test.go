package main

import (
	"time"
)

func main() {
	layout := "2006-01-02 15:04:05"
	nextWeekOneAM := getThisWeekMondayMidnight(time.Now())
	nextNextWeekOneAM := getNextWeekdayAtMidnight(time.Now())
	startTime := nextWeekOneAM.Format(layout)
	endTime := nextNextWeekOneAM.Format(layout)
	println(startTime)
	println(endTime)
}
func getNextMonday() time.Time {
	today := time.Now()
	daysUntilMonday := int(time.Monday-today.Weekday()+7) % 7
	nextMonday := today.AddDate(0, 0, daysUntilMonday)
	return nextMonday
}
func getNextWeekdayAtMidnight(currentTime time.Time) time.Time {
	return getThisWeekMondayMidnight(currentTime).Add(time.Hour * 24 * 7)
}

func getMidnightByDay(currentTime time.Time, day int) time.Time {
	tomorrow := currentTime.AddDate(0, 0, day)
	tomorrowMidnight := time.Date(tomorrow.Year(), tomorrow.Month(), tomorrow.Day(), 0, 0, 0, 0, tomorrow.Location())
	return tomorrowMidnight
}
func getThisMonthMidnight(currentTime time.Time) time.Time {
	thisMonth := time.Date(currentTime.Year(), currentTime.Month(), 1, 0, 0, 0, 0, currentTime.Location())
	return thisMonth
}
func getNextMonthMidnight(currentTime time.Time) time.Time {
	nextMonth := time.Date(currentTime.Year(), currentTime.Month()+1, 1, 0, 0, 0, 0, currentTime.Location())
	return nextMonth
}

func getTodayMidnight(currentTime time.Time) time.Time {
	todayMidnight := time.Date(currentTime.Year(), currentTime.Month(), currentTime.Day(), 0, 0, 0, 0, currentTime.Location())

	return todayMidnight
}
func getTomorrowMidnight(currentTime time.Time) time.Time {
	tomorrow := currentTime.AddDate(0, 0, 1)
	tomorrowMidnight := time.Date(tomorrow.Year(), tomorrow.Month(), tomorrow.Day(), 0, 0, 0, 0, tomorrow.Location())
	return tomorrowMidnight
}
func getThisWeekMondayMidnight(currentTime time.Time) time.Time {
	weekday := int(currentTime.Weekday())
	if weekday == 0 {
		weekday = 7
	}
	daysSinceMonday := time.Duration(weekday - 1)
	thisWeekMonday := currentTime.Add(-daysSinceMonday * 24 * time.Hour)
	thisWeekMondayMidnight := time.Date(thisWeekMonday.Year(), thisWeekMonday.Month(), thisWeekMonday.Day(), 0, 0, 0, 0, thisWeekMonday.Location())

	return thisWeekMondayMidnight
}
func getNextMondayMidnight(currentTime time.Time) time.Time {
	daysUntilMonday := (1 + 7 - int(currentTime.Weekday())) % 7
	nextMonday := currentTime.AddDate(0, 0, daysUntilMonday)
	return nextMonday
}
