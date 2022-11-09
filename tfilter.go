package util

import (
	"path/filepath"
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
func TFilter(filnames []string, tr *TimeRange, count int, ignore bool) []string {
  if count > 0 {
    return filterCount(filnames, tr, count, ignore)
  }
  return filterTime(filnames, tr, count, ignore)
}

func filterTime(filnames []string, tr *TimeRange, count int, ignore bool) []string {
	var results []string
	// now sure why there is a one off error of 2 here... maybe timezones?
  days := -tr.Days-2 + -tr.Weeks*7
	timeHoriz := time.Now().AddDate(0, -tr.Months, days)
	for _, filename := range filnames {
		abs, _ := filepath.Abs(filename)
		t, err := parseDateFileName(filename)

		if ignore && err != nil {
			continue
		}

		if t.After(timeHoriz) {
			results = append(results, abs)
		}
	}
	return results;
}

func filterCount(filenames []string, tr *TimeRange, count int, ignore bool) []string {
  var results []string

	for _, filename := range filenames {
		abs, _ := filepath.Abs(filename)
		_, err := parseDateFileName(filename)

		if ignore && err != nil {
			continue
		}
    results = append(results, abs)
  }

  if count > len(results) {
    count = len(results)
  }

  // I want the last x of the array
  return results[len(results) - count:];
}
