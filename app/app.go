package app

import (
	"time"
)

func GetCharacter(today time.Time) string {
    if InXmasSeason(today){
        return "🎄"
    }
    if InThanksgivingSeason(today) {
        return "🦃"
    }
    if InNewYearsSeason(today) {
        return "🍾"
    }
    if InValentineSeason(today) {
        return "❤️"
    }
    return "🐚"
}

func InXmasSeason(today time.Time) bool {
    end := GetEndOfDay(today, time.December, 25)
    return InSeason(today, end, 14)
}

func InThanksgivingSeason(today time.Time) bool {
    // Get the end of the day for the fourth Thursday in November
    end := GetEndOfNthWeekdayOfMonth(today, time.November, 3, time.Thursday)
    return InSeason(today, end, 7)
}

func InNewYearsSeason(today time.Time) bool {
    // This is a deliberate change from previous versions where December 31 was used.
    end := GetEndOfDay(today, time.January, 1)
    return InSeason(today, end, 6)
}

func InValentineSeason(today time.Time) bool {
    end := GetEndOfDay(today, time.February, 14)
    return InSeason(today, end, 7)
}

func InSeason(today time.Time, end time.Time, spanDays int) bool {
    spanSeconds := float64(spanDays) * 24 * 60 * 60
    diffSeconds := end.Sub(today).Seconds()
    return diffSeconds >= 0 && diffSeconds < spanSeconds
}

func GetEndOfDay(today time.Time, month time.Month, day int) time.Time {
    end := time.Date(today.Year(), month, day, 23, 59, 59, 0, today.Location())
    // Always get the next occurrence of the date
    if end.Before(today) {
        end = time.Date(today.Year()+1, month, day, 23, 59, 59, 0, today.Location())
    }
    return end
}


// wow this is a terrible name
func GetEndOfNthWeekdayOfMonth(today time.Time, month time.Month, nthWeek int, weekday time.Weekday) time.Time {
    end := GetEndOfDay(today, month, 1)
    daysToAdd := 0

    if weekDaysDiff := int(weekday - end.Weekday()); weekDaysDiff >= 0 {
        daysToAdd += weekDaysDiff
    } else {
        daysToAdd += int(7 - end.Weekday() + weekday)
    }

    // Advance to whichever nth week
    daysToAdd += nthWeek * 7
    return end.AddDate(0, 0, daysToAdd)
}
