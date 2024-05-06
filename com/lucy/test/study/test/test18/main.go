package main

import "time"

func main() {
	layout := "2006-01-02 15:04:05"
	now := time.Now()
	now = getTodayMidnight(now)
	nextNextWeekOneAM := getNextMondayMidnight(now)
	endTime := nextNextWeekOneAM.Format(layout)
	println(endTime)
	format := getNextWeekdayAtMidnight(time.Now()).Format(layout)
	println(format)

}
func getTodayMidnight(currentTime time.Time) time.Time {
	todayMidnight := time.Date(currentTime.Year(), currentTime.Month(), currentTime.Day(), 0, 0, 0, 0, currentTime.Location())

	return todayMidnight
}
func getNextMondayMidnight(currentTime time.Time) time.Time {
	daysUntilMonday := (1 + 7 - int(currentTime.Weekday())) % 7
	nextMonday := currentTime.AddDate(0, 0, daysUntilMonday)
	return nextMonday
}
func getNextWeekdayAtMidnight(currentTime time.Time) time.Time {
	return getThisWeekMondayMidnight(currentTime).Add(time.Hour * 24 * 7)
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
