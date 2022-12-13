package util

import (
	"path/filepath"
	"sort"
	"time"
)

type TimeRange struct {
	Months int
	Weeks  int
	Days   int
}

// Given some strings and a range of time from now
// filters those strings based on whether or not the date parsed from those
// strings fall within the range of time from now
func TFilter(filenames []string, tr *TimeRange, count int, ignore bool) []string {
	return filter(filenames, tr, count, ignore)
}

// filter
// if neither a count or time is specified then its the first 20
//
// If a time is specified and a count is specified, it returns up to count
// within the timerange
//
// if Just a time is specified then its all within that timerange
//
// if Just a count is specified then its all up to that count
func filter(filenames []string, tr *TimeRange, count int, ignore bool) []string {
	var results []string
	// not sure why there is a one off error of 2 here... maybe timezones?
	days := -tr.Days - 2 + -tr.Weeks*7
	timeHoriz := time.Now().AddDate(0, -tr.Months, days)
	isTimeActive := tr.Days+tr.Weeks+tr.Months > 1
	isCountActive := count > 0

	// count is whatever is specified or, if no time range is specified, 20
	if !isCountActive && !isTimeActive {
		count = 20
		isCountActive = true
	}

	// Start from most recent time
	sort.Sort(sort.Reverse(sort.StringSlice(filenames)))

	for _, filename := range filenames {
		abs, _ := filepath.Abs(filename)
		t, err := parseDateFileName(filename)

		if ignore && err != nil {
			continue
		}

		// dont include if time range is passed
		if isTimeActive && !t.After(timeHoriz) {
			continue
		}

		// dont include if there is no count left
		if isCountActive && count <= 0 {
			continue
		}

		results = append(results, abs)
		count -= 1
	}

	// return in original Chronological order
	sort.Sort(sort.StringSlice(results))
	return results
}
