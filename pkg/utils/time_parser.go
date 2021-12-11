package utils

import (
	"errors"
	"strconv"
	"strings"
	"time"
)

func ParseDate(date string) (*time.Time, error) {
	dates := strings.Split(date, ".")
	day, err := strconv.Atoi(dates[0])
	if err != nil {
		return nil, err
	}

	month, err := strconv.Atoi(dates[1])
	if err != nil {
		return nil, err
	}

	year, err := strconv.Atoi(dates[2])
	if err != nil {
		return nil, err
	}

	t := time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.Local)
	return &t, nil
}

func ParseTime(date *time.Time, t string) (start time.Time, end time.Time, err error) {
	ts := strings.Split(t, "-")
	ts1 := strings.Split(ts[0], ":")
	ts2 := strings.Split(ts[1], ":")

	if len(ts) != 2 || len(ts1) != 2 || len(ts2) != 2 {
		return time.Time{}, time.Time{}, errors.New("invalid time")
	}

	tsStartHours, err := strconv.Atoi(ts1[0])
	if err != nil {
		return time.Time{}, time.Time{}, err
	}
	tsStartMinutes, err := strconv.Atoi(ts1[1])
	if err != nil {
		return time.Time{}, time.Time{}, err
	}
	tsEndHours, err := strconv.Atoi(ts2[0])
	if err != nil {
		return time.Time{}, time.Time{}, err
	}
	tsEndMinutes, err := strconv.Atoi(ts2[1])
	if err != nil {
		return time.Time{}, time.Time{}, err
	}

	s := date.Format(time.ANSIC)
	start, _ = time.Parse(time.ANSIC, s)
	start = start.Add(time.Duration(tsStartHours) * time.Hour)
	start = start.Add(time.Duration(tsStartMinutes) * time.Minute)

	end, _ = time.Parse(time.ANSIC, s)
	end = end.Add(time.Duration(tsEndHours) * time.Hour)
	end = end.Add(time.Duration(tsEndMinutes) * time.Minute)

	return start, end, nil
}
