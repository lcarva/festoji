package app

import (
	"testing"
	"time"
)

func TestGetEndOfDay(t *testing.T) {
	date := time.Date(2021, time.August, 1, 0, 0, 0, 0, time.UTC)

	expectedAfter := time.Date(2021, time.September, 1, 23, 59, 59, 0, time.UTC)
	after := GetEndOfDay(date, time.September, 1)
	if !after.Equal(expectedAfter) {
		t.Error("After date", after, "is not as expected", expectedAfter)
	}

	expectedBefore := time.Date(2022, time.July, 1, 23, 59, 59, 0, time.UTC)
	before := GetEndOfDay(date, time.July, 1)
	if !before.Equal(expectedBefore) {
		t.Error("Before date", before, "is not as expected", expectedBefore)
	}
}

func TestGetEndOfNthWeekdayOfMonth_NextMonthThisWeekday(t *testing.T) {
	// The month has not already passed, and the first of the month is before the first occurence
	// of the requested weekday.
	date := time.Date(2021, time.August, 1, 0, 0, 0, 0, time.UTC)
	expectedNextDate := time.Date(2021, time.November, 25, 23, 59, 59, 0, time.UTC)
	nextDate := GetEndOfNthWeekdayOfMonth(date, time.November, 4, time.Thursday)
	if !nextDate.Equal(expectedNextDate) {
		t.Error("Date", nextDate, "is not as expected", expectedNextDate)
	}
}

func TestGetEndOfNthWeekdayOfMonth_NextMonthNextWeek(t *testing.T) {
	// The month has not already passed, and the first of the month is after the first occurence
	// of the requested weekday.
	date := time.Date(2021, time.August, 1, 0, 0, 0, 0, time.UTC)
	expectedNextDate := time.Date(2021, time.September, 13, 23, 59, 59, 0, time.UTC)
	nextDate := GetEndOfNthWeekdayOfMonth(date, time.September, 2, time.Monday)
	if !nextDate.Equal(expectedNextDate) {
		t.Error("Date", nextDate, "is not as expected", expectedNextDate)
	}
}

func TestGetEndOfNthWeekdayOfMonth_PreviousMonthThisWeekday(t *testing.T) {
	// The month has already passed, and the first of the month is before the first occurrence of
	// the requested weekday.
	date := time.Date(2020, time.December, 1, 0, 0, 0, 0, time.UTC)
	expectedNextDate := time.Date(2021, time.November, 25, 23, 59, 59, 0, time.UTC)
	nextDate := GetEndOfNthWeekdayOfMonth(date, time.November, 4, time.Thursday)
	if !nextDate.Equal(expectedNextDate) {
		t.Error("Date", nextDate, "is not as expected", expectedNextDate)
	}
}

func TestGetEndOfNthWeekdayOfMonth_PreviousMonthNextWeek(t *testing.T) {
	// The month has already passed, and the first of the month is after the first occurence of the
	// requested weekday.
	date := time.Date(2020, time.December, 1, 0, 0, 0, 0, time.UTC)
	expectedNextDate := time.Date(2021, time.September, 13, 23, 59, 59, 0, time.UTC)
	nextDate := GetEndOfNthWeekdayOfMonth(date, time.September, 2, time.Monday)
	if !nextDate.Equal(expectedNextDate) {
		t.Error("Date", nextDate, "is not as expected", expectedNextDate)
	}
}

func TestInSeason(t *testing.T) {
	todayDate := time.Date(2021, time.December, 12, 0, 0, 0, 0, time.UTC)
	endDate := time.Date(2021, time.December, 26, 0, 0, 0, 0, time.UTC)
	span := 14

	if InSeason(todayDate, endDate, span) {
		t.Error(todayDate, "should not be in season:", endDate, span)
	}
	for i := 0; i < span; i++ {
		todayDate = todayDate.AddDate(0, 0, 1)
		if !InSeason(todayDate, endDate, span) {
			t.Error(todayDate, "should not be in season:", endDate, span)
		}
	}
	todayDate = todayDate.AddDate(0, 0, 1)
	if InSeason(todayDate, endDate, span) {
		t.Error(todayDate, "should not be in season:", endDate, span)
	}
}
