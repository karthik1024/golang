package booking

import (
	"fmt"
	"time"
)

// Schedule returns a time.Time from a string containing a date.
func Schedule(date string) time.Time {
	t, e := time.Parse("1/2/2006 15:04:05", date)
	if e != nil {
		panic(e)
	}
	return t
}

// HasPassed returns whether a date has passed.
func HasPassed(date string) bool {
	t, e := time.Parse("January 2, 2006 15:04:05", date)
	if e != nil {
		panic(e)
	}
	return t.Before(time.Now())
}

// IsAfternoonAppointment returns whether a time is in the afternoon.
func IsAfternoonAppointment(date string) bool {
	t, e := time.Parse("Monday, January 2, 2006 15:04:05", date)
	if e != nil {
		panic(e)
	}
	return t.Hour() >= 12 && t.Hour() < 18
}

// Description returns a formatted string of the appointment time.
func Description(date string) string {
	t, e := time.Parse("1/2/2006 15:04:05", date)
	if e != nil {
		panic(e)
	}
	return t.Format("You have an appointment on Monday, January 2, 2006, at 15:04.")
}

// AnniversaryDate returns a Time with this year's anniversary.
func AnniversaryDate() time.Time {
	date := fmt.Sprintf("%d-09-15", time.Now().Year())
	t, e := time.Parse("2006-01-02", date)
	if e != nil {
		panic(e)
	}
	return t
}
