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
func TFilter(filenames []string, tr *TimeRange, ignore bool, isReject bool) []string {

  if isReject  {
    return reject(filenames, tr, ignore)
  }

	return filter(filenames, tr, ignore)
}



func reject(filenames []string, tr *TimeRange, ignore bool) []string {
	var results []string
	timeHoriz := genTimeHoriz(tr)
	isTimeActive := tr.Days+tr.Weeks+tr.Months >= 1

	// Start from most recent time
	sort.Sort(sort.Reverse(sort.StringSlice(filenames)))

	for _, filename := range filenames {
		abs, _ := filepath.Abs(filename)
		t, err := parseDateFileName(filename)

		if ignore && err != nil {
			continue
		}

		// dont include if time range is passed
		if isTimeActive && !t.Before(timeHoriz) {
			continue
		}

		results = append(results, abs)
  }

	// return in original Chronological order
	sort.Sort(sort.StringSlice(results))
  return results
}


func genTimeHoriz(tr *TimeRange) time.Time {
  return time.Now().AddDate(
    0,
    -tr.Months,
    -tr.Days+ -tr.Weeks*7,
    )
}

// filter
// if Just a time is specified then its all within that timerange
//
// if Just a count is specified then its all up to that count
func filter(filenames []string, tr *TimeRange, ignore bool) []string {
  var results []string
  timeHoriz := genTimeHoriz(tr)
  isTimeActive := tr.Days+tr.Weeks+tr.Months >= 1

	// Start from most recent time
	// sort.Sort(sort.Reverse(sort.StringSlice(filenames)))

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

		results = append(results, abs)
	}

	// return in original Chronological order
	sort.Sort(sort.StringSlice(results))
	return results
}

