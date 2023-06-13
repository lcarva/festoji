package app

import (
	"errors"
	"fmt"
	"time"
)

func Character(today time.Time, config Config) (string, error) {

	for _, rule := range config.Rules {
		var end time.Time
		month := time.Month(rule.Month)

		if rule.Day != 0 {
			end = endOfDay(today, month, rule.Day)
		} else if rule.Weekday != 0 && rule.Week != 0 {
			weekday := time.Weekday(rule.Weekday)
			end = endOfNthWeekdayOfMonth(today, month, rule.Week, weekday)
		} else {
			return "", errors.New(fmt.Sprint(rule.Name, " is not a valid rule"))
		}
		if inSeason(today, end, rule.Span) {
			return rule.Emoji, nil
		}
	}
	return config.Default, nil
}

func inSeason(today time.Time, end time.Time, spanDays int) bool {
	spanSeconds := float64(spanDays) * 24 * 60 * 60
	diffSeconds := end.Sub(today).Seconds()
	return diffSeconds >= 0 && diffSeconds < spanSeconds
}

func endOfDay(today time.Time, month time.Month, day int) time.Time {
	end := endOfDayThisYear(today, month, day)
	// Always get the next occurrence of the date
	if end.Before(today) {
		end = endOfDayNextYear(today, month, day)
	}
	return end
}

func endOfDayThisYear(today time.Time, month time.Month, day int) time.Time {
	end := time.Date(today.Year(), month, day, 23, 59, 59, 0, today.Location())
	return end
}

func endOfDayNextYear(today time.Time, month time.Month, day int) time.Time {
	end := time.Date(today.Year()+1, month, day, 23, 59, 59, 0, today.Location())
	return end
}

// wow this is a terrible name
func endOfNthWeekdayOfMonth(today time.Time, month time.Month, nthWeek int, weekday time.Weekday) time.Time {
	// cannot use the more generic GetEndOfDay function here because we only want to forward
	// to the next year if the *adjusted* end date is in the past.
	end := endOfNthWeekdayOfMonthStrict(endOfDayThisYear(today, month, 1), nthWeek, weekday)
	if end.Before(today) {
		end = endOfNthWeekdayOfMonthStrict(endOfDayNextYear(today, month, 1), nthWeek, weekday)
	}
	return end
}

func endOfNthWeekdayOfMonthStrict(start time.Time, nthWeek int, weekday time.Weekday) time.Time {
	daysToAdd := 0

	if weekDaysDiff := int(weekday - start.Weekday()); weekDaysDiff >= 0 {
		daysToAdd += weekDaysDiff
	} else {
		daysToAdd += int(7 - start.Weekday() + weekday)
	}

	// Advance to whichever nth week
	daysToAdd += (nthWeek - 1) * 7
	return start.AddDate(0, 0, daysToAdd)
}
