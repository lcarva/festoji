package app

import (
    "errors"
    "fmt"
    "time"
)

func GetCharacter(today time.Time, config FestojiConfig) (string, error) {

    for _, rule := range config.Rules {
        var end time.Time
        month := time.Month(rule.Month)

        if rule.Day != 0 {
            end = GetEndOfDay(today, month, rule.Day)
        } else if rule.Weekday != 0 && rule.Week != 0 {
            weekday := time.Weekday(rule.Weekday)
            end = GetEndOfNthWeekdayOfMonth(today, month, rule.Week, weekday)
        } else {
            return "", errors.New(fmt.Sprint(rule.Name, " is not a valid rule"))
        }
        if InSeason(today, end, rule.Span) {
            return rule.Emoji, nil
        }
    }
    return config.Default, nil
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
