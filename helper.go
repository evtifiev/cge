package main

import (
	"strconv"
	"time"
)

func DateFormatXml(dateOfBirth, dateOpenCase string) string {
	dt1, err := time.Parse("2006-01-02", dateOfBirth)
	if err != nil {
		panic(err)
	}
	dt2, err := time.Parse("2006-01-02", dateOpenCase)
	if err != nil {
		panic(err)
	}
	_, years, months, days, _, _, _, _ := Elapsed(dt1, dt2)
	if years >= 1 && years < 4 {
		return strconv.Itoa(years*10000 + months*100)
	} else if years < 1 {
		if months >= 1 {
			return strconv.Itoa(months * 100)
		} else {
			return strconv.Itoa(days + 100)
		}
	} else {
		return strconv.Itoa(years * 10000)
	}
}
func DaysIn(year int, month time.Month) int {
	return time.Date(year, month+1, 0, 0, 0, 0, 0, time.UTC).Day()
}

func Elapsed(from, to time.Time) (inverted bool, years, months, days, hours, minutes, seconds, nanoseconds int) {
	if from.Location() != to.Location() {
		to = to.In(to.Location())
	}

	inverted = false
	if from.After(to) {
		inverted = true
		from, to = to, from
	}

	y1, M1, d1 := from.Date()
	y2, M2, d2 := to.Date()

	h1, m1, s1 := from.Clock()
	h2, m2, s2 := to.Clock()

	ns1, ns2 := from.Nanosecond(), to.Nanosecond()

	years = y2 - y1
	months = int(M2 - M1)
	days = d2 - d1

	hours = h2 - h1
	minutes = m2 - m1
	seconds = s2 - s1
	nanoseconds = ns2 - ns1

	if nanoseconds < 0 {
		nanoseconds += 1e9
		seconds--
	}
	if seconds < 0 {
		seconds += 60
		minutes--
	}
	if minutes < 0 {
		minutes += 60
		hours--
	}
	if hours < 0 {
		hours += 24
		days--
	}
	if days < 0 {
		days += DaysIn(y2, M2-1)
		months--
	}
	if months < 0 {
		months += 12
		years--
	}
	return
}
